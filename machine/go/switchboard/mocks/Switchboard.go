// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	switchboard "go.skia.org/infra/machine/go/switchboard"
)

// Switchboard is an autogenerated mock type for the Switchboard type
type Switchboard struct {
	mock.Mock
}

// AddPod provides a mock function with given fields: ctx, PodName
func (_m *Switchboard) AddPod(ctx context.Context, PodName string) error {
	ret := _m.Called(ctx, PodName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, PodName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearMeetingPoint provides a mock function with given fields: ctx, meeingPoint
func (_m *Switchboard) ClearMeetingPoint(ctx context.Context, meeingPoint switchboard.MeetingPoint) error {
	ret := _m.Called(ctx, meeingPoint)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, switchboard.MeetingPoint) error); ok {
		r0 = rf(ctx, meeingPoint)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMeetingPoint provides a mock function with given fields: ctx, machineID
func (_m *Switchboard) GetMeetingPoint(ctx context.Context, machineID string) (switchboard.MeetingPoint, error) {
	ret := _m.Called(ctx, machineID)

	var r0 switchboard.MeetingPoint
	if rf, ok := ret.Get(0).(func(context.Context, string) switchboard.MeetingPoint); ok {
		r0 = rf(ctx, machineID)
	} else {
		r0 = ret.Get(0).(switchboard.MeetingPoint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, machineID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeepAliveMeetingPoint provides a mock function with given fields: ctx, meetingPoint
func (_m *Switchboard) KeepAliveMeetingPoint(ctx context.Context, meetingPoint switchboard.MeetingPoint) error {
	ret := _m.Called(ctx, meetingPoint)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, switchboard.MeetingPoint) error); ok {
		r0 = rf(ctx, meetingPoint)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListMeetingPoints provides a mock function with given fields: ctx
func (_m *Switchboard) ListMeetingPoints(ctx context.Context) ([]switchboard.MeetingPoint, error) {
	ret := _m.Called(ctx)

	var r0 []switchboard.MeetingPoint
	if rf, ok := ret.Get(0).(func(context.Context) []switchboard.MeetingPoint); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]switchboard.MeetingPoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPods provides a mock function with given fields: ctx
func (_m *Switchboard) ListPods(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemovePod provides a mock function with given fields: ctx, PodName
func (_m *Switchboard) RemovePod(ctx context.Context, PodName string) error {
	ret := _m.Called(ctx, PodName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, PodName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReserveMeetingPoint provides a mock function with given fields: ctx, machineID, Username
func (_m *Switchboard) ReserveMeetingPoint(ctx context.Context, machineID string, Username string) (switchboard.MeetingPoint, error) {
	ret := _m.Called(ctx, machineID, Username)

	var r0 switchboard.MeetingPoint
	if rf, ok := ret.Get(0).(func(context.Context, string, string) switchboard.MeetingPoint); ok {
		r0 = rf(ctx, machineID, Username)
	} else {
		r0 = ret.Get(0).(switchboard.MeetingPoint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, machineID, Username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}