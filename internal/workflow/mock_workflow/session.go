// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: internal/workflow/session.go

// Package mock_workflow is a generated GoMock package.
package mock_workflow

import (
	gomock "github.com/golang/mock/gomock"
	dipper "github.com/honeydipper/honeydipper/pkg/dipper"
	reflect "reflect"
	time "time"
)

// MockSessionHandler is a mock of SessionHandler interface
type MockSessionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSessionHandlerMockRecorder
}

// MockSessionHandlerMockRecorder is the mock recorder for MockSessionHandler
type MockSessionHandlerMockRecorder struct {
	mock *MockSessionHandler
}

// NewMockSessionHandler creates a new mock instance
func NewMockSessionHandler(ctrl *gomock.Controller) *MockSessionHandler {
	mock := &MockSessionHandler{ctrl: ctrl}
	mock.recorder = &MockSessionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionHandler) EXPECT() *MockSessionHandlerMockRecorder {
	return m.recorder
}

// prepare mocks base method
func (m *MockSessionHandler) prepare(msg *dipper.Message, parent interface{}, ctx map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "prepare", msg, parent, ctx)
}

// prepare indicates an expected call of prepare
func (mr *MockSessionHandlerMockRecorder) prepare(msg, parent, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "prepare", reflect.TypeOf((*MockSessionHandler)(nil).prepare), msg, parent, ctx)
}

// execute mocks base method
func (m *MockSessionHandler) execute(msg *dipper.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "execute", msg)
}

// execute indicates an expected call of execute
func (mr *MockSessionHandlerMockRecorder) execute(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "execute", reflect.TypeOf((*MockSessionHandler)(nil).execute), msg)
}

// continueExec mocks base method
func (m *MockSessionHandler) continueExec(msg *dipper.Message, exports []map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "continueExec", msg, exports)
}

// continueExec indicates an expected call of continueExec
func (mr *MockSessionHandlerMockRecorder) continueExec(msg, exports interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "continueExec", reflect.TypeOf((*MockSessionHandler)(nil).continueExec), msg, exports)
}

// onError mocks base method
func (m *MockSessionHandler) onError() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "onError")
}

// onError indicates an expected call of onError
func (mr *MockSessionHandlerMockRecorder) onError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "onError", reflect.TypeOf((*MockSessionHandler)(nil).onError))
}

// GetName mocks base method
func (m *MockSessionHandler) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockSessionHandlerMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockSessionHandler)(nil).GetName))
}

// GetDescription mocks base method
func (m *MockSessionHandler) GetDescription() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDescription")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDescription indicates an expected call of GetDescription
func (mr *MockSessionHandlerMockRecorder) GetDescription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDescription", reflect.TypeOf((*MockSessionHandler)(nil).GetDescription))
}

// GetParent mocks base method
func (m *MockSessionHandler) GetParent() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParent")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetParent indicates an expected call of GetParent
func (mr *MockSessionHandlerMockRecorder) GetParent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParent", reflect.TypeOf((*MockSessionHandler)(nil).GetParent))
}

// GetEventID mocks base method
func (m *MockSessionHandler) GetEventID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEventID indicates an expected call of GetEventID
func (mr *MockSessionHandlerMockRecorder) GetEventID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventID", reflect.TypeOf((*MockSessionHandler)(nil).GetEventID))
}

// GetEventName mocks base method
func (m *MockSessionHandler) GetEventName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEventName indicates an expected call of GetEventName
func (mr *MockSessionHandlerMockRecorder) GetEventName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventName", reflect.TypeOf((*MockSessionHandler)(nil).GetEventName))
}

// GetStatus mocks base method
func (m *MockSessionHandler) GetStatus() (string, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockSessionHandlerMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockSessionHandler)(nil).GetStatus))
}

// GetExported mocks base method
func (m *MockSessionHandler) GetExported() []map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExported")
	ret0, _ := ret[0].([]map[string]interface{})
	return ret0
}

// GetExported indicates an expected call of GetExported
func (mr *MockSessionHandlerMockRecorder) GetExported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExported", reflect.TypeOf((*MockSessionHandler)(nil).GetExported))
}

// Watch mocks base method
func (m *MockSessionHandler) Watch() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Watch indicates an expected call of Watch
func (mr *MockSessionHandlerMockRecorder) Watch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockSessionHandler)(nil).Watch))
}

// GetStartTime mocks base method
func (m *MockSessionHandler) GetStartTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStartTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetStartTime indicates an expected call of GetStartTime
func (mr *MockSessionHandlerMockRecorder) GetStartTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStartTime", reflect.TypeOf((*MockSessionHandler)(nil).GetStartTime))
}

// GetCompletionTime mocks base method
func (m *MockSessionHandler) GetCompletionTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompletionTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetCompletionTime indicates an expected call of GetCompletionTime
func (mr *MockSessionHandlerMockRecorder) GetCompletionTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompletionTime", reflect.TypeOf((*MockSessionHandler)(nil).GetCompletionTime))
}
