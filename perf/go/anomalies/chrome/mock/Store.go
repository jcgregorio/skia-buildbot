// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	chromeperf "go.skia.org/infra/perf/go/chromeperf"

	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// GetAnomalies provides a mock function with given fields: ctx, traceNames, startCommitPosition, endCommitPosition
func (_m *Store) GetAnomalies(ctx context.Context, traceNames []string, startCommitPosition int, endCommitPosition int) (chromeperf.AnomalyMap, error) {
	ret := _m.Called(ctx, traceNames, startCommitPosition, endCommitPosition)

	if len(ret) == 0 {
		panic("no return value specified for GetAnomalies")
	}

	var r0 chromeperf.AnomalyMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) (chromeperf.AnomalyMap, error)); ok {
		return rf(ctx, traceNames, startCommitPosition, endCommitPosition)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) chromeperf.AnomalyMap); ok {
		r0 = rf(ctx, traceNames, startCommitPosition, endCommitPosition)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chromeperf.AnomalyMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, int, int) error); ok {
		r1 = rf(ctx, traceNames, startCommitPosition, endCommitPosition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAnomaliesAroundRevision provides a mock function with given fields: ctx, revision
func (_m *Store) GetAnomaliesAroundRevision(ctx context.Context, revision int) ([]chromeperf.AnomalyForRevision, error) {
	ret := _m.Called(ctx, revision)

	if len(ret) == 0 {
		panic("no return value specified for GetAnomaliesAroundRevision")
	}

	var r0 []chromeperf.AnomalyForRevision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]chromeperf.AnomalyForRevision, error)); ok {
		return rf(ctx, revision)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []chromeperf.AnomalyForRevision); ok {
		r0 = rf(ctx, revision)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chromeperf.AnomalyForRevision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, revision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
