// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	filters "github.com/esnet/gdg/internal/service/filters"
	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana-openapi-client-go/models"
)

// DashboardsApi is an autogenerated mock type for the DashboardsApi type
type DashboardsApi struct {
	mock.Mock
}

type DashboardsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *DashboardsApi) EXPECT() *DashboardsApi_Expecter {
	return &DashboardsApi_Expecter{mock: &_m.Mock}
}

// DeleteAllDashboards provides a mock function with given fields: filter
func (_m *DashboardsApi) DeleteAllDashboards(filter filters.Filter) []string {
	ret := _m.Called(filter)

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// DashboardsApi_DeleteAllDashboards_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAllDashboards'
type DashboardsApi_DeleteAllDashboards_Call struct {
	*mock.Call
}

// DeleteAllDashboards is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *DashboardsApi_Expecter) DeleteAllDashboards(filter interface{}) *DashboardsApi_DeleteAllDashboards_Call {
	return &DashboardsApi_DeleteAllDashboards_Call{Call: _e.mock.On("DeleteAllDashboards", filter)}
}

func (_c *DashboardsApi_DeleteAllDashboards_Call) Run(run func(filter filters.Filter)) *DashboardsApi_DeleteAllDashboards_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *DashboardsApi_DeleteAllDashboards_Call) Return(_a0 []string) *DashboardsApi_DeleteAllDashboards_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DashboardsApi_DeleteAllDashboards_Call) RunAndReturn(run func(filters.Filter) []string) *DashboardsApi_DeleteAllDashboards_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadDashboards provides a mock function with given fields: filter
func (_m *DashboardsApi) DownloadDashboards(filter filters.Filter) []string {
	ret := _m.Called(filter)

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// DashboardsApi_DownloadDashboards_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadDashboards'
type DashboardsApi_DownloadDashboards_Call struct {
	*mock.Call
}

// DownloadDashboards is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *DashboardsApi_Expecter) DownloadDashboards(filter interface{}) *DashboardsApi_DownloadDashboards_Call {
	return &DashboardsApi_DownloadDashboards_Call{Call: _e.mock.On("DownloadDashboards", filter)}
}

func (_c *DashboardsApi_DownloadDashboards_Call) Run(run func(filter filters.Filter)) *DashboardsApi_DownloadDashboards_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *DashboardsApi_DownloadDashboards_Call) Return(_a0 []string) *DashboardsApi_DownloadDashboards_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DashboardsApi_DownloadDashboards_Call) RunAndReturn(run func(filters.Filter) []string) *DashboardsApi_DownloadDashboards_Call {
	_c.Call.Return(run)
	return _c
}

// ListDashboards provides a mock function with given fields: filter
func (_m *DashboardsApi) ListDashboards(filter filters.Filter) []*models.Hit {
	ret := _m.Called(filter)

	var r0 []*models.Hit
	if rf, ok := ret.Get(0).(func(filters.Filter) []*models.Hit); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Hit)
		}
	}

	return r0
}

// DashboardsApi_ListDashboards_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDashboards'
type DashboardsApi_ListDashboards_Call struct {
	*mock.Call
}

// ListDashboards is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *DashboardsApi_Expecter) ListDashboards(filter interface{}) *DashboardsApi_ListDashboards_Call {
	return &DashboardsApi_ListDashboards_Call{Call: _e.mock.On("ListDashboards", filter)}
}

func (_c *DashboardsApi_ListDashboards_Call) Run(run func(filter filters.Filter)) *DashboardsApi_ListDashboards_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *DashboardsApi_ListDashboards_Call) Return(_a0 []*models.Hit) *DashboardsApi_ListDashboards_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DashboardsApi_ListDashboards_Call) RunAndReturn(run func(filters.Filter) []*models.Hit) *DashboardsApi_ListDashboards_Call {
	_c.Call.Return(run)
	return _c
}

// UploadDashboards provides a mock function with given fields: filter
func (_m *DashboardsApi) UploadDashboards(filter filters.Filter) {
	_m.Called(filter)
}

// DashboardsApi_UploadDashboards_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadDashboards'
type DashboardsApi_UploadDashboards_Call struct {
	*mock.Call
}

// UploadDashboards is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *DashboardsApi_Expecter) UploadDashboards(filter interface{}) *DashboardsApi_UploadDashboards_Call {
	return &DashboardsApi_UploadDashboards_Call{Call: _e.mock.On("UploadDashboards", filter)}
}

func (_c *DashboardsApi_UploadDashboards_Call) Run(run func(filter filters.Filter)) *DashboardsApi_UploadDashboards_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *DashboardsApi_UploadDashboards_Call) Return() *DashboardsApi_UploadDashboards_Call {
	_c.Call.Return()
	return _c
}

func (_c *DashboardsApi_UploadDashboards_Call) RunAndReturn(run func(filters.Filter)) *DashboardsApi_UploadDashboards_Call {
	_c.Call.Return(run)
	return _c
}

// NewDashboardsApi creates a new instance of DashboardsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDashboardsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *DashboardsApi {
	mock := &DashboardsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
