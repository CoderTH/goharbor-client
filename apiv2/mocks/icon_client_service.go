// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	icon "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/icon"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockIconClientService is an autogenerated mock type for the ClientService type
type MockIconClientService struct {
	mock.Mock
}

// GetIcon provides a mock function with given fields: params, authInfo, opts
func (_m *MockIconClientService) GetIcon(params *icon.GetIconParams, authInfo runtime.ClientAuthInfoWriter, opts ...icon.ClientOption) (*icon.GetIconOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *icon.GetIconOK
	if rf, ok := ret.Get(0).(func(*icon.GetIconParams, runtime.ClientAuthInfoWriter, ...icon.ClientOption) *icon.GetIconOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*icon.GetIconOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*icon.GetIconParams, runtime.ClientAuthInfoWriter, ...icon.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockIconClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
