// Code generated by MockGen. DO NOT EDIT.
// Source: init.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockpromoCodesManager is a mock of promoCodesManager interface.
type MockpromoCodesManager struct {
	ctrl     *gomock.Controller
	recorder *MockpromoCodesManagerMockRecorder
}

// MockpromoCodesManagerMockRecorder is the mock recorder for MockpromoCodesManager.
type MockpromoCodesManagerMockRecorder struct {
	mock *MockpromoCodesManager
}

// NewMockpromoCodesManager creates a new mock instance.
func NewMockpromoCodesManager(ctrl *gomock.Controller) *MockpromoCodesManager {
	mock := &MockpromoCodesManager{ctrl: ctrl}
	mock.recorder = &MockpromoCodesManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpromoCodesManager) EXPECT() *MockpromoCodesManagerMockRecorder {
	return m.recorder
}

// DeletePromoCode mocks base method.
func (m *MockpromoCodesManager) DeletePromoCode(ctx context.Context, userID, promoID uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePromoCode", ctx, userID, promoID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePromoCode indicates an expected call of DeletePromoCode.
func (mr *MockpromoCodesManagerMockRecorder) DeletePromoCode(ctx, userID, promoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePromoCode", reflect.TypeOf((*MockpromoCodesManager)(nil).DeletePromoCode), ctx, userID, promoID)
}

// GetPromoCode mocks base method.
func (m *MockpromoCodesManager) GetPromoCode(ctx context.Context, userID uint32, promoCodeName string) (model.PromoCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPromoCode", ctx, userID, promoCodeName)
	ret0, _ := ret[0].(model.PromoCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPromoCode indicates an expected call of GetPromoCode.
func (mr *MockpromoCodesManagerMockRecorder) GetPromoCode(ctx, userID, promoCodeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPromoCode", reflect.TypeOf((*MockpromoCodesManager)(nil).GetPromoCode), ctx, userID, promoCodeName)
}

// MockOrdersRepo is a mock of OrdersRepo interface.
type MockOrdersRepo struct {
	ctrl     *gomock.Controller
	recorder *MockOrdersRepoMockRecorder
}

// MockOrdersRepoMockRecorder is the mock recorder for MockOrdersRepo.
type MockOrdersRepoMockRecorder struct {
	mock *MockOrdersRepo
}

// NewMockOrdersRepo creates a new mock instance.
func NewMockOrdersRepo(ctrl *gomock.Controller) *MockOrdersRepo {
	mock := &MockOrdersRepo{ctrl: ctrl}
	mock.recorder = &MockOrdersRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrdersRepo) EXPECT() *MockOrdersRepoMockRecorder {
	return m.recorder
}

// CreateOrderFromCart mocks base method.
func (m *MockOrdersRepo) CreateOrderFromCart(ctx context.Context, orderData *model.OrderFromCart) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderFromCart", ctx, orderData)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderFromCart indicates an expected call of CreateOrderFromCart.
func (mr *MockOrdersRepoMockRecorder) CreateOrderFromCart(ctx, orderData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderFromCart", reflect.TypeOf((*MockOrdersRepo)(nil).CreateOrderFromCart), ctx, orderData)
}

// GetNearestDeliveryDate mocks base method.
func (m *MockOrdersRepo) GetNearestDeliveryDate(ctx context.Context, userID uint32) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNearestDeliveryDate", ctx, userID)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNearestDeliveryDate indicates an expected call of GetNearestDeliveryDate.
func (mr *MockOrdersRepoMockRecorder) GetNearestDeliveryDate(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNearestDeliveryDate", reflect.TypeOf((*MockOrdersRepo)(nil).GetNearestDeliveryDate), ctx, userID)
}

// GetOrderById mocks base method.
func (m *MockOrdersRepo) GetOrderById(ctx context.Context, id uuid.UUID, userID uint32) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", ctx, id, userID)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockOrdersRepoMockRecorder) GetOrderById(ctx, id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockOrdersRepo)(nil).GetOrderById), ctx, id, userID)
}

// GetOrders mocks base method.
func (m *MockOrdersRepo) GetOrders(ctx context.Context, userId uint32) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", ctx, userId)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrdersRepoMockRecorder) GetOrders(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrdersRepo)(nil).GetOrders), ctx, userId)
}

// MockcartGetter is a mock of cartGetter interface.
type MockcartGetter struct {
	ctrl     *gomock.Controller
	recorder *MockcartGetterMockRecorder
}

// MockcartGetterMockRecorder is the mock recorder for MockcartGetter.
type MockcartGetterMockRecorder struct {
	mock *MockcartGetter
}

// NewMockcartGetter creates a new mock instance.
func NewMockcartGetter(ctrl *gomock.Controller) *MockcartGetter {
	mock := &MockcartGetter{ctrl: ctrl}
	mock.recorder = &MockcartGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcartGetter) EXPECT() *MockcartGetterMockRecorder {
	return m.recorder
}

// GetSelectedCartItems mocks base method.
func (m *MockcartGetter) GetSelectedCartItems(ctx context.Context, userID uint32) ([]model.ProductOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelectedCartItems", ctx, userID)
	ret0, _ := ret[0].([]model.ProductOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSelectedCartItems indicates an expected call of GetSelectedCartItems.
func (mr *MockcartGetterMockRecorder) GetSelectedCartItems(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelectedCartItems", reflect.TypeOf((*MockcartGetter)(nil).GetSelectedCartItems), ctx, userID)
}
