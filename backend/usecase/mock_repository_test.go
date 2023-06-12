// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repository.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"
	time "time"

	auth "firebase.google.com/go/v4/auth"
	gomock "github.com/golang/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// UpsertUser mocks base method.
func (m *MockDatabase) UpsertUser(ctx context.Context, name, email, pictureUrl string, updatedAt time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertUser", ctx, name, email, pictureUrl, updatedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertUser indicates an expected call of UpsertUser.
func (mr *MockDatabaseMockRecorder) UpsertUser(ctx, name, email, pictureUrl, updatedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertUser", reflect.TypeOf((*MockDatabase)(nil).UpsertUser), ctx, name, email, pictureUrl, updatedAt)
}

// MockFirebase is a mock of Firebase interface.
type MockFirebase struct {
	ctrl     *gomock.Controller
	recorder *MockFirebaseMockRecorder
}

// MockFirebaseMockRecorder is the mock recorder for MockFirebase.
type MockFirebaseMockRecorder struct {
	mock *MockFirebase
}

// NewMockFirebase creates a new mock instance.
func NewMockFirebase(ctrl *gomock.Controller) *MockFirebase {
	mock := &MockFirebase{ctrl: ctrl}
	mock.recorder = &MockFirebaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirebase) EXPECT() *MockFirebaseMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockFirebase) CreateSession(ctx context.Context, idToken string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, idToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockFirebaseMockRecorder) CreateSession(ctx, idToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockFirebase)(nil).CreateSession), ctx, idToken)
}

// VerifyIDToken mocks base method.
func (m *MockFirebase) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIDToken", ctx, idToken)
	ret0, _ := ret[0].(*auth.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyIDToken indicates an expected call of VerifyIDToken.
func (mr *MockFirebaseMockRecorder) VerifyIDToken(ctx, idToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIDToken", reflect.TypeOf((*MockFirebase)(nil).VerifyIDToken), ctx, idToken)
}

// VerifySessionCookie mocks base method.
func (m *MockFirebase) VerifySessionCookie(ctx context.Context, cookie string) (*auth.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySessionCookie", ctx, cookie)
	ret0, _ := ret[0].(*auth.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifySessionCookie indicates an expected call of VerifySessionCookie.
func (mr *MockFirebaseMockRecorder) VerifySessionCookie(ctx, cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySessionCookie", reflect.TypeOf((*MockFirebase)(nil).VerifySessionCookie), ctx, cookie)
}
