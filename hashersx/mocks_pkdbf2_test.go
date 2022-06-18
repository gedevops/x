// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/x/hashersx (interfaces: PKBDF2Configurator)

// Package hashersx_test is a generated GoMock package.
package hashersx_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	hashersx "github.com/ory/x/hashersx"
)

// MockPKBDF2Configurator is a mock of PKBDF2Configurator interface.
type MockPKBDF2Configurator struct {
	ctrl     *gomock.Controller
	recorder *MockPKBDF2ConfiguratorMockRecorder
}

// MockPKBDF2ConfiguratorMockRecorder is the mock recorder for MockPKBDF2Configurator.
type MockPKBDF2ConfiguratorMockRecorder struct {
	mock *MockPKBDF2Configurator
}

// NewMockPKBDF2Configurator creates a new mock instance.
func NewMockPKBDF2Configurator(ctrl *gomock.Controller) *MockPKBDF2Configurator {
	mock := &MockPKBDF2Configurator{ctrl: ctrl}
	mock.recorder = &MockPKBDF2ConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPKBDF2Configurator) EXPECT() *MockPKBDF2ConfiguratorMockRecorder {
	return m.recorder
}

// HasherPKBDF2Config mocks base method.
func (m *MockPKBDF2Configurator) HasherPKBDF2Config(arg0 context.Context) *hashersx.PKBDF2Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasherPKBDF2Config", arg0)
	ret0, _ := ret[0].(*hashersx.PKBDF2Config)
	return ret0
}

// HasherPKBDF2Config indicates an expected call of HasherPKBDF2Config.
func (mr *MockPKBDF2ConfiguratorMockRecorder) HasherPKBDF2Config(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasherPKBDF2Config", reflect.TypeOf((*MockPKBDF2Configurator)(nil).HasherPKBDF2Config), arg0)
}
