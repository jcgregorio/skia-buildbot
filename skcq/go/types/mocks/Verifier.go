// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	gerrit "go.skia.org/infra/go/gerrit"

	types "go.skia.org/infra/skcq/go/types"
)

// Verifier is an autogenerated mock type for the Verifier type
type Verifier struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: ctx, ci, cleanupPatchsetID
func (_m *Verifier) Cleanup(ctx context.Context, ci *gerrit.ChangeInfo, cleanupPatchsetID int64) {
	_m.Called(ctx, ci, cleanupPatchsetID)
}

// Name provides a mock function with given fields:
func (_m *Verifier) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Verify provides a mock function with given fields: ctx, ci, startTime
func (_m *Verifier) Verify(ctx context.Context, ci *gerrit.ChangeInfo, startTime int64) (types.VerifierState, string, error) {
	ret := _m.Called(ctx, ci, startTime)

	var r0 types.VerifierState
	if rf, ok := ret.Get(0).(func(context.Context, *gerrit.ChangeInfo, int64) types.VerifierState); ok {
		r0 = rf(ctx, ci, startTime)
	} else {
		r0 = ret.Get(0).(types.VerifierState)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, *gerrit.ChangeInfo, int64) string); ok {
		r1 = rf(ctx, ci, startTime)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *gerrit.ChangeInfo, int64) error); ok {
		r2 = rf(ctx, ci, startTime)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}