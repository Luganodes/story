// Code generated by MockGen. DO NOT EDIT.
// Source: ../client/x/evmengine/types/expected_keepers.go
//
// Generated by this command:
//
//	mockgen -source=../client/x/evmengine/types/expected_keepers.go -package testutil -destination ../client/x/evmengine/testutil/expected_keepers_mocks.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	math "cosmossdk.io/math"
	types "cosmossdk.io/x/upgrade/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/ethereum/go-ethereum/core/types"
	types2 "github.com/piplabs/story/client/x/evmengine/types"
	bindings "github.com/piplabs/story/contracts/bindings"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(moduleName string) types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", moduleName)
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(moduleName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), moduleName)
}

// MockEvmStakingKeeper is a mock of EvmStakingKeeper interface.
type MockEvmStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockEvmStakingKeeperMockRecorder
}

// MockEvmStakingKeeperMockRecorder is the mock recorder for MockEvmStakingKeeper.
type MockEvmStakingKeeperMockRecorder struct {
	mock *MockEvmStakingKeeper
}

// NewMockEvmStakingKeeper creates a new mock instance.
func NewMockEvmStakingKeeper(ctrl *gomock.Controller) *MockEvmStakingKeeper {
	mock := &MockEvmStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockEvmStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvmStakingKeeper) EXPECT() *MockEvmStakingKeeperMockRecorder {
	return m.recorder
}

// DequeueEligibleRewardWithdrawals mocks base method.
func (m *MockEvmStakingKeeper) DequeueEligibleRewardWithdrawals(ctx context.Context, maxDequeue uint32) (types1.Withdrawals, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DequeueEligibleRewardWithdrawals", ctx, maxDequeue)
	ret0, _ := ret[0].(types1.Withdrawals)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DequeueEligibleRewardWithdrawals indicates an expected call of DequeueEligibleRewardWithdrawals.
func (mr *MockEvmStakingKeeperMockRecorder) DequeueEligibleRewardWithdrawals(ctx, maxDequeue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DequeueEligibleRewardWithdrawals", reflect.TypeOf((*MockEvmStakingKeeper)(nil).DequeueEligibleRewardWithdrawals), ctx, maxDequeue)
}

// DequeueEligibleWithdrawals mocks base method.
func (m *MockEvmStakingKeeper) DequeueEligibleWithdrawals(ctx context.Context, maxDequeue uint32) (types1.Withdrawals, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DequeueEligibleWithdrawals", ctx, maxDequeue)
	ret0, _ := ret[0].(types1.Withdrawals)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DequeueEligibleWithdrawals indicates an expected call of DequeueEligibleWithdrawals.
func (mr *MockEvmStakingKeeperMockRecorder) DequeueEligibleWithdrawals(ctx, maxDequeue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DequeueEligibleWithdrawals", reflect.TypeOf((*MockEvmStakingKeeper)(nil).DequeueEligibleWithdrawals), ctx, maxDequeue)
}

// MaxWithdrawalPerBlock mocks base method.
func (m *MockEvmStakingKeeper) MaxWithdrawalPerBlock(ctx context.Context) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxWithdrawalPerBlock", ctx)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaxWithdrawalPerBlock indicates an expected call of MaxWithdrawalPerBlock.
func (mr *MockEvmStakingKeeperMockRecorder) MaxWithdrawalPerBlock(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxWithdrawalPerBlock", reflect.TypeOf((*MockEvmStakingKeeper)(nil).MaxWithdrawalPerBlock), ctx)
}

// ParseDepositLog mocks base method.
func (m *MockEvmStakingKeeper) ParseDepositLog(ethlog types1.Log) (*bindings.IPTokenStakingDeposit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseDepositLog", ethlog)
	ret0, _ := ret[0].(*bindings.IPTokenStakingDeposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseDepositLog indicates an expected call of ParseDepositLog.
func (mr *MockEvmStakingKeeperMockRecorder) ParseDepositLog(ethlog any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseDepositLog", reflect.TypeOf((*MockEvmStakingKeeper)(nil).ParseDepositLog), ethlog)
}

// ParseWithdrawLog mocks base method.
func (m *MockEvmStakingKeeper) ParseWithdrawLog(ethlog types1.Log) (*bindings.IPTokenStakingWithdraw, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseWithdrawLog", ethlog)
	ret0, _ := ret[0].(*bindings.IPTokenStakingWithdraw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseWithdrawLog indicates an expected call of ParseWithdrawLog.
func (mr *MockEvmStakingKeeperMockRecorder) ParseWithdrawLog(ethlog any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseWithdrawLog", reflect.TypeOf((*MockEvmStakingKeeper)(nil).ParseWithdrawLog), ethlog)
}

// PeekEligibleRewardWithdrawals mocks base method.
func (m *MockEvmStakingKeeper) PeekEligibleRewardWithdrawals(ctx context.Context, maxPeek uint32) (types1.Withdrawals, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekEligibleRewardWithdrawals", ctx, maxPeek)
	ret0, _ := ret[0].(types1.Withdrawals)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeekEligibleRewardWithdrawals indicates an expected call of PeekEligibleRewardWithdrawals.
func (mr *MockEvmStakingKeeperMockRecorder) PeekEligibleRewardWithdrawals(ctx, maxPeek any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekEligibleRewardWithdrawals", reflect.TypeOf((*MockEvmStakingKeeper)(nil).PeekEligibleRewardWithdrawals), ctx, maxPeek)
}

// PeekEligibleWithdrawals mocks base method.
func (m *MockEvmStakingKeeper) PeekEligibleWithdrawals(ctx context.Context, maxPeek uint32) (types1.Withdrawals, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeekEligibleWithdrawals", ctx, maxPeek)
	ret0, _ := ret[0].(types1.Withdrawals)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeekEligibleWithdrawals indicates an expected call of PeekEligibleWithdrawals.
func (mr *MockEvmStakingKeeperMockRecorder) PeekEligibleWithdrawals(ctx, maxPeek any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeekEligibleWithdrawals", reflect.TypeOf((*MockEvmStakingKeeper)(nil).PeekEligibleWithdrawals), ctx, maxPeek)
}

// ProcessStakingEvents mocks base method.
func (m *MockEvmStakingKeeper) ProcessStakingEvents(ctx context.Context, height uint64, logs []*types2.EVMEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessStakingEvents", ctx, height, logs)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessStakingEvents indicates an expected call of ProcessStakingEvents.
func (mr *MockEvmStakingKeeperMockRecorder) ProcessStakingEvents(ctx, height, logs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessStakingEvents", reflect.TypeOf((*MockEvmStakingKeeper)(nil).ProcessStakingEvents), ctx, height, logs)
}

// MockUpgradeKeeper is a mock of UpgradeKeeper interface.
type MockUpgradeKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeKeeperMockRecorder
}

// MockUpgradeKeeperMockRecorder is the mock recorder for MockUpgradeKeeper.
type MockUpgradeKeeperMockRecorder struct {
	mock *MockUpgradeKeeper
}

// NewMockUpgradeKeeper creates a new mock instance.
func NewMockUpgradeKeeper(ctrl *gomock.Controller) *MockUpgradeKeeper {
	mock := &MockUpgradeKeeper{ctrl: ctrl}
	mock.recorder = &MockUpgradeKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeKeeper) EXPECT() *MockUpgradeKeeperMockRecorder {
	return m.recorder
}

// ClearUpgradePlan mocks base method.
func (m *MockUpgradeKeeper) ClearUpgradePlan(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearUpgradePlan", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearUpgradePlan indicates an expected call of ClearUpgradePlan.
func (mr *MockUpgradeKeeperMockRecorder) ClearUpgradePlan(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearUpgradePlan", reflect.TypeOf((*MockUpgradeKeeper)(nil).ClearUpgradePlan), ctx)
}

// ScheduleUpgrade mocks base method.
func (m *MockUpgradeKeeper) ScheduleUpgrade(ctx context.Context, plan types.Plan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleUpgrade", ctx, plan)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleUpgrade indicates an expected call of ScheduleUpgrade.
func (mr *MockUpgradeKeeperMockRecorder) ScheduleUpgrade(ctx, plan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleUpgrade", reflect.TypeOf((*MockUpgradeKeeper)(nil).ScheduleUpgrade), ctx, plan)
}

// MockDistrKeeper is a mock of DistrKeeper interface.
type MockDistrKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockDistrKeeperMockRecorder
}

// MockDistrKeeperMockRecorder is the mock recorder for MockDistrKeeper.
type MockDistrKeeperMockRecorder struct {
	mock *MockDistrKeeper
}

// NewMockDistrKeeper creates a new mock instance.
func NewMockDistrKeeper(ctrl *gomock.Controller) *MockDistrKeeper {
	mock := &MockDistrKeeper{ctrl: ctrl}
	mock.recorder = &MockDistrKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDistrKeeper) EXPECT() *MockDistrKeeperMockRecorder {
	return m.recorder
}

// SetUbi mocks base method.
func (m *MockDistrKeeper) SetUbi(ctx context.Context, newUbi math.LegacyDec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUbi", ctx, newUbi)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUbi indicates an expected call of SetUbi.
func (mr *MockDistrKeeperMockRecorder) SetUbi(ctx, newUbi any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUbi", reflect.TypeOf((*MockDistrKeeper)(nil).SetUbi), ctx, newUbi)
}
