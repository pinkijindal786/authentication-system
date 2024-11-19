// Code generated by mockery v2.47.0. DO NOT EDIT.

package services

import mock "github.com/stretchr/testify/mock"

// MockIAuthService is an autogenerated mock type for the IAuthService type
type MockIAuthService struct {
	mock.Mock
}

// RefreshToken provides a mock function with given fields: oldToken
func (_m *MockIAuthService) RefreshToken(oldToken string) (string, error) {
	ret := _m.Called(oldToken)

	if len(ret) == 0 {
		panic("no return value specified for RefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(oldToken)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(oldToken)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(oldToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeToken provides a mock function with given fields: token
func (_m *MockIAuthService) RevokeToken(token string) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for RevokeToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignIn provides a mock function with given fields: email, password
func (_m *MockIAuthService) SignIn(email string, password string) (string, error) {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for SignIn")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: email, password
func (_m *MockIAuthService) SignUp(email string, password string) error {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for SignUp")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockIAuthService creates a new instance of MockIAuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIAuthService {
	mock := &MockIAuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
