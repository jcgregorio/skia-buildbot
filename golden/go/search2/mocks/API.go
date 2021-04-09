// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	search2 "go.skia.org/infra/golden/go/search2"

	time "time"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

// ChangelistLastUpdated provides a mock function with given fields: ctx, crs, clID
func (_m *API) ChangelistLastUpdated(ctx context.Context, crs string, clID string) (time.Time, error) {
	ret := _m.Called(ctx, crs, clID)

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(context.Context, string, string) time.Time); ok {
		r0 = rf(ctx, crs, clID)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, crs, clID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAndUntriagedSummaryForCL provides a mock function with given fields: ctx, crs, clID
func (_m *API) NewAndUntriagedSummaryForCL(ctx context.Context, crs string, clID string) (search2.NewAndUntriagedSummary, error) {
	ret := _m.Called(ctx, crs, clID)

	var r0 search2.NewAndUntriagedSummary
	if rf, ok := ret.Get(0).(func(context.Context, string, string) search2.NewAndUntriagedSummary); ok {
		r0 = rf(ctx, crs, clID)
	} else {
		r0 = ret.Get(0).(search2.NewAndUntriagedSummary)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, crs, clID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}