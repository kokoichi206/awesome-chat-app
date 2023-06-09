// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/usecase.go

// Package handler_test is a generated GoMock package.
package handler_test

import (
	context "context"
	net "net"
	reflect "reflect"
	time "time"

	auth "firebase.google.com/go/v4/auth"
	gomock "github.com/golang/mock/gomock"
	model "github.com/kokoichi206/awesome-chat-app/backend/model"
	response "github.com/kokoichi206/awesome-chat-app/backend/model/response"
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

// GetMessages mocks base method.
func (m *MockUsecase) GetMessages(ctx context.Context, roomID string, lastReadAt time.Time) ([]*response.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessages", ctx, roomID, lastReadAt)
	ret0, _ := ret[0].([]*response.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessages indicates an expected call of GetMessages.
func (mr *MockUsecaseMockRecorder) GetMessages(ctx, roomID, lastReadAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockUsecase)(nil).GetMessages), ctx, roomID, lastReadAt)
}

// GetRoomUsers mocks base method.
func (m *MockUsecase) GetRoomUsers(ctx context.Context, roomID string) ([]*response.RoomUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomUsers", ctx, roomID)
	ret0, _ := ret[0].([]*response.RoomUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomUsers indicates an expected call of GetRoomUsers.
func (mr *MockUsecaseMockRecorder) GetRoomUsers(ctx, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomUsers", reflect.TypeOf((*MockUsecase)(nil).GetRoomUsers), ctx, roomID)
}

// GetUser mocks base method.
func (m *MockUsecase) GetUser(ctx context.Context, email string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUsecaseMockRecorder) GetUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUsecase)(nil).GetUser), ctx, email)
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

// PostMessage mocks base method.
func (m *MockUsecase) PostMessage(ctx context.Context, roomID, userID, content string, messageType model.MessageType, postedAt time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostMessage", ctx, roomID, userID, content, messageType, postedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostMessage indicates an expected call of PostMessage.
func (mr *MockUsecaseMockRecorder) PostMessage(ctx, roomID, userID, content, messageType, postedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessage", reflect.TypeOf((*MockUsecase)(nil).PostMessage), ctx, roomID, userID, content, messageType, postedAt)
}

// SubscribeMessages mocks base method.
func (m *MockUsecase) SubscribeMessages(ctx context.Context, conn *net.Conn, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeMessages", ctx, conn, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeMessages indicates an expected call of SubscribeMessages.
func (mr *MockUsecaseMockRecorder) SubscribeMessages(ctx, conn, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeMessages", reflect.TypeOf((*MockUsecase)(nil).SubscribeMessages), ctx, conn, email)
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
func (m *MockUsecase) VerifySessionCookie(ctx context.Context, session string) (*auth.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySessionCookie", ctx, session)
	ret0, _ := ret[0].(*auth.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifySessionCookie indicates an expected call of VerifySessionCookie.
func (mr *MockUsecaseMockRecorder) VerifySessionCookie(ctx, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySessionCookie", reflect.TypeOf((*MockUsecase)(nil).VerifySessionCookie), ctx, session)
}
