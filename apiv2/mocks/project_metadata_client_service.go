// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	project_metadata "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/project_metadata"
	mock "github.com/stretchr/testify/mock"
)

// MockProject_metadataClientService is an autogenerated mock type for the ClientService type
type MockProject_metadataClientService struct {
	mock.Mock
}

// AddProjectMetadatas provides a mock function with given fields: params, authInfo
func (_m *MockProject_metadataClientService) AddProjectMetadatas(params *project_metadata.AddProjectMetadatasParams, authInfo runtime.ClientAuthInfoWriter) (*project_metadata.AddProjectMetadatasOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project_metadata.AddProjectMetadatasOK
	if rf, ok := ret.Get(0).(func(*project_metadata.AddProjectMetadatasParams, runtime.ClientAuthInfoWriter) *project_metadata.AddProjectMetadatasOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project_metadata.AddProjectMetadatasOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project_metadata.AddProjectMetadatasParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProjectMetadata provides a mock function with given fields: params, authInfo
func (_m *MockProject_metadataClientService) DeleteProjectMetadata(params *project_metadata.DeleteProjectMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*project_metadata.DeleteProjectMetadataOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project_metadata.DeleteProjectMetadataOK
	if rf, ok := ret.Get(0).(func(*project_metadata.DeleteProjectMetadataParams, runtime.ClientAuthInfoWriter) *project_metadata.DeleteProjectMetadataOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project_metadata.DeleteProjectMetadataOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project_metadata.DeleteProjectMetadataParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectMetadata provides a mock function with given fields: params, authInfo
func (_m *MockProject_metadataClientService) GetProjectMetadata(params *project_metadata.GetProjectMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*project_metadata.GetProjectMetadataOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project_metadata.GetProjectMetadataOK
	if rf, ok := ret.Get(0).(func(*project_metadata.GetProjectMetadataParams, runtime.ClientAuthInfoWriter) *project_metadata.GetProjectMetadataOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project_metadata.GetProjectMetadataOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project_metadata.GetProjectMetadataParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjectMetadatas provides a mock function with given fields: params, authInfo
func (_m *MockProject_metadataClientService) ListProjectMetadatas(params *project_metadata.ListProjectMetadatasParams, authInfo runtime.ClientAuthInfoWriter) (*project_metadata.ListProjectMetadatasOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project_metadata.ListProjectMetadatasOK
	if rf, ok := ret.Get(0).(func(*project_metadata.ListProjectMetadatasParams, runtime.ClientAuthInfoWriter) *project_metadata.ListProjectMetadatasOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project_metadata.ListProjectMetadatasOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project_metadata.ListProjectMetadatasParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockProject_metadataClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateProjectMetadata provides a mock function with given fields: params, authInfo
func (_m *MockProject_metadataClientService) UpdateProjectMetadata(params *project_metadata.UpdateProjectMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*project_metadata.UpdateProjectMetadataOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *project_metadata.UpdateProjectMetadataOK
	if rf, ok := ret.Get(0).(func(*project_metadata.UpdateProjectMetadataParams, runtime.ClientAuthInfoWriter) *project_metadata.UpdateProjectMetadataOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*project_metadata.UpdateProjectMetadataOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*project_metadata.UpdateProjectMetadataParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockProject_metadataClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockProject_metadataClientService creates a new instance of MockProject_metadataClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockProject_metadataClientService(t mockConstructorTestingTNewMockProject_metadataClientService) *MockProject_metadataClientService {
	mock := &MockProject_metadataClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
