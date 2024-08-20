// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	context "context"

	protobufs "github.com/open-telemetry/opamp-go/protobufs"
	mock "github.com/stretchr/testify/mock"
)

// MockMonitor is an autogenerated mock type for the Monitor type
type MockMonitor struct {
	mock.Mock
}

// MonitorForSuccess provides a mock function with given fields: ctx, packageName
func (_m *MockMonitor) MonitorForSuccess(ctx context.Context, packageName string) error {
	ret := _m.Called(ctx, packageName)

	if len(ret) == 0 {
		panic("no return value specified for MonitorForSuccess")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, packageName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetState provides a mock function with given fields: packageName, status, statusErr
func (_m *MockMonitor) SetState(packageName string, status protobufs.PackageStatusEnum, statusErr error) error {
	ret := _m.Called(packageName, status, statusErr)

	if len(ret) == 0 {
		panic("no return value specified for SetState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, protobufs.PackageStatusEnum, error) error); ok {
		r0 = rf(packageName, status, statusErr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockMonitor creates a new instance of MockMonitor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMonitor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMonitor {
	mock := &MockMonitor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
