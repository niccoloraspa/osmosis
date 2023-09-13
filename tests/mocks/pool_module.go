// Code generated by MockGen. DO NOT EDIT.
// Source: x/poolmanager/types/expected_keepers.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/bank/types"
	gomock "github.com/golang/mock/gomock"
	osmomath "github.com/osmosis-labs/osmosis/osmomath"
	types2 "github.com/osmosis-labs/osmosis/v19/x/poolmanager/types"
)

// MockAccountI is a mock of AccountI interface.
type MockAccountI struct {
	ctrl     *gomock.Controller
	recorder *MockAccountIMockRecorder
}

// MockAccountIMockRecorder is the mock recorder for MockAccountI.
type MockAccountIMockRecorder struct {
	mock *MockAccountI
}

// NewMockAccountI creates a new mock instance.
func NewMockAccountI(ctrl *gomock.Controller) *MockAccountI {
	mock := &MockAccountI{ctrl: ctrl}
	mock.recorder = &MockAccountIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountI) EXPECT() *MockAccountIMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountI) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountIMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountI)(nil).GetAccount), ctx, addr)
}

// GetModuleAccount mocks base method.
func (m *MockAccountI) GetModuleAccount(ctx types.Context, moduleName string) types0.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, moduleName)
	ret0, _ := ret[0].(types0.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountIMockRecorder) GetModuleAccount(ctx, moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountI)(nil).GetModuleAccount), ctx, moduleName)
}

// NewAccount mocks base method.
func (m *MockAccountI) NewAccount(arg0 types.Context, arg1 types0.AccountI) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0, arg1)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// NewAccount indicates an expected call of NewAccount.
func (mr *MockAccountIMockRecorder) NewAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountI)(nil).NewAccount), arg0, arg1)
}

// SetAccount mocks base method.
func (m *MockAccountI) SetAccount(ctx types.Context, acc types0.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountIMockRecorder) SetAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountI)(nil).SetAccount), ctx, acc)
}

// MockBankI is a mock of BankI interface.
type MockBankI struct {
	ctrl     *gomock.Controller
	recorder *MockBankIMockRecorder
}

// MockBankIMockRecorder is the mock recorder for MockBankI.
type MockBankIMockRecorder struct {
	mock *MockBankI
}

// NewMockBankI creates a new mock instance.
func NewMockBankI(ctrl *gomock.Controller) *MockBankI {
	mock := &MockBankI{ctrl: ctrl}
	mock.recorder = &MockBankIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankI) EXPECT() *MockBankIMockRecorder {
	return m.recorder
}

// GetAllBalances mocks base method.
func (m *MockBankI) GetAllBalances(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankIMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankI)(nil).GetAllBalances), ctx, addr)
}

// SendCoins mocks base method.
func (m *MockBankI) SendCoins(ctx types.Context, fromAddr, toAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, fromAddr, toAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankIMockRecorder) SendCoins(ctx, fromAddr, toAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankI)(nil).SendCoins), ctx, fromAddr, toAddr, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankI) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankIMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankI)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SetDenomMetaData mocks base method.
func (m *MockBankI) SetDenomMetaData(ctx types.Context, denomMetaData types1.Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDenomMetaData", ctx, denomMetaData)
}

// SetDenomMetaData indicates an expected call of SetDenomMetaData.
func (mr *MockBankIMockRecorder) SetDenomMetaData(ctx, denomMetaData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDenomMetaData", reflect.TypeOf((*MockBankI)(nil).SetDenomMetaData), ctx, denomMetaData)
}

// MockCommunityPoolI is a mock of CommunityPoolI interface.
type MockCommunityPoolI struct {
	ctrl     *gomock.Controller
	recorder *MockCommunityPoolIMockRecorder
}

// MockCommunityPoolIMockRecorder is the mock recorder for MockCommunityPoolI.
type MockCommunityPoolIMockRecorder struct {
	mock *MockCommunityPoolI
}

// NewMockCommunityPoolI creates a new mock instance.
func NewMockCommunityPoolI(ctrl *gomock.Controller) *MockCommunityPoolI {
	mock := &MockCommunityPoolI{ctrl: ctrl}
	mock.recorder = &MockCommunityPoolIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommunityPoolI) EXPECT() *MockCommunityPoolIMockRecorder {
	return m.recorder
}

// FundCommunityPool mocks base method.
func (m *MockCommunityPoolI) FundCommunityPool(ctx types.Context, amount types.Coins, sender types.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FundCommunityPool", ctx, amount, sender)
	ret0, _ := ret[0].(error)
	return ret0
}

// FundCommunityPool indicates an expected call of FundCommunityPool.
func (mr *MockCommunityPoolIMockRecorder) FundCommunityPool(ctx, amount, sender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FundCommunityPool", reflect.TypeOf((*MockCommunityPoolI)(nil).FundCommunityPool), ctx, amount, sender)
}

// MockPoolModuleI is a mock of PoolModuleI interface.
type MockPoolModuleI struct {
	ctrl     *gomock.Controller
	recorder *MockPoolModuleIMockRecorder
}

// MockPoolModuleIMockRecorder is the mock recorder for MockPoolModuleI.
type MockPoolModuleIMockRecorder struct {
	mock *MockPoolModuleI
}

// NewMockPoolModuleI creates a new mock instance.
func NewMockPoolModuleI(ctrl *gomock.Controller) *MockPoolModuleI {
	mock := &MockPoolModuleI{ctrl: ctrl}
	mock.recorder = &MockPoolModuleIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPoolModuleI) EXPECT() *MockPoolModuleIMockRecorder {
	return m.recorder
}

// CalcInAmtGivenOut mocks base method.
func (m *MockPoolModuleI) CalcInAmtGivenOut(ctx types.Context, poolI types2.PoolI, tokenOut types.Coin, tokenInDenom string, spreadFactor osmomath.Dec) (types.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcInAmtGivenOut", ctx, poolI, tokenOut, tokenInDenom, spreadFactor)
	ret0, _ := ret[0].(types.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalcInAmtGivenOut indicates an expected call of CalcInAmtGivenOut.
func (mr *MockPoolModuleIMockRecorder) CalcInAmtGivenOut(ctx, poolI, tokenOut, tokenInDenom, spreadFactor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcInAmtGivenOut", reflect.TypeOf((*MockPoolModuleI)(nil).CalcInAmtGivenOut), ctx, poolI, tokenOut, tokenInDenom, spreadFactor)
}

// CalcOutAmtGivenIn mocks base method.
func (m *MockPoolModuleI) CalcOutAmtGivenIn(ctx types.Context, poolI types2.PoolI, tokenIn types.Coin, tokenOutDenom string, spreadFactor osmomath.Dec) (types.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcOutAmtGivenIn", ctx, poolI, tokenIn, tokenOutDenom, spreadFactor)
	ret0, _ := ret[0].(types.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalcOutAmtGivenIn indicates an expected call of CalcOutAmtGivenIn.
func (mr *MockPoolModuleIMockRecorder) CalcOutAmtGivenIn(ctx, poolI, tokenIn, tokenOutDenom, spreadFactor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcOutAmtGivenIn", reflect.TypeOf((*MockPoolModuleI)(nil).CalcOutAmtGivenIn), ctx, poolI, tokenIn, tokenOutDenom, spreadFactor)
}

// CalculateSpotPrice mocks base method.
func (m *MockPoolModuleI) CalculateSpotPrice(ctx types.Context, poolId uint64, quoteAssetDenom, baseAssetDenom string) (osmomath.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateSpotPrice", ctx, poolId, quoteAssetDenom, baseAssetDenom)
	ret0, _ := ret[0].(osmomath.Dec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateSpotPrice indicates an expected call of CalculateSpotPrice.
func (mr *MockPoolModuleIMockRecorder) CalculateSpotPrice(ctx, poolId, quoteAssetDenom, baseAssetDenom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateSpotPrice", reflect.TypeOf((*MockPoolModuleI)(nil).CalculateSpotPrice), ctx, poolId, quoteAssetDenom, baseAssetDenom)
}

// GetPool mocks base method.
func (m *MockPoolModuleI) GetPool(ctx types.Context, poolId uint64) (types2.PoolI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPool", ctx, poolId)
	ret0, _ := ret[0].(types2.PoolI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPool indicates an expected call of GetPool.
func (mr *MockPoolModuleIMockRecorder) GetPool(ctx, poolId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPool", reflect.TypeOf((*MockPoolModuleI)(nil).GetPool), ctx, poolId)
}

// GetPoolDenoms mocks base method.
func (m *MockPoolModuleI) GetPoolDenoms(ctx types.Context, poolId uint64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoolDenoms", ctx, poolId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPoolDenoms indicates an expected call of GetPoolDenoms.
func (mr *MockPoolModuleIMockRecorder) GetPoolDenoms(ctx, poolId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoolDenoms", reflect.TypeOf((*MockPoolModuleI)(nil).GetPoolDenoms), ctx, poolId)
}

// GetPools mocks base method.
func (m *MockPoolModuleI) GetPools(ctx types.Context) ([]types2.PoolI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPools", ctx)
	ret0, _ := ret[0].([]types2.PoolI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPools indicates an expected call of GetPools.
func (mr *MockPoolModuleIMockRecorder) GetPools(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPools", reflect.TypeOf((*MockPoolModuleI)(nil).GetPools), ctx)
}

// GetTotalLiquidity mocks base method.
func (m *MockPoolModuleI) GetTotalLiquidity(ctx types.Context) (types.Coins, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalLiquidity", ctx)
	ret0, _ := ret[0].(types.Coins)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalLiquidity indicates an expected call of GetTotalLiquidity.
func (mr *MockPoolModuleIMockRecorder) GetTotalLiquidity(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalLiquidity", reflect.TypeOf((*MockPoolModuleI)(nil).GetTotalLiquidity), ctx)
}

// GetTotalPoolLiquidity mocks base method.
func (m *MockPoolModuleI) GetTotalPoolLiquidity(ctx types.Context, poolId uint64) (types.Coins, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalPoolLiquidity", ctx, poolId)
	ret0, _ := ret[0].(types.Coins)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalPoolLiquidity indicates an expected call of GetTotalPoolLiquidity.
func (mr *MockPoolModuleIMockRecorder) GetTotalPoolLiquidity(ctx, poolId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalPoolLiquidity", reflect.TypeOf((*MockPoolModuleI)(nil).GetTotalPoolLiquidity), ctx, poolId)
}

// InitializePool mocks base method.
func (m *MockPoolModuleI) InitializePool(ctx types.Context, pool types2.PoolI, creatorAddress types.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializePool", ctx, pool, creatorAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializePool indicates an expected call of InitializePool.
func (mr *MockPoolModuleIMockRecorder) InitializePool(ctx, pool, creatorAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializePool", reflect.TypeOf((*MockPoolModuleI)(nil).InitializePool), ctx, pool, creatorAddress)
}

// SwapExactAmountIn mocks base method.
func (m *MockPoolModuleI) SwapExactAmountIn(ctx types.Context, sender types.AccAddress, pool types2.PoolI, tokenIn types.Coin, tokenOutDenom string, tokenOutMinAmount osmomath.Int, spreadFactor osmomath.Dec) (osmomath.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwapExactAmountIn", ctx, sender, pool, tokenIn, tokenOutDenom, tokenOutMinAmount, spreadFactor)
	ret0, _ := ret[0].(osmomath.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SwapExactAmountIn indicates an expected call of SwapExactAmountIn.
func (mr *MockPoolModuleIMockRecorder) SwapExactAmountIn(ctx, sender, pool, tokenIn, tokenOutDenom, tokenOutMinAmount, spreadFactor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwapExactAmountIn", reflect.TypeOf((*MockPoolModuleI)(nil).SwapExactAmountIn), ctx, sender, pool, tokenIn, tokenOutDenom, tokenOutMinAmount, spreadFactor)
}

// SwapExactAmountOut mocks base method.
func (m *MockPoolModuleI) SwapExactAmountOut(ctx types.Context, sender types.AccAddress, pool types2.PoolI, tokenInDenom string, tokenInMaxAmount osmomath.Int, tokenOut types.Coin, spreadFactor osmomath.Dec) (osmomath.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwapExactAmountOut", ctx, sender, pool, tokenInDenom, tokenInMaxAmount, tokenOut, spreadFactor)
	ret0, _ := ret[0].(osmomath.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SwapExactAmountOut indicates an expected call of SwapExactAmountOut.
func (mr *MockPoolModuleIMockRecorder) SwapExactAmountOut(ctx, sender, pool, tokenInDenom, tokenInMaxAmount, tokenOut, spreadFactor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwapExactAmountOut", reflect.TypeOf((*MockPoolModuleI)(nil).SwapExactAmountOut), ctx, sender, pool, tokenInDenom, tokenInMaxAmount, tokenOut, spreadFactor)
}

// ValidatePermissionlessPoolCreationEnabled mocks base method.
func (m *MockPoolModuleI) ValidatePermissionlessPoolCreationEnabled(ctx types.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePermissionlessPoolCreationEnabled", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidatePermissionlessPoolCreationEnabled indicates an expected call of ValidatePermissionlessPoolCreationEnabled.
func (mr *MockPoolModuleIMockRecorder) ValidatePermissionlessPoolCreationEnabled(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePermissionlessPoolCreationEnabled", reflect.TypeOf((*MockPoolModuleI)(nil).ValidatePermissionlessPoolCreationEnabled), ctx)
}

// ValidatePoolHasUniqueParams mocks base method.
func (m *MockPoolModuleI) ValidatePoolHasUniqueParams(ctx types.Context, pool types2.PoolI, msg types2.CreatePoolMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePoolHasUniqueParams", ctx, pool, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidatePoolHasUniqueParams indicates an expected call of ValidatePoolHasUniqueParams.
func (mr *MockPoolModuleIMockRecorder) ValidatePoolHasUniqueParams(ctx, pool, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePoolHasUniqueParams", reflect.TypeOf((*MockPoolModuleI)(nil).ValidatePoolHasUniqueParams), ctx, pool, msg)
}

// MockPoolIncentivesKeeperI is a mock of PoolIncentivesKeeperI interface.
type MockPoolIncentivesKeeperI struct {
	ctrl     *gomock.Controller
	recorder *MockPoolIncentivesKeeperIMockRecorder
}

// MockPoolIncentivesKeeperIMockRecorder is the mock recorder for MockPoolIncentivesKeeperI.
type MockPoolIncentivesKeeperIMockRecorder struct {
	mock *MockPoolIncentivesKeeperI
}

// NewMockPoolIncentivesKeeperI creates a new mock instance.
func NewMockPoolIncentivesKeeperI(ctrl *gomock.Controller) *MockPoolIncentivesKeeperI {
	mock := &MockPoolIncentivesKeeperI{ctrl: ctrl}
	mock.recorder = &MockPoolIncentivesKeeperIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPoolIncentivesKeeperI) EXPECT() *MockPoolIncentivesKeeperIMockRecorder {
	return m.recorder
}

// IsPoolIncentivized mocks base method.
func (m *MockPoolIncentivesKeeperI) IsPoolIncentivized(ctx types.Context, poolId uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPoolIncentivized", ctx, poolId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPoolIncentivized indicates an expected call of IsPoolIncentivized.
func (mr *MockPoolIncentivesKeeperIMockRecorder) IsPoolIncentivized(ctx, poolId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPoolIncentivized", reflect.TypeOf((*MockPoolIncentivesKeeperI)(nil).IsPoolIncentivized), ctx, poolId)
}

// MockMultihopRoute is a mock of MultihopRoute interface.
type MockMultihopRoute struct {
	ctrl     *gomock.Controller
	recorder *MockMultihopRouteMockRecorder
}

// MockMultihopRouteMockRecorder is the mock recorder for MockMultihopRoute.
type MockMultihopRouteMockRecorder struct {
	mock *MockMultihopRoute
}

// NewMockMultihopRoute creates a new mock instance.
func NewMockMultihopRoute(ctrl *gomock.Controller) *MockMultihopRoute {
	mock := &MockMultihopRoute{ctrl: ctrl}
	mock.recorder = &MockMultihopRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMultihopRoute) EXPECT() *MockMultihopRouteMockRecorder {
	return m.recorder
}

// IntermediateDenoms mocks base method.
func (m *MockMultihopRoute) IntermediateDenoms() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntermediateDenoms")
	ret0, _ := ret[0].([]string)
	return ret0
}

// IntermediateDenoms indicates an expected call of IntermediateDenoms.
func (mr *MockMultihopRouteMockRecorder) IntermediateDenoms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntermediateDenoms", reflect.TypeOf((*MockMultihopRoute)(nil).IntermediateDenoms))
}

// Length mocks base method.
func (m *MockMultihopRoute) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockMultihopRouteMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockMultihopRoute)(nil).Length))
}

// PoolIds mocks base method.
func (m *MockMultihopRoute) PoolIds() []uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PoolIds")
	ret0, _ := ret[0].([]uint64)
	return ret0
}

// PoolIds indicates an expected call of PoolIds.
func (mr *MockMultihopRouteMockRecorder) PoolIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PoolIds", reflect.TypeOf((*MockMultihopRoute)(nil).PoolIds))
}

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// BondDenom mocks base method.
func (m *MockStakingKeeper) BondDenom(ctx types.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BondDenom", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// BondDenom indicates an expected call of BondDenom.
func (mr *MockStakingKeeperMockRecorder) BondDenom(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BondDenom", reflect.TypeOf((*MockStakingKeeper)(nil).BondDenom), ctx)
}

// MockProtorevKeeper is a mock of ProtorevKeeper interface.
type MockProtorevKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockProtorevKeeperMockRecorder
}

// MockProtorevKeeperMockRecorder is the mock recorder for MockProtorevKeeper.
type MockProtorevKeeperMockRecorder struct {
	mock *MockProtorevKeeper
}

// NewMockProtorevKeeper creates a new mock instance.
func NewMockProtorevKeeper(ctrl *gomock.Controller) *MockProtorevKeeper {
	mock := &MockProtorevKeeper{ctrl: ctrl}
	mock.recorder = &MockProtorevKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProtorevKeeper) EXPECT() *MockProtorevKeeperMockRecorder {
	return m.recorder
}

// GetPoolForDenomPair mocks base method.
func (m *MockProtorevKeeper) GetPoolForDenomPair(ctx types.Context, baseDenom, denomToMatch string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoolForDenomPair", ctx, baseDenom, denomToMatch)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPoolForDenomPair indicates an expected call of GetPoolForDenomPair.
func (mr *MockProtorevKeeperMockRecorder) GetPoolForDenomPair(ctx, baseDenom, denomToMatch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoolForDenomPair", reflect.TypeOf((*MockProtorevKeeper)(nil).GetPoolForDenomPair), ctx, baseDenom, denomToMatch)
}
