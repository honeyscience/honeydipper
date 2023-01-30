// Copyright 2022 PayPal Inc.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this file,
// you can obtain one at https://mit-license.org/.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: internal/workflow/store.go

// Package mock_workflow is a generated GoMock package.
package mock_workflow

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	config "github.com/honeydipper/honeydipper/internal/config"
	dipper "github.com/honeydipper/honeydipper/pkg/dipper"
)

// MockSessionStoreHelper is a mock of SessionStoreHelper interface.
type MockSessionStoreHelper struct {
	ctrl     *gomock.Controller
	recorder *MockSessionStoreHelperMockRecorder
}

// MockSessionStoreHelperMockRecorder is the mock recorder for MockSessionStoreHelper.
type MockSessionStoreHelperMockRecorder struct {
	mock *MockSessionStoreHelper
}

// NewMockSessionStoreHelper creates a new mock instance.
func NewMockSessionStoreHelper(ctrl *gomock.Controller) *MockSessionStoreHelper {
	mock := &MockSessionStoreHelper{ctrl: ctrl}
	mock.recorder = &MockSessionStoreHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionStoreHelper) EXPECT() *MockSessionStoreHelperMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockSessionStoreHelper) GetConfig() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockSessionStoreHelperMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockSessionStoreHelper)(nil).GetConfig))
}

// SendMessage mocks base method.
func (m *MockSessionStoreHelper) SendMessage(msg *dipper.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendMessage", msg)
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockSessionStoreHelperMockRecorder) SendMessage(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockSessionStoreHelper)(nil).SendMessage), msg)
}
