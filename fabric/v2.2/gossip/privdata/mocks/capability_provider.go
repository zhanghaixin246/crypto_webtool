// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import channelconfig "crypto_webtool/fabric/v2.2/common/channelconfig"
import mock "github.com/stretchr/testify/mock"

// CapabilityProvider is an autogenerated mock type for the CapabilityProvider type
type CapabilityProvider struct {
	mock.Mock
}

// Capabilities provides a mock function with given fields:
func (_m *CapabilityProvider) Capabilities() channelconfig.ApplicationCapabilities {
	ret := _m.Called()

	var r0 channelconfig.ApplicationCapabilities
	if rf, ok := ret.Get(0).(func() channelconfig.ApplicationCapabilities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(channelconfig.ApplicationCapabilities)
		}
	}

	return r0
}
