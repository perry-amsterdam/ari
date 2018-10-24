// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"

// Recorder is an autogenerated mock type for the Recorder type
type Recorder struct {
	mock.Mock
}

// Record provides a mock function with given fields: _a0, _a1
func (_m *Recorder) Record(_a0 string, _a1 *ari.RecordingOptions) (*ari.LiveRecordingHandle, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ari.LiveRecordingHandle
	if rf, ok := ret.Get(0).(func(string, *ari.RecordingOptions) *ari.LiveRecordingHandle); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.LiveRecordingHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *ari.RecordingOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StageRecord provides a mock function with given fields: _a0, _a1
func (_m *Recorder) StageRecord(_a0 string, _a1 *ari.RecordingOptions) (*ari.LiveRecordingHandle, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ari.LiveRecordingHandle
	if rf, ok := ret.Get(0).(func(string, *ari.RecordingOptions) *ari.LiveRecordingHandle); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.LiveRecordingHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *ari.RecordingOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: n
func (_m *Recorder) Subscribe(n ...string) ari.Subscription {
	_va := make([]interface{}, len(n))
	for _i := range n {
		_va[_i] = n[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ari.Subscription
	if rf, ok := ret.Get(0).(func(...string) ari.Subscription); ok {
		r0 = rf(n...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Subscription)
		}
	}

	return r0
}
