// Copyright 2022 PayPal Inc.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this file,
// you can obtain one at https://mit-license.org/.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: drivers/cmd/gcloud-secret/main.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	context "context"
	reflect "reflect"

	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	gomock "github.com/golang/mock/gomock"
	v2 "github.com/googleapis/gax-go/v2"
)

// MockSecretManagerClient is a mock of SecretManagerClient interface.
type MockSecretManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretManagerClientMockRecorder
}

// MockSecretManagerClientMockRecorder is the mock recorder for MockSecretManagerClient.
type MockSecretManagerClientMockRecorder struct {
	mock *MockSecretManagerClient
}

// NewMockSecretManagerClient creates a new mock instance.
func NewMockSecretManagerClient(ctrl *gomock.Controller) *MockSecretManagerClient {
	mock := &MockSecretManagerClient{ctrl: ctrl}
	mock.recorder = &MockSecretManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretManagerClient) EXPECT() *MockSecretManagerClientMockRecorder {
	return m.recorder
}

// AccessSecretVersion mocks base method.
func (m *MockSecretManagerClient) AccessSecretVersion(ctx context.Context, req *secretmanagerpb.AccessSecretVersionRequest, opts ...v2.CallOption) (*secretmanagerpb.AccessSecretVersionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AccessSecretVersion", varargs...)
	ret0, _ := ret[0].(*secretmanagerpb.AccessSecretVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessSecretVersion indicates an expected call of AccessSecretVersion.
func (mr *MockSecretManagerClientMockRecorder) AccessSecretVersion(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessSecretVersion", reflect.TypeOf((*MockSecretManagerClient)(nil).AccessSecretVersion), varargs...)
}

// Close mocks base method.
func (m *MockSecretManagerClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSecretManagerClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSecretManagerClient)(nil).Close))
}
