// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GenesisEducationKyiv/main-project-delveper/internal/subscription (interfaces: RateGetter)

// Package subscription is a generated GoMock package.
package subscription

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRateGetter is a mock of RateGetter interface.
type MockRateGetter struct {
	ctrl     *gomock.Controller
	recorder *MockRateGetterMockRecorder
}

// MockRateGetterMockRecorder is the mock recorder for MockRateGetter.
type MockRateGetterMockRecorder struct {
	mock *MockRateGetter
}

// NewMockRateGetter creates a new mock instance.
func NewMockRateGetter(ctrl *gomock.Controller) *MockRateGetter {
	mock := &MockRateGetter{ctrl: ctrl}
	mock.recorder = &MockRateGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateGetter) EXPECT() *MockRateGetterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRateGetter) Get() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRateGetterMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRateGetter)(nil).Get))
}