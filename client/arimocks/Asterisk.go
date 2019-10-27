// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/CyCoreSystems/ari"
	mock "github.com/stretchr/testify/mock"
)

// Asterisk is an autogenerated mock type for the Asterisk type
type Asterisk struct {
	mock.Mock
}

// Config provides a mock function with given fields:
func (_m *Asterisk) Config() ari.Config {
	ret := _m.Called()

	var r0 ari.Config
	if rf, ok := ret.Get(0).(func() ari.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Config)
		}
	}

	return r0
}

// Info provides a mock function with given fields: key
func (_m *Asterisk) Info(key *ari.Key) (*ari.AsteriskInfo, error) {
	ret := _m.Called(key)

	var r0 *ari.AsteriskInfo
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.AsteriskInfo); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.AsteriskInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logging provides a mock function with given fields:
func (_m *Asterisk) Logging() ari.Logging {
	ret := _m.Called()

	var r0 ari.Logging
	if rf, ok := ret.Get(0).(func() ari.Logging); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Logging)
		}
	}

	return r0
}

// Modules provides a mock function with given fields:
func (_m *Asterisk) Modules() ari.Modules {
	ret := _m.Called()

	var r0 ari.Modules
	if rf, ok := ret.Get(0).(func() ari.Modules); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Modules)
		}
	}

	return r0
}

// Variables provides a mock function with given fields:
func (_m *Asterisk) Variables() ari.AsteriskVariables {
	ret := _m.Called()

	var r0 ari.AsteriskVariables
	if rf, ok := ret.Get(0).(func() ari.AsteriskVariables); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.AsteriskVariables)
		}
	}

	return r0
}
