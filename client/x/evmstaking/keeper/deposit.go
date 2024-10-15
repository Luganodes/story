package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	skeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/piplabs/story/client/x/evmstaking/types"
	"github.com/piplabs/story/contracts/bindings"
	"github.com/piplabs/story/lib/errors"
	"github.com/piplabs/story/lib/k1util"
	"github.com/piplabs/story/lib/log"
)

func (k Keeper) ProcessDeposit(ctx context.Context, ev *bindings.IPTokenStakingDeposit) error {
	delCmpPubkey, err := UncmpPubKeyToCmpPubKey(ev.DelegatorUncmpPubkey)
	if err != nil {
		return errors.Wrap(err, "compress delegator pubkey")
	}
	depositorPubkey, err := k1util.PubKeyBytesToCosmos(delCmpPubkey)
	if err != nil {
		return errors.Wrap(err, "depositor pubkey to cosmos")
	}

	valCmpPubkey, err := UncmpPubKeyToCmpPubKey(ev.ValidatorUnCmpPubkey)
	if err != nil {
		return errors.Wrap(err, "compress validator pubkey")
	}
	validatorPubkey, err := k1util.PubKeyBytesToCosmos(valCmpPubkey)
	if err != nil {
		return errors.Wrap(err, "validator pubkey to cosmos")
	}

	depositorAddr := sdk.AccAddress(depositorPubkey.Address().Bytes())
	validatorAddr := sdk.ValAddress(validatorPubkey.Address().Bytes())

	valEvmAddr, err := k1util.CosmosPubkeyToEVMAddress(validatorPubkey.Bytes())
	if err != nil {
		return errors.Wrap(err, "validator pubkey to evm address")
	}
	delEvmAddr, err := k1util.CosmosPubkeyToEVMAddress(depositorPubkey.Bytes())
	if err != nil {
		return errors.Wrap(err, "delegator pubkey to evm address")
	}

	amountCoin, amountCoins := IPTokenToBondCoin(ev.StakeAmount)

	// Create account if not exists
	if !k.authKeeper.HasAccount(ctx, depositorAddr) {
		acc := k.authKeeper.NewAccountWithAddress(ctx, depositorAddr)
		k.authKeeper.SetAccount(ctx, acc)
		log.Debug(ctx, "Created account for depositor",
			"address", depositorAddr.String(),
			"evm_address", delEvmAddr.String(),
		)
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, amountCoins); err != nil {
		return errors.Wrap(err, "create stake coin for depositor: mint coins")
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, depositorAddr, amountCoins); err != nil {
		return errors.Wrap(err, "create stake coin for depositor: send coins")
	}

	log.Info(ctx, "EVM staking deposit detected, delegating to validator",
		"del_story", depositorAddr.String(),
		"val_story", validatorAddr.String(),
		"del_evm_addr", delEvmAddr.String(),
		"val_evm_addr", valEvmAddr.String(),
		"amount_coin", amountCoin.String(),
	)

	// Note that, after minting, we save the mapping between delegator bech32 address and evm address, which will be used in the withdrawal queue.
	// The saving is done regardless of any error below, as the money is already minted and sent to the delegator, who can withdraw the minted amount.
	// TODO: Confirm that bech32 address and evm address can be used interchangeably. Must be one-to-one or many-bech32-to-one-evm.
	if err := k.DelegatorMap.Set(ctx, depositorAddr.String(), delEvmAddr.String()); err != nil {
		return errors.Wrap(err, "set delegator map")
	}

	// TODO: Check if we can instantiate the msgServer without type assertion
	evmstakingSKeeper, ok := k.stakingKeeper.(*skeeper.Keeper)
	if !ok {
		return errors.New("type assertion failed")
	}
	skeeperMsgServer := skeeper.NewMsgServerImpl(evmstakingSKeeper)

	var periodType stypes.PeriodType
	switch ev.StakingPeriod.Int64() {
	case int64(stypes.PeriodType_FLEXIBLE):
		periodType = stypes.PeriodType_FLEXIBLE
	case int64(stypes.PeriodType_THREE_MONTHS):
		periodType = stypes.PeriodType_THREE_MONTHS
	case int64(stypes.PeriodType_ONE_YEAR):
		periodType = stypes.PeriodType_ONE_YEAR
	case int64(stypes.PeriodType_EIGHTEEN_MONTHS):
		periodType = stypes.PeriodType_EIGHTEEN_MONTHS
	default:
		return errors.New("invalid staking period")
	}

	// Delegation by the depositor on the validator (validator existence is checked in msgServer.Delegate)
	msg := stypes.NewMsgDelegate(
		depositorAddr.String(), validatorAddr.String(), amountCoin,
		ev.DelegationId.String(), periodType,
	)
	_, err = skeeperMsgServer.Delegate(ctx, msg)
	if err != nil {
		return errors.Wrap(err, "delegate")
	}

	return nil
}

func (k Keeper) ParseDepositLog(ethlog ethtypes.Log) (*bindings.IPTokenStakingDeposit, error) {
	return k.ipTokenStakingContract.ParseDeposit(ethlog)
}
