// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import time "time"

// TaskInterface is an autogenerated mock type for the TaskInterface type
type TaskInterface struct {
	mock.Mock
}

// Execute provides a mock function with given fields: instance
func (_m *TaskInterface) Execute(instance time.Time) (bool, error) {
	ret := _m.Called(instance)

	var r0 bool
	if rf, ok := ret.Get(0).(func(time.Time) bool); ok {
		r0 = rf(instance)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time) error); ok {
		r1 = rf(instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
