// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/usecase.go

// Package handler_test is a generated GoMock package.
package handler_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// PostLogin mocks base method.
func (m *MockUsecase) PostLogin(ctx context.Context, token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostLogin", ctx, token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostLogin indicates an expected call of PostLogin.
func (mr *MockUsecaseMockRecorder) PostLogin(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostLogin", reflect.TypeOf((*MockUsecase)(nil).PostLogin), ctx, token)
}

// VerifyIDToken mocks base method.
func (m *MockUsecase) VerifyIDToken(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIDToken", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyIDToken indicates an expected call of VerifyIDToken.
func (mr *MockUsecaseMockRecorder) VerifyIDToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIDToken", reflect.TypeOf((*MockUsecase)(nil).VerifyIDToken), ctx, token)
}

// VerifySessionCookie mocks base method.
func (m *MockUsecase) VerifySessionCookie(ctx context.Context, session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySessionCookie", ctx, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifySessionCookie indicates an expected call of VerifySessionCookie.
func (mr *MockUsecaseMockRecorder) VerifySessionCookie(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySessionCookie", reflect.TypeOf((*MockUsecase)(nil).VerifySessionCookie), ctx, session)
}
