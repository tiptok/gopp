// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Stringer is an autogenerated mock type for the Stringer type
type Stringer struct {
	mock.Mock
}

// String provides a mock function with given fields:
func (_m *Stringer) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewStringer interface {
	mock.TestingT
	Cleanup(func())
}

// NewStringer creates a new instance of Stringer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStringer(t mockConstructorTestingTNewStringer) *Stringer {
	mock := &Stringer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
