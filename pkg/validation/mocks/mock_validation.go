// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dhuki/go-date/pkg/validation (interfaces: Validation)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidation is a mock of Validation interface.
type MockValidation struct {
	ctrl     *gomock.Controller
	recorder *MockValidationMockRecorder
}

// MockValidationMockRecorder is the mock recorder for MockValidation.
type MockValidationMockRecorder struct {
	mock *MockValidation
}

// NewMockValidation creates a new mock instance.
func NewMockValidation(ctrl *gomock.Controller) *MockValidation {
	mock := &MockValidation{ctrl: ctrl}
	mock.recorder = &MockValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidation) EXPECT() *MockValidationMockRecorder {
	return m.recorder
}

// GenerateJWTAccessToken mocks base method.
func (m *MockValidation) GenerateJWTAccessToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateJWTAccessToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateJWTAccessToken indicates an expected call of GenerateJWTAccessToken.
func (mr *MockValidationMockRecorder) GenerateJWTAccessToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateJWTAccessToken", reflect.TypeOf((*MockValidation)(nil).GenerateJWTAccessToken))
}
