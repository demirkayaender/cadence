// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/cadence/common/persistence (interfaces: VisibilityStore)

// Package persistence is a generated GoMock package.
package persistence

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVisibilityStore is a mock of VisibilityStore interface.
type MockVisibilityStore struct {
	ctrl     *gomock.Controller
	recorder *MockVisibilityStoreMockRecorder
}

// MockVisibilityStoreMockRecorder is the mock recorder for MockVisibilityStore.
type MockVisibilityStoreMockRecorder struct {
	mock *MockVisibilityStore
}

// NewMockVisibilityStore creates a new mock instance.
func NewMockVisibilityStore(ctrl *gomock.Controller) *MockVisibilityStore {
	mock := &MockVisibilityStore{ctrl: ctrl}
	mock.recorder = &MockVisibilityStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVisibilityStore) EXPECT() *MockVisibilityStoreMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockVisibilityStore) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockVisibilityStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockVisibilityStore)(nil).Close))
}

// CountWorkflowExecutions mocks base method.
func (m *MockVisibilityStore) CountWorkflowExecutions(arg0 context.Context, arg1 *CountWorkflowExecutionsRequest) (*CountWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*CountWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountWorkflowExecutions indicates an expected call of CountWorkflowExecutions.
func (mr *MockVisibilityStoreMockRecorder) CountWorkflowExecutions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountWorkflowExecutions", reflect.TypeOf((*MockVisibilityStore)(nil).CountWorkflowExecutions), arg0, arg1)
}

// DeleteUninitializedWorkflowExecution mocks base method.
func (m *MockVisibilityStore) DeleteUninitializedWorkflowExecution(arg0 context.Context, arg1 *VisibilityDeleteWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUninitializedWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUninitializedWorkflowExecution indicates an expected call of DeleteUninitializedWorkflowExecution.
func (mr *MockVisibilityStoreMockRecorder) DeleteUninitializedWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUninitializedWorkflowExecution", reflect.TypeOf((*MockVisibilityStore)(nil).DeleteUninitializedWorkflowExecution), arg0, arg1)
}

// DeleteWorkflowExecution mocks base method.
func (m *MockVisibilityStore) DeleteWorkflowExecution(arg0 context.Context, arg1 *VisibilityDeleteWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkflowExecution indicates an expected call of DeleteWorkflowExecution.
func (mr *MockVisibilityStoreMockRecorder) DeleteWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkflowExecution", reflect.TypeOf((*MockVisibilityStore)(nil).DeleteWorkflowExecution), arg0, arg1)
}

// GetClosedWorkflowExecution mocks base method.
func (m *MockVisibilityStore) GetClosedWorkflowExecution(arg0 context.Context, arg1 *InternalGetClosedWorkflowExecutionRequest) (*InternalGetClosedWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClosedWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(*InternalGetClosedWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClosedWorkflowExecution indicates an expected call of GetClosedWorkflowExecution.
func (mr *MockVisibilityStoreMockRecorder) GetClosedWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClosedWorkflowExecution", reflect.TypeOf((*MockVisibilityStore)(nil).GetClosedWorkflowExecution), arg0, arg1)
}

// GetName mocks base method.
func (m *MockVisibilityStore) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockVisibilityStoreMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockVisibilityStore)(nil).GetName))
}

// ListClosedWorkflowExecutions mocks base method.
func (m *MockVisibilityStore) ListClosedWorkflowExecutions(arg0 context.Context, arg1 *InternalListWorkflowExecutionsRequest) (*InternalListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClosedWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*InternalListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClosedWorkflowExecutions indicates an expected call of ListClosedWorkflowExecutions.
func (mr *MockVisibilityStoreMockRecorder) ListClosedWorkflowExecutions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClosedWorkflowExecutions", reflect.TypeOf((*MockVisibilityStore)(nil).ListClosedWorkflowExecutions), arg0, arg1)
}

// ListClosedWorkflowExecutionsByStatus mocks base method.
func (m *MockVisibilityStore) ListClosedWorkflowExecutionsByStatus(arg0 context.Context, arg1 *InternalListClosedWorkflowExecutionsByStatusRequest) (*InternalListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClosedWorkflowExecutionsByStatus", arg0, arg1)
	ret0, _ := ret[0].(*InternalListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClosedWorkflowExecutionsByStatus indicates an expected call of ListClosedWorkflowExecutionsByStatus.
func (mr *MockVisibilityStoreMockRecorder) ListClosedWorkflowExecutionsByStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClosedWorkflowExecutionsByStatus", reflect.TypeOf((*MockVisibilityStore)(nil).ListClosedWorkflowExecutionsByStatus), arg0, arg1)
}

// ListClosedWorkflowExecutionsByType mocks base method.
func (m *MockVisibilityStore) ListClosedWorkflowExecutionsByType(arg0 context.Context, arg1 *InternalListWorkflowExecutionsByTypeRequest) (*InternalListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClosedWorkflowExecutionsByType", arg0, arg1)
	ret0, _ := ret[0].(*InternalListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClosedWorkflowExecutionsByType indicates an expected call of ListClosedWorkflowExecutionsByType.
func (mr *MockVisibilityStoreMockRecorder) ListClosedWorkflowExecutionsByType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClosedWorkflowExecutionsByType", reflect.TypeOf((*MockVisibilityStore)(nil).ListClosedWorkflowExecutionsByType), arg0, arg1)
}

// ListClosedWorkflowExecutionsByWorkflowID mocks base method.
func (m *MockVisibilityStore) ListClosedWorkflowExecutionsByWorkflowID(arg0 context.Context, arg1 *InternalListWorkflowExecutionsByWorkflowIDRequest) (*InternalListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClosedWorkflowExecutionsByWorkflowID", arg0, arg1)
	ret0, _ := ret[0].(*InternalListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClosedWorkflowExecutionsByWorkflowID indicates an expected call of ListClosedWorkflowExecutionsByWorkflowID.
func (mr *MockVisibilityStoreMockRecorder) ListClosedWorkflowExecutionsByWorkflowID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClosedWorkflowExecutionsByWorkflowID", reflect.TypeOf((*MockVisibilityStore)(nil).ListClosedWorkflowExecutionsByWorkflowID), arg0, arg1)
}

// ListOpenWorkflowExecutions mocks base method.
func (m *MockVisibilityStore) ListOpenWorkflowExecutions(arg0 context.Context, arg1 *InternalListWorkflowExecutionsRequest) (*InternalListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOpenWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*InternalListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOpenWorkflowExecutions indicates an expected call of ListOpenWorkflowExecutions.
func (mr *MockVisibilityStoreMockRecorder) ListOpenWorkflowExecutions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenWorkflowExecutions", reflect.TypeOf((*MockVisibilityStore)(nil).ListOpenWorkflowExecutions), arg0, arg1)
}

// ListOpenWorkflowExecutionsByType mocks base method.
func (m *MockVisibilityStore) ListOpenWorkflowExecutionsByType(arg0 context.Context, arg1 *InternalListWorkflowExecutionsByTypeRequest) (*InternalListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOpenWorkflowExecutionsByType", arg0, arg1)
	ret0, _ := ret[0].(*InternalListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOpenWorkflowExecutionsByType indicates an expected call of ListOpenWorkflowExecutionsByType.
func (mr *MockVisibilityStoreMockRecorder) ListOpenWorkflowExecutionsByType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenWorkflowExecutionsByType", reflect.TypeOf((*MockVisibilityStore)(nil).ListOpenWorkflowExecutionsByType), arg0, arg1)
}

// ListOpenWorkflowExecutionsByWorkflowID mocks base method.
func (m *MockVisibilityStore) ListOpenWorkflowExecutionsByWorkflowID(arg0 context.Context, arg1 *InternalListWorkflowExecutionsByWorkflowIDRequest) (*InternalListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOpenWorkflowExecutionsByWorkflowID", arg0, arg1)
	ret0, _ := ret[0].(*InternalListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOpenWorkflowExecutionsByWorkflowID indicates an expected call of ListOpenWorkflowExecutionsByWorkflowID.
func (mr *MockVisibilityStoreMockRecorder) ListOpenWorkflowExecutionsByWorkflowID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenWorkflowExecutionsByWorkflowID", reflect.TypeOf((*MockVisibilityStore)(nil).ListOpenWorkflowExecutionsByWorkflowID), arg0, arg1)
}

// ListWorkflowExecutions mocks base method.
func (m *MockVisibilityStore) ListWorkflowExecutions(arg0 context.Context, arg1 *ListWorkflowExecutionsByQueryRequest) (*InternalListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*InternalListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkflowExecutions indicates an expected call of ListWorkflowExecutions.
func (mr *MockVisibilityStoreMockRecorder) ListWorkflowExecutions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflowExecutions", reflect.TypeOf((*MockVisibilityStore)(nil).ListWorkflowExecutions), arg0, arg1)
}

// RecordWorkflowExecutionClosed mocks base method.
func (m *MockVisibilityStore) RecordWorkflowExecutionClosed(arg0 context.Context, arg1 *InternalRecordWorkflowExecutionClosedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordWorkflowExecutionClosed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordWorkflowExecutionClosed indicates an expected call of RecordWorkflowExecutionClosed.
func (mr *MockVisibilityStoreMockRecorder) RecordWorkflowExecutionClosed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordWorkflowExecutionClosed", reflect.TypeOf((*MockVisibilityStore)(nil).RecordWorkflowExecutionClosed), arg0, arg1)
}

// RecordWorkflowExecutionStarted mocks base method.
func (m *MockVisibilityStore) RecordWorkflowExecutionStarted(arg0 context.Context, arg1 *InternalRecordWorkflowExecutionStartedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordWorkflowExecutionStarted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordWorkflowExecutionStarted indicates an expected call of RecordWorkflowExecutionStarted.
func (mr *MockVisibilityStoreMockRecorder) RecordWorkflowExecutionStarted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordWorkflowExecutionStarted", reflect.TypeOf((*MockVisibilityStore)(nil).RecordWorkflowExecutionStarted), arg0, arg1)
}

// RecordWorkflowExecutionUninitialized mocks base method.
func (m *MockVisibilityStore) RecordWorkflowExecutionUninitialized(arg0 context.Context, arg1 *InternalRecordWorkflowExecutionUninitializedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordWorkflowExecutionUninitialized", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordWorkflowExecutionUninitialized indicates an expected call of RecordWorkflowExecutionUninitialized.
func (mr *MockVisibilityStoreMockRecorder) RecordWorkflowExecutionUninitialized(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordWorkflowExecutionUninitialized", reflect.TypeOf((*MockVisibilityStore)(nil).RecordWorkflowExecutionUninitialized), arg0, arg1)
}

// ScanWorkflowExecutions mocks base method.
func (m *MockVisibilityStore) ScanWorkflowExecutions(arg0 context.Context, arg1 *ListWorkflowExecutionsByQueryRequest) (*InternalListWorkflowExecutionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanWorkflowExecutions", arg0, arg1)
	ret0, _ := ret[0].(*InternalListWorkflowExecutionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScanWorkflowExecutions indicates an expected call of ScanWorkflowExecutions.
func (mr *MockVisibilityStoreMockRecorder) ScanWorkflowExecutions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanWorkflowExecutions", reflect.TypeOf((*MockVisibilityStore)(nil).ScanWorkflowExecutions), arg0, arg1)
}

// UpsertWorkflowExecution mocks base method.
func (m *MockVisibilityStore) UpsertWorkflowExecution(arg0 context.Context, arg1 *InternalUpsertWorkflowExecutionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertWorkflowExecution", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertWorkflowExecution indicates an expected call of UpsertWorkflowExecution.
func (mr *MockVisibilityStoreMockRecorder) UpsertWorkflowExecution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertWorkflowExecution", reflect.TypeOf((*MockVisibilityStore)(nil).UpsertWorkflowExecution), arg0, arg1)
}