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

// MockusersRepository is a mock of usersRepository interface.
type MockusersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockusersRepositoryMockRecorder
	isgomock struct{}
}

// MockusersRepositoryMockRecorder is the mock recorder for MockusersRepository.
type MockusersRepositoryMockRecorder struct {
	mock *MockusersRepository
}

// NewMockusersRepository creates a new mock instance.
func NewMockusersRepository(ctrl *gomock.Controller) *MockusersRepository {
	mock := &MockusersRepository{ctrl: ctrl}
	mock.recorder = &MockusersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockusersRepository) EXPECT() *MockusersRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockusersRepository) CreateUser(ctx context.Context, userModel model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, userModel)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockusersRepositoryMockRecorder) CreateUser(ctx, userModel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockusersRepository)(nil).CreateUser), ctx, userModel)
}

// GetUserByEmail mocks base method.
func (m *MockusersRepository) GetUserByEmail(ctx context.Context, userModel model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, userModel)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockusersRepositoryMockRecorder) GetUserByEmail(ctx, userModel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockusersRepository)(nil).GetUserByEmail), ctx, userModel)
}

// GetUserByUserID mocks base method.
func (m *MockusersRepository) GetUserByUserID(ctx context.Context, id uint32) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUserID", ctx, id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUserID indicates an expected call of GetUserByUserID.
func (mr *MockusersRepositoryMockRecorder) GetUserByUserID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUserID", reflect.TypeOf((*MockusersRepository)(nil).GetUserByUserID), ctx, id)
}

// MocksessionGetter is a mock of sessionGetter interface.
type MocksessionGetter struct {
	ctrl     *gomock.Controller
	recorder *MocksessionGetterMockRecorder
	isgomock struct{}
}

// MocksessionGetterMockRecorder is the mock recorder for MocksessionGetter.
type MocksessionGetterMockRecorder struct {
	mock *MocksessionGetter
}

// NewMocksessionGetter creates a new mock instance.
func NewMocksessionGetter(ctrl *gomock.Controller) *MocksessionGetter {
	mock := &MocksessionGetter{ctrl: ctrl}
	mock.recorder = &MocksessionGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksessionGetter) EXPECT() *MocksessionGetterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MocksessionGetter) Get(ctx context.Context, sessionID string) (model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, sessionID)
	ret0, _ := ret[0].(model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MocksessionGetterMockRecorder) Get(ctx, sessionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MocksessionGetter)(nil).Get), ctx, sessionID)
}

// MocksessionCreator is a mock of sessionCreator interface.
type MocksessionCreator struct {
	ctrl     *gomock.Controller
	recorder *MocksessionCreatorMockRecorder
	isgomock struct{}
}

// MocksessionCreatorMockRecorder is the mock recorder for MocksessionCreator.
type MocksessionCreatorMockRecorder struct {
	mock *MocksessionCreator
}

// NewMocksessionCreator creates a new mock instance.
func NewMocksessionCreator(ctrl *gomock.Controller) *MocksessionCreator {
	mock := &MocksessionCreator{ctrl: ctrl}
	mock.recorder = &MocksessionCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksessionCreator) EXPECT() *MocksessionCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MocksessionCreator) Create(ctx context.Context, userID uint32) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MocksessionCreatorMockRecorder) Create(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MocksessionCreator)(nil).Create), ctx, userID)
}
