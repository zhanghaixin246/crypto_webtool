// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	blockledger "crypto_webtool/fabric/v2.2/common/ledger/blockledger"
	mock "github.com/stretchr/testify/mock"
)

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// ChannelIDs provides a mock function with given fields:
func (_m *Factory) ChannelIDs() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Factory) Close() {
	_m.Called()
}

// GetOrCreate provides a mock function with given fields: chainID
func (_m *Factory) GetOrCreate(chainID string) (blockledger.ReadWriter, error) {
	ret := _m.Called(chainID)

	var r0 blockledger.ReadWriter
	if rf, ok := ret.Get(0).(func(string) blockledger.ReadWriter); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blockledger.ReadWriter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
