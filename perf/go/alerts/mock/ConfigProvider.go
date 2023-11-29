// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	alerts "go.skia.org/infra/perf/go/alerts"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// ConfigProvider is an autogenerated mock type for the ConfigProvider type
type ConfigProvider struct {
	mock.Mock
}

// GetAllAlertConfigs provides a mock function with given fields: ctx, includeDeleted
func (_m *ConfigProvider) GetAllAlertConfigs(ctx context.Context, includeDeleted bool) ([]*alerts.Alert, error) {
	ret := _m.Called(ctx, includeDeleted)

	var r0 []*alerts.Alert
	if rf, ok := ret.Get(0).(func(context.Context, bool) []*alerts.Alert); ok {
		r0 = rf(ctx, includeDeleted)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*alerts.Alert)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool) error); ok {
		r1 = rf(ctx, includeDeleted)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields:
func (_m *ConfigProvider) Refresh() {
	_m.Called()
}

// NewConfigProvider creates a new instance of ConfigProvider. It also registers a cleanup function to assert the mocks expectations.
func NewConfigProvider(t testing.TB) *ConfigProvider {
	mock := &ConfigProvider{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}