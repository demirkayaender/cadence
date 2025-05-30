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
// Source: persister.go
//
// Generated by this command:
//
//	mockgen -package types -source persister.go -destination persister_mock.go -package types github.com/uber/cadence/common/mapq/types Persister
//

// Package types is a generated GoMock package.
package types

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPersister is a mock of Persister interface.
type MockPersister struct {
	ctrl     *gomock.Controller
	recorder *MockPersisterMockRecorder
	isgomock struct{}
}

// MockPersisterMockRecorder is the mock recorder for MockPersister.
type MockPersisterMockRecorder struct {
	mock *MockPersister
}

// NewMockPersister creates a new mock instance.
func NewMockPersister(ctrl *gomock.Controller) *MockPersister {
	mock := &MockPersister{ctrl: ctrl}
	mock.recorder = &MockPersisterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersister) EXPECT() *MockPersisterMockRecorder {
	return m.recorder
}

// CommitOffsets mocks base method.
func (m *MockPersister) CommitOffsets(ctx context.Context, offsets *Offsets) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitOffsets", ctx, offsets)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitOffsets indicates an expected call of CommitOffsets.
func (mr *MockPersisterMockRecorder) CommitOffsets(ctx, offsets any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitOffsets", reflect.TypeOf((*MockPersister)(nil).CommitOffsets), ctx, offsets)
}

// Fetch mocks base method.
func (m *MockPersister) Fetch(ctx context.Context, partitions ItemPartitions, pageInfo PageInfo) ([]Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, partitions, pageInfo)
	ret0, _ := ret[0].([]Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockPersisterMockRecorder) Fetch(ctx, partitions, pageInfo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockPersister)(nil).Fetch), ctx, partitions, pageInfo)
}

// GetOffsets mocks base method.
func (m *MockPersister) GetOffsets(ctx context.Context) (*Offsets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffsets", ctx)
	ret0, _ := ret[0].(*Offsets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffsets indicates an expected call of GetOffsets.
func (mr *MockPersisterMockRecorder) GetOffsets(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffsets", reflect.TypeOf((*MockPersister)(nil).GetOffsets), ctx)
}

// Persist mocks base method.
func (m *MockPersister) Persist(ctx context.Context, items []ItemToPersist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Persist", ctx, items)
	ret0, _ := ret[0].(error)
	return ret0
}

// Persist indicates an expected call of Persist.
func (mr *MockPersisterMockRecorder) Persist(ctx, items any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Persist", reflect.TypeOf((*MockPersister)(nil).Persist), ctx, items)
}
