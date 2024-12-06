// Code generated by MockGen. DO NOT EDIT.
// Source: init.go
//
// Generated by this command:
//
//	mockgen -source=init.go -destination=mocks/mockgen.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockcartRepository is a mock of cartRepository interface.
type MockcartRepository struct {
	ctrl     *gomock.Controller
	recorder *MockcartRepositoryMockRecorder
	isgomock struct{}
}

// MockcartRepositoryMockRecorder is the mock recorder for MockcartRepository.
type MockcartRepositoryMockRecorder struct {
	mock *MockcartRepository
}

// NewMockcartRepository creates a new mock instance.
func NewMockcartRepository(ctrl *gomock.Controller) *MockcartRepository {
	mock := &MockcartRepository{ctrl: ctrl}
	mock.recorder = &MockcartRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcartRepository) EXPECT() *MockcartRepositoryMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockcartRepository) AddProduct(ctx context.Context, productID, userID uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", ctx, productID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockcartRepositoryMockRecorder) AddProduct(ctx, productID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockcartRepository)(nil).AddProduct), ctx, productID, userID)
}

// ChangeCartProductCount mocks base method.
func (m *MockcartRepository) ChangeCartProductCount(ctx context.Context, productID uint32, count int32, userID uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCartProductCount", ctx, productID, count, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCartProductCount indicates an expected call of ChangeCartProductCount.
func (mr *MockcartRepositoryMockRecorder) ChangeCartProductCount(ctx, productID, count, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCartProductCount", reflect.TypeOf((*MockcartRepository)(nil).ChangeCartProductCount), ctx, productID, count, userID)
}

// ChangeCartProductDeletedState mocks base method.
func (m *MockcartRepository) ChangeCartProductDeletedState(ctx context.Context, productID, userID uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCartProductDeletedState", ctx, productID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCartProductDeletedState indicates an expected call of ChangeCartProductDeletedState.
func (mr *MockcartRepositoryMockRecorder) ChangeCartProductDeletedState(ctx, productID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCartProductDeletedState", reflect.TypeOf((*MockcartRepository)(nil).ChangeCartProductDeletedState), ctx, productID, userID)
}

// ChangeCartProductSelectedState mocks base method.
func (m *MockcartRepository) ChangeCartProductSelectedState(ctx context.Context, productID, userID uint32, isSelected bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCartProductSelectedState", ctx, productID, userID, isSelected)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCartProductSelectedState indicates an expected call of ChangeCartProductSelectedState.
func (mr *MockcartRepositoryMockRecorder) ChangeCartProductSelectedState(ctx, productID, userID, isSelected any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCartProductSelectedState", reflect.TypeOf((*MockcartRepository)(nil).ChangeCartProductSelectedState), ctx, productID, userID, isSelected)
}

// GetCartProduct mocks base method.
func (m *MockcartRepository) GetCartProduct(ctx context.Context, productID, userID uint32) (model.CartProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartProduct", ctx, productID, userID)
	ret0, _ := ret[0].(model.CartProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartProduct indicates an expected call of GetCartProduct.
func (mr *MockcartRepositoryMockRecorder) GetCartProduct(ctx, productID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartProduct", reflect.TypeOf((*MockcartRepository)(nil).GetCartProduct), ctx, productID, userID)
}

// GetSelectedCartItems mocks base method.
func (m *MockcartRepository) GetSelectedCartItems(ctx context.Context, userID uint32) ([]model.ProductOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelectedCartItems", ctx, userID)
	ret0, _ := ret[0].([]model.ProductOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSelectedCartItems indicates an expected call of GetSelectedCartItems.
func (mr *MockcartRepositoryMockRecorder) GetSelectedCartItems(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelectedCartItems", reflect.TypeOf((*MockcartRepository)(nil).GetSelectedCartItems), ctx, userID)
}

// GetSelectedFromCart mocks base method.
func (m *MockcartRepository) GetSelectedFromCart(ctx context.Context, userID uint32) (*model.CartProductsForOrderWithUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelectedFromCart", ctx, userID)
	ret0, _ := ret[0].(*model.CartProductsForOrderWithUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSelectedFromCart indicates an expected call of GetSelectedFromCart.
func (mr *MockcartRepositoryMockRecorder) GetSelectedFromCart(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelectedFromCart", reflect.TypeOf((*MockcartRepository)(nil).GetSelectedFromCart), ctx, userID)
}

// RemoveCartProduct mocks base method.
func (m *MockcartRepository) RemoveCartProduct(ctx context.Context, productID uint32, count int32, userID uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCartProduct", ctx, productID, count, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCartProduct indicates an expected call of RemoveCartProduct.
func (mr *MockcartRepositoryMockRecorder) RemoveCartProduct(ctx, productID, count, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCartProduct", reflect.TypeOf((*MockcartRepository)(nil).RemoveCartProduct), ctx, productID, count, userID)
}

// MockproductCountGetter is a mock of productCountGetter interface.
type MockproductCountGetter struct {
	ctrl     *gomock.Controller
	recorder *MockproductCountGetterMockRecorder
	isgomock struct{}
}

// MockproductCountGetterMockRecorder is the mock recorder for MockproductCountGetter.
type MockproductCountGetterMockRecorder struct {
	mock *MockproductCountGetter
}

// NewMockproductCountGetter creates a new mock instance.
func NewMockproductCountGetter(ctrl *gomock.Controller) *MockproductCountGetter {
	mock := &MockproductCountGetter{ctrl: ctrl}
	mock.recorder = &MockproductCountGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockproductCountGetter) EXPECT() *MockproductCountGetterMockRecorder {
	return m.recorder
}

// GetProductCount mocks base method.
func (m *MockproductCountGetter) GetProductCount(ctx context.Context, productID uint32) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductCount", ctx, productID)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductCount indicates an expected call of GetProductCount.
func (mr *MockproductCountGetterMockRecorder) GetProductCount(ctx, productID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductCount", reflect.TypeOf((*MockproductCountGetter)(nil).GetProductCount), ctx, productID)
}