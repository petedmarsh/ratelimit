// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/envoyproxy/ratelimit/src/memcached (interfaces: Client)

// Package mock_memcached is a generated GoMock package.
package mock_memcached

import (
	memcache "github.com/rainycape/memcache"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockClient) Add(arg0 *memcache.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockClientMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockClient)(nil).Add), arg0)
}

// GetMulti mocks base method
func (m *MockClient) GetMulti(arg0 []string) (map[string]*memcache.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMulti", arg0)
	ret0, _ := ret[0].(map[string]*memcache.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMulti indicates an expected call of GetMulti
func (mr *MockClientMockRecorder) GetMulti(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockClient)(nil).GetMulti), arg0)
}

// Increment mocks base method
func (m *MockClient) Increment(arg0 string, arg1 uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Increment", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Increment indicates an expected call of Increment
func (mr *MockClientMockRecorder) Increment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockClient)(nil).Increment), arg0, arg1)
}
