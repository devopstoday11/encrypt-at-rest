// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	libcnb "github.com/buildpacks/libcnb"
	libpak "github.com/paketo-buildpacks/libpak"

	mock "github.com/stretchr/testify/mock"
)

// KeyProvider is an autogenerated mock type for the KeyProvider type
type KeyProvider struct {
	mock.Mock
}

// Detect provides a mock function with given fields: context, result
func (_m *KeyProvider) Detect(context libcnb.DetectContext, result *libcnb.DetectResult) error {
	ret := _m.Called(context, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(libcnb.DetectContext, *libcnb.DetectResult) error); ok {
		r0 = rf(context, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Key provides a mock function with given fields: context
func (_m *KeyProvider) Key(context libcnb.BuildContext) ([]byte, error) {
	ret := _m.Called(context)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(libcnb.BuildContext) []byte); ok {
		r0 = rf(context)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(libcnb.BuildContext) error); ok {
		r1 = rf(context)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Participate provides a mock function with given fields: resolver
func (_m *KeyProvider) Participate(resolver libpak.PlanEntryResolver) (bool, error) {
	ret := _m.Called(resolver)

	var r0 bool
	if rf, ok := ret.Get(0).(func(libpak.PlanEntryResolver) bool); ok {
		r0 = rf(resolver)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(libpak.PlanEntryResolver) error); ok {
		r1 = rf(resolver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
