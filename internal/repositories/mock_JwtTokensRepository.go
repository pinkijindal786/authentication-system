// Code generated by mockery v2.47.0. DO NOT EDIT.

package repositories

import mock "github.com/stretchr/testify/mock"

// MockJwtTokensRepository is an autogenerated mock type for the JwtTokensRepository type
type MockJwtTokensRepository struct {
	mock.Mock
}

// IsTokenRevoked provides a mock function with given fields: token
func (_m *MockJwtTokensRepository) IsTokenRevoked(token string) (bool, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for IsTokenRevoked")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevokeToken provides a mock function with given fields: token
func (_m *MockJwtTokensRepository) RevokeToken(token string) error {
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

// NewMockJwtTokensRepository creates a new instance of MockJwtTokensRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockJwtTokensRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockJwtTokensRepository {
	mock := &MockJwtTokensRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
