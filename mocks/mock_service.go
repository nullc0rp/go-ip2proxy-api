// Code generated by MockGen. DO NOT EDIT.
// Source: service/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	service "github.com/nullc0rp/go-ip2proxy-api/service"
	net "net"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetIPInfo mocks base method
func (m *MockService) GetIPInfo(ip net.IP) (*service.IPData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPInfo", ip)
	ret0, _ := ret[0].(*service.IPData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPInfo indicates an expected call of GetIPInfo
func (mr *MockServiceMockRecorder) GetIPInfo(ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPInfo", reflect.TypeOf((*MockService)(nil).GetIPInfo), ip)
}

// GetIPCountry mocks base method
func (m *MockService) GetIPCountry(country string, limit int) (*service.IPCountryData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPCountry", country, limit)
	ret0, _ := ret[0].(*service.IPCountryData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPCountry indicates an expected call of GetIPCountry
func (mr *MockServiceMockRecorder) GetIPCountry(country, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPCountry", reflect.TypeOf((*MockService)(nil).GetIPCountry), country, limit)
}

// GetISPCountry mocks base method
func (m *MockService) GetISPCountry(country string) (*service.ISPCountryData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetISPCountry", country)
	ret0, _ := ret[0].(*service.ISPCountryData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetISPCountry indicates an expected call of GetISPCountry
func (mr *MockServiceMockRecorder) GetISPCountry(country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetISPCountry", reflect.TypeOf((*MockService)(nil).GetISPCountry), country)
}

// GetCountryTotal mocks base method
func (m *MockService) GetCountryTotal(country string) (*service.IPCountryTotal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountryTotal", country)
	ret0, _ := ret[0].(*service.IPCountryTotal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountryTotal indicates an expected call of GetCountryTotal
func (mr *MockServiceMockRecorder) GetCountryTotal(country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountryTotal", reflect.TypeOf((*MockService)(nil).GetCountryTotal), country)
}

// MostProxyTypes mocks base method
func (m *MockService) MostProxyTypes() (*service.MostProxyTypeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MostProxyTypes")
	ret0, _ := ret[0].(*service.MostProxyTypeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MostProxyTypes indicates an expected call of MostProxyTypes
func (mr *MockServiceMockRecorder) MostProxyTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MostProxyTypes", reflect.TypeOf((*MockService)(nil).MostProxyTypes))
}
