// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-api/cluster (interfaces: JobPluginAPI)

// Package mock_cluster is a generated GoMock package.
package mock_cluster

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost/server/public/model"
)

// MockJobPluginAPI is a mock of JobPluginAPI interface.
type MockJobPluginAPI struct {
	ctrl     *gomock.Controller
	recorder *MockJobPluginAPIMockRecorder
}

// MockJobPluginAPIMockRecorder is the mock recorder for MockJobPluginAPI.
type MockJobPluginAPIMockRecorder struct {
	mock *MockJobPluginAPI
}

// NewMockJobPluginAPI creates a new mock instance.
func NewMockJobPluginAPI(ctrl *gomock.Controller) *MockJobPluginAPI {
	mock := &MockJobPluginAPI{ctrl: ctrl}
	mock.recorder = &MockJobPluginAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobPluginAPI) EXPECT() *MockJobPluginAPIMockRecorder {
	return m.recorder
}

// KVDelete mocks base method.
func (m *MockJobPluginAPI) KVDelete(arg0 string) *model.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KVDelete", arg0)
	ret0, _ := ret[0].(*model.AppError)
	return ret0
}

// KVDelete indicates an expected call of KVDelete.
func (mr *MockJobPluginAPIMockRecorder) KVDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KVDelete", reflect.TypeOf((*MockJobPluginAPI)(nil).KVDelete), arg0)
}

// KVGet mocks base method.
func (m *MockJobPluginAPI) KVGet(arg0 string) ([]byte, *model.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KVGet", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*model.AppError)
	return ret0, ret1
}

// KVGet indicates an expected call of KVGet.
func (mr *MockJobPluginAPIMockRecorder) KVGet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KVGet", reflect.TypeOf((*MockJobPluginAPI)(nil).KVGet), arg0)
}

// KVList mocks base method.
func (m *MockJobPluginAPI) KVList(arg0, arg1 int) ([]string, *model.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KVList", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*model.AppError)
	return ret0, ret1
}

// KVList indicates an expected call of KVList.
func (mr *MockJobPluginAPIMockRecorder) KVList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KVList", reflect.TypeOf((*MockJobPluginAPI)(nil).KVList), arg0, arg1)
}

// KVSetWithOptions mocks base method.
func (m *MockJobPluginAPI) KVSetWithOptions(arg0 string, arg1 []byte, arg2 model.PluginKVSetOptions) (bool, *model.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KVSetWithOptions", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*model.AppError)
	return ret0, ret1
}

// KVSetWithOptions indicates an expected call of KVSetWithOptions.
func (mr *MockJobPluginAPIMockRecorder) KVSetWithOptions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KVSetWithOptions", reflect.TypeOf((*MockJobPluginAPI)(nil).KVSetWithOptions), arg0, arg1, arg2)
}

// LogError mocks base method.
func (m *MockJobPluginAPI) LogError(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LogError", varargs...)
}

// LogError indicates an expected call of LogError.
func (mr *MockJobPluginAPIMockRecorder) LogError(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogError", reflect.TypeOf((*MockJobPluginAPI)(nil).LogError), varargs...)
}
