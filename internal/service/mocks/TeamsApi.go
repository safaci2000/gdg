// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	filters "github.com/esnet/gdg/internal/service/filters"
	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana-openapi-client-go/models"
)

// TeamsApi is an autogenerated mock type for the TeamsApi type
type TeamsApi struct {
	mock.Mock
}

type TeamsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *TeamsApi) EXPECT() *TeamsApi_Expecter {
	return &TeamsApi_Expecter{mock: &_m.Mock}
}

// DeleteTeam provides a mock function with given fields: filter
func (_m *TeamsApi) DeleteTeam(filter filters.Filter) ([]*models.TeamDTO, error) {
	ret := _m.Called(filter)

	var r0 []*models.TeamDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(filters.Filter) ([]*models.TeamDTO, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(filters.Filter) []*models.TeamDTO); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TeamDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(filters.Filter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TeamsApi_DeleteTeam_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTeam'
type TeamsApi_DeleteTeam_Call struct {
	*mock.Call
}

// DeleteTeam is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *TeamsApi_Expecter) DeleteTeam(filter interface{}) *TeamsApi_DeleteTeam_Call {
	return &TeamsApi_DeleteTeam_Call{Call: _e.mock.On("DeleteTeam", filter)}
}

func (_c *TeamsApi_DeleteTeam_Call) Run(run func(filter filters.Filter)) *TeamsApi_DeleteTeam_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *TeamsApi_DeleteTeam_Call) Return(_a0 []*models.TeamDTO, _a1 error) *TeamsApi_DeleteTeam_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TeamsApi_DeleteTeam_Call) RunAndReturn(run func(filters.Filter) ([]*models.TeamDTO, error)) *TeamsApi_DeleteTeam_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadTeams provides a mock function with given fields: filter
func (_m *TeamsApi) DownloadTeams(filter filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO {
	ret := _m.Called(filter)

	var r0 map[*models.TeamDTO][]*models.TeamMemberDTO
	if rf, ok := ret.Get(0).(func(filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[*models.TeamDTO][]*models.TeamMemberDTO)
		}
	}

	return r0
}

// TeamsApi_DownloadTeams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadTeams'
type TeamsApi_DownloadTeams_Call struct {
	*mock.Call
}

// DownloadTeams is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *TeamsApi_Expecter) DownloadTeams(filter interface{}) *TeamsApi_DownloadTeams_Call {
	return &TeamsApi_DownloadTeams_Call{Call: _e.mock.On("DownloadTeams", filter)}
}

func (_c *TeamsApi_DownloadTeams_Call) Run(run func(filter filters.Filter)) *TeamsApi_DownloadTeams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *TeamsApi_DownloadTeams_Call) Return(_a0 map[*models.TeamDTO][]*models.TeamMemberDTO) *TeamsApi_DownloadTeams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TeamsApi_DownloadTeams_Call) RunAndReturn(run func(filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO) *TeamsApi_DownloadTeams_Call {
	_c.Call.Return(run)
	return _c
}

// ListTeams provides a mock function with given fields: filter
func (_m *TeamsApi) ListTeams(filter filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO {
	ret := _m.Called(filter)

	var r0 map[*models.TeamDTO][]*models.TeamMemberDTO
	if rf, ok := ret.Get(0).(func(filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[*models.TeamDTO][]*models.TeamMemberDTO)
		}
	}

	return r0
}

// TeamsApi_ListTeams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTeams'
type TeamsApi_ListTeams_Call struct {
	*mock.Call
}

// ListTeams is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *TeamsApi_Expecter) ListTeams(filter interface{}) *TeamsApi_ListTeams_Call {
	return &TeamsApi_ListTeams_Call{Call: _e.mock.On("ListTeams", filter)}
}

func (_c *TeamsApi_ListTeams_Call) Run(run func(filter filters.Filter)) *TeamsApi_ListTeams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *TeamsApi_ListTeams_Call) Return(_a0 map[*models.TeamDTO][]*models.TeamMemberDTO) *TeamsApi_ListTeams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TeamsApi_ListTeams_Call) RunAndReturn(run func(filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO) *TeamsApi_ListTeams_Call {
	_c.Call.Return(run)
	return _c
}

// UploadTeams provides a mock function with given fields: filter
func (_m *TeamsApi) UploadTeams(filter filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO {
	ret := _m.Called(filter)

	var r0 map[*models.TeamDTO][]*models.TeamMemberDTO
	if rf, ok := ret.Get(0).(func(filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[*models.TeamDTO][]*models.TeamMemberDTO)
		}
	}

	return r0
}

// TeamsApi_UploadTeams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadTeams'
type TeamsApi_UploadTeams_Call struct {
	*mock.Call
}

// UploadTeams is a helper method to define mock.On call
//   - filter filters.Filter
func (_e *TeamsApi_Expecter) UploadTeams(filter interface{}) *TeamsApi_UploadTeams_Call {
	return &TeamsApi_UploadTeams_Call{Call: _e.mock.On("UploadTeams", filter)}
}

func (_c *TeamsApi_UploadTeams_Call) Run(run func(filter filters.Filter)) *TeamsApi_UploadTeams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(filters.Filter))
	})
	return _c
}

func (_c *TeamsApi_UploadTeams_Call) Return(_a0 map[*models.TeamDTO][]*models.TeamMemberDTO) *TeamsApi_UploadTeams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TeamsApi_UploadTeams_Call) RunAndReturn(run func(filters.Filter) map[*models.TeamDTO][]*models.TeamMemberDTO) *TeamsApi_UploadTeams_Call {
	_c.Call.Return(run)
	return _c
}

// NewTeamsApi creates a new instance of TeamsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTeamsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *TeamsApi {
	mock := &TeamsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
