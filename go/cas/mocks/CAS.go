// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CAS is an autogenerated mock type for the CAS type
type CAS struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *CAS) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Download provides a mock function with given fields: ctx, root, digest
func (_m *CAS) Download(ctx context.Context, root string, digest string) error {
	ret := _m.Called(ctx, root, digest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, root, digest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Merge provides a mock function with given fields: ctx, digests
func (_m *CAS) Merge(ctx context.Context, digests []string) (string, error) {
	ret := _m.Called(ctx, digests)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, []string) string); ok {
		r0 = rf(ctx, digests)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, digests)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upload provides a mock function with given fields: ctx, root, paths, excludes
func (_m *CAS) Upload(ctx context.Context, root string, paths []string, excludes []string) (string, error) {
	ret := _m.Called(ctx, root, paths, excludes)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, []string) string); ok {
		r0 = rf(ctx, root, paths, excludes)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, []string) error); ok {
		r1 = rf(ctx, root, paths, excludes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}