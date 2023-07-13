// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Auth is an autogenerated mock type for the Auth type
type Auth struct {
	mock.Mock
}

// Init provides a mock function with given fields: ctx
func (_m *Auth) Init(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoggedInAs provides a mock function with given fields: r
func (_m *Auth) LoggedInAs(r *http.Request) string {
	ret := _m.Called(r)

	var r0 string
	if rf, ok := ret.Get(0).(func(*http.Request) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LoginURL provides a mock function with given fields: w, r
func (_m *Auth) LoginURL(w http.ResponseWriter, r *http.Request) string {
	ret := _m.Called(w, r)

	var r0 string
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) string); ok {
		r0 = rf(w, r)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewAuth creates a new instance of Auth. It also registers a cleanup function to assert the mocks expectations.
func NewAuth(t testing.TB) *Auth {
	mock := &Auth{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
