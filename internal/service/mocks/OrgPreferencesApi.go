// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	models "github.com/grafana/grafana-openapi-client-go/models"
	mock "github.com/stretchr/testify/mock"
)

// OrgPreferencesApi is an autogenerated mock type for the OrgPreferencesApi type
type OrgPreferencesApi struct {
	mock.Mock
}

type OrgPreferencesApi_Expecter struct {
	mock *mock.Mock
}

func (_m *OrgPreferencesApi) EXPECT() *OrgPreferencesApi_Expecter {
	return &OrgPreferencesApi_Expecter{mock: &_m.Mock}
}

// GetOrgPreferences provides a mock function with given fields: orgName
func (_m *OrgPreferencesApi) GetOrgPreferences(orgName string) (*models.Preferences, error) {
	ret := _m.Called(orgName)

	if len(ret) == 0 {
		panic("no return value specified for GetOrgPreferences")
	}

	var r0 *models.Preferences
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Preferences, error)); ok {
		return rf(orgName)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Preferences); ok {
		r0 = rf(orgName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Preferences)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orgName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrgPreferencesApi_GetOrgPreferences_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrgPreferences'
type OrgPreferencesApi_GetOrgPreferences_Call struct {
	*mock.Call
}

// GetOrgPreferences is a helper method to define mock.On call
//   - orgName string
func (_e *OrgPreferencesApi_Expecter) GetOrgPreferences(orgName interface{}) *OrgPreferencesApi_GetOrgPreferences_Call {
	return &OrgPreferencesApi_GetOrgPreferences_Call{Call: _e.mock.On("GetOrgPreferences", orgName)}
}

func (_c *OrgPreferencesApi_GetOrgPreferences_Call) Run(run func(orgName string)) *OrgPreferencesApi_GetOrgPreferences_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *OrgPreferencesApi_GetOrgPreferences_Call) Return(_a0 *models.Preferences, _a1 error) *OrgPreferencesApi_GetOrgPreferences_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrgPreferencesApi_GetOrgPreferences_Call) RunAndReturn(run func(string) (*models.Preferences, error)) *OrgPreferencesApi_GetOrgPreferences_Call {
	_c.Call.Return(run)
	return _c
}

// UploadOrgPreferences provides a mock function with given fields: orgName, pref
func (_m *OrgPreferencesApi) UploadOrgPreferences(orgName string, pref *models.Preferences) error {
	ret := _m.Called(orgName, pref)

	if len(ret) == 0 {
		panic("no return value specified for UploadOrgPreferences")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *models.Preferences) error); ok {
		r0 = rf(orgName, pref)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrgPreferencesApi_UploadOrgPreferences_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadOrgPreferences'
type OrgPreferencesApi_UploadOrgPreferences_Call struct {
	*mock.Call
}

// UploadOrgPreferences is a helper method to define mock.On call
//   - orgName string
//   - pref *models.Preferences
func (_e *OrgPreferencesApi_Expecter) UploadOrgPreferences(orgName interface{}, pref interface{}) *OrgPreferencesApi_UploadOrgPreferences_Call {
	return &OrgPreferencesApi_UploadOrgPreferences_Call{Call: _e.mock.On("UploadOrgPreferences", orgName, pref)}
}

func (_c *OrgPreferencesApi_UploadOrgPreferences_Call) Run(run func(orgName string, pref *models.Preferences)) *OrgPreferencesApi_UploadOrgPreferences_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*models.Preferences))
	})
	return _c
}

func (_c *OrgPreferencesApi_UploadOrgPreferences_Call) Return(_a0 error) *OrgPreferencesApi_UploadOrgPreferences_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrgPreferencesApi_UploadOrgPreferences_Call) RunAndReturn(run func(string, *models.Preferences) error) *OrgPreferencesApi_UploadOrgPreferences_Call {
	_c.Call.Return(run)
	return _c
}

// NewOrgPreferencesApi creates a new instance of OrgPreferencesApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrgPreferencesApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrgPreferencesApi {
	mock := &OrgPreferencesApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}