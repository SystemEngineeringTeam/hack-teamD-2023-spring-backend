// Code generated by MockGen. DO NOT EDIT.
// Source: chat_completions.go

// Package mock_chat is a generated GoMock package.
package mock_chat

import (
	reflect "reflect"
	chat "suggest-be/src/chat"

	gomock "github.com/golang/mock/gomock"
)

// MockCompletion is a mock of Completion interface.
type MockCompletion struct {
	ctrl     *gomock.Controller
	recorder *MockCompletionMockRecorder
}

// MockCompletionMockRecorder is the mock recorder for MockCompletion.
type MockCompletionMockRecorder struct {
	mock *MockCompletion
}

// NewMockCompletion creates a new mock instance.
func NewMockCompletion(ctrl *gomock.Controller) *MockCompletion {
	mock := &MockCompletion{ctrl: ctrl}
	mock.recorder = &MockCompletionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompletion) EXPECT() *MockCompletionMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockCompletion) SendMessage(arg0 []*chat.RequestMessage) (*chat.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0)
	ret0, _ := ret[0].(*chat.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockCompletionMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockCompletion)(nil).SendMessage), arg0)
}
