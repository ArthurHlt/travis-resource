// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/Orange-OpenSource/travis-resource/travis (interfaces: BuildsInterface,JobsInterface)

package mock_travis

import (
	travis "github.com/Orange-OpenSource/travis-resource/travis"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
)

// Mock of BuildsInterface interface
type MockBuildsInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockBuildsInterfaceRecorder
}

// Recorder for MockBuildsInterface (not exported)
type _MockBuildsInterfaceRecorder struct {
	mock *MockBuildsInterface
}

func NewMockBuildsInterface(ctrl *gomock.Controller) *MockBuildsInterface {
	mock := &MockBuildsInterface{ctrl: ctrl}
	mock.recorder = &_MockBuildsInterfaceRecorder{mock}
	return mock
}

func (_m *MockBuildsInterface) EXPECT() *_MockBuildsInterfaceRecorder {
	return _m.recorder
}

func (_m *MockBuildsInterface) Cancel(_param0 uint) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "Cancel", _param0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBuildsInterfaceRecorder) Cancel(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Cancel", arg0)
}

func (_m *MockBuildsInterface) Get(_param0 uint) (*travis.Build, []travis.Job, *travis.Commit, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*travis.Build)
	ret1, _ := ret[1].([]travis.Job)
	ret2, _ := ret[2].(*travis.Commit)
	ret3, _ := ret[3].(*http.Response)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

func (_mr *_MockBuildsInterfaceRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockBuildsInterface) GetFirstBuildFromBuildNumber(_param0 string, _param1 string) (travis.Build, error) {
	ret := _m.ctrl.Call(_m, "GetFirstBuildFromBuildNumber", _param0, _param1)
	ret0, _ := ret[0].(travis.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBuildsInterfaceRecorder) GetFirstBuildFromBuildNumber(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFirstBuildFromBuildNumber", arg0, arg1)
}

func (_m *MockBuildsInterface) GetFirstFinishedBuild(_param0 string) (travis.Build, error) {
	ret := _m.ctrl.Call(_m, "GetFirstFinishedBuild", _param0)
	ret0, _ := ret[0].(travis.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBuildsInterfaceRecorder) GetFirstFinishedBuild(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFirstFinishedBuild", arg0)
}

func (_m *MockBuildsInterface) GetFirstFinishedBuildWithBranch(_param0 string, _param1 string) (travis.Build, error) {
	ret := _m.ctrl.Call(_m, "GetFirstFinishedBuildWithBranch", _param0, _param1)
	ret0, _ := ret[0].(travis.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBuildsInterfaceRecorder) GetFirstFinishedBuildWithBranch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFirstFinishedBuildWithBranch", arg0, arg1)
}

func (_m *MockBuildsInterface) List(_param0 *travis.BuildListOptions) ([]travis.Build, []travis.Job, []travis.Commit, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "List", _param0)
	ret0, _ := ret[0].([]travis.Build)
	ret1, _ := ret[1].([]travis.Job)
	ret2, _ := ret[2].([]travis.Commit)
	ret3, _ := ret[3].(*http.Response)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

func (_mr *_MockBuildsInterfaceRecorder) List(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List", arg0)
}

func (_m *MockBuildsInterface) ListFromRepository(_param0 string, _param1 *travis.BuildListOptions) ([]travis.Build, []travis.Job, []travis.Commit, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "ListFromRepository", _param0, _param1)
	ret0, _ := ret[0].([]travis.Build)
	ret1, _ := ret[1].([]travis.Job)
	ret2, _ := ret[2].([]travis.Commit)
	ret3, _ := ret[3].(*http.Response)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

func (_mr *_MockBuildsInterfaceRecorder) ListFromRepository(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListFromRepository", arg0, arg1)
}

func (_m *MockBuildsInterface) ListFromRepositoryWithBranch(_param0 string, _param1 string, _param2 *travis.BuildListOptions) ([]travis.Build, []travis.Job, []travis.Commit, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "ListFromRepositoryWithBranch", _param0, _param1, _param2)
	ret0, _ := ret[0].([]travis.Build)
	ret1, _ := ret[1].([]travis.Job)
	ret2, _ := ret[2].([]travis.Commit)
	ret3, _ := ret[3].(*http.Response)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

func (_mr *_MockBuildsInterfaceRecorder) ListFromRepositoryWithBranch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListFromRepositoryWithBranch", arg0, arg1, arg2)
}

func (_m *MockBuildsInterface) ListSucceededFromRepository(_param0 string, _param1 *travis.BuildListOptions) ([]travis.Build, []travis.Job, []travis.Commit, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "ListSucceededFromRepository", _param0, _param1)
	ret0, _ := ret[0].([]travis.Build)
	ret1, _ := ret[1].([]travis.Job)
	ret2, _ := ret[2].([]travis.Commit)
	ret3, _ := ret[3].(*http.Response)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

func (_mr *_MockBuildsInterfaceRecorder) ListSucceededFromRepository(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListSucceededFromRepository", arg0, arg1)
}

func (_m *MockBuildsInterface) ListSucceededFromRepositoryWithBranch(_param0 string, _param1 string, _param2 *travis.BuildListOptions) ([]travis.Build, []travis.Job, []travis.Commit, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "ListSucceededFromRepositoryWithBranch", _param0, _param1, _param2)
	ret0, _ := ret[0].([]travis.Build)
	ret1, _ := ret[1].([]travis.Job)
	ret2, _ := ret[2].([]travis.Commit)
	ret3, _ := ret[3].(*http.Response)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

func (_mr *_MockBuildsInterfaceRecorder) ListSucceededFromRepositoryWithBranch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListSucceededFromRepositoryWithBranch", arg0, arg1, arg2)
}

func (_m *MockBuildsInterface) Restart(_param0 uint) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "Restart", _param0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockBuildsInterfaceRecorder) Restart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Restart", arg0)
}

// Mock of JobsInterface interface
type MockJobsInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockJobsInterfaceRecorder
}

// Recorder for MockJobsInterface (not exported)
type _MockJobsInterfaceRecorder struct {
	mock *MockJobsInterface
}

func NewMockJobsInterface(ctrl *gomock.Controller) *MockJobsInterface {
	mock := &MockJobsInterface{ctrl: ctrl}
	mock.recorder = &_MockJobsInterfaceRecorder{mock}
	return mock
}

func (_m *MockJobsInterface) EXPECT() *_MockJobsInterfaceRecorder {
	return _m.recorder
}

func (_m *MockJobsInterface) Cancel(_param0 uint) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "Cancel", _param0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobsInterfaceRecorder) Cancel(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Cancel", arg0)
}

func (_m *MockJobsInterface) Find(_param0 *travis.JobFindOptions) ([]travis.Job, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "Find", _param0)
	ret0, _ := ret[0].([]travis.Job)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJobsInterfaceRecorder) Find(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Find", arg0)
}

func (_m *MockJobsInterface) Get(_param0 uint) (*travis.Job, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*travis.Job)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJobsInterfaceRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockJobsInterface) ListFromBuild(_param0 uint) ([]travis.Job, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "ListFromBuild", _param0)
	ret0, _ := ret[0].([]travis.Job)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJobsInterfaceRecorder) ListFromBuild(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListFromBuild", arg0)
}

func (_m *MockJobsInterface) RawLog(_param0 uint) ([]byte, *http.Response, error) {
	ret := _m.ctrl.Call(_m, "RawLog", _param0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJobsInterfaceRecorder) RawLog(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RawLog", arg0)
}

func (_m *MockJobsInterface) Restart(_param0 uint) (*http.Response, error) {
	ret := _m.ctrl.Call(_m, "Restart", _param0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJobsInterfaceRecorder) Restart(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Restart", arg0)
}
