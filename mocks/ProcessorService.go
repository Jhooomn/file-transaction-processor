// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ProcessorService is an autogenerated mock type for the ProcessorService type
type ProcessorService struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *ProcessorService) Execute() {
	_m.Called()
}

// NewProcessorService creates a new instance of ProcessorService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProcessorService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProcessorService {
	mock := &ProcessorService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
