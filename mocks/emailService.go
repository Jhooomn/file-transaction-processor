// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// EmailService is an autogenerated mock type for the EmailService type
type EmailService struct {
	mock.Mock
}

// Send provides a mock function with given fields: ctx, recipient, subject, body
func (_m *EmailService) Send(ctx context.Context, recipient string, subject string, body string) error {
	ret := _m.Called(ctx, recipient, subject, body)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, recipient, subject, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEmailService creates a new instance of EmailService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEmailService(t interface {
	mock.TestingT
	Cleanup(func())
}) *EmailService {
	mock := &EmailService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
