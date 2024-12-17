// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockBlobClient is an autogenerated mock type for the blobClient type
type MockBlobClient struct {
	mock.Mock
}

type MockBlobClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBlobClient) EXPECT() *MockBlobClient_Expecter {
	return &MockBlobClient_Expecter{mock: &_m.Mock}
}

// UploadBuffer provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockBlobClient) UploadBuffer(_a0 context.Context, _a1 string, _a2 string, _a3 []byte) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for UploadBuffer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []byte) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBlobClient_UploadBuffer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadBuffer'
type MockBlobClient_UploadBuffer_Call struct {
	*mock.Call
}

// UploadBuffer is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
//   - _a2 string
//   - _a3 []byte
func (_e *MockBlobClient_Expecter) UploadBuffer(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}) *MockBlobClient_UploadBuffer_Call {
	return &MockBlobClient_UploadBuffer_Call{Call: _e.mock.On("UploadBuffer", _a0, _a1, _a2, _a3)}
}

func (_c *MockBlobClient_UploadBuffer_Call) Run(run func(_a0 context.Context, _a1 string, _a2 string, _a3 []byte)) *MockBlobClient_UploadBuffer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].([]byte))
	})
	return _c
}

func (_c *MockBlobClient_UploadBuffer_Call) Return(_a0 error) *MockBlobClient_UploadBuffer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBlobClient_UploadBuffer_Call) RunAndReturn(run func(context.Context, string, string, []byte) error) *MockBlobClient_UploadBuffer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBlobClient creates a new instance of MockBlobClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBlobClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBlobClient {
	mock := &MockBlobClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
