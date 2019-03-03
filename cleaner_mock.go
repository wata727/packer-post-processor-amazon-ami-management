// Code generated by MockGen. DO NOT EDIT.
// Source: cleaner.go

// Package main is a generated GoMock package.
package main

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAbstractCleaner is a mock of AbstractCleaner interface
type MockAbstractCleaner struct {
	ctrl     *gomock.Controller
	recorder *MockAbstractCleanerMockRecorder
}

// MockAbstractCleanerMockRecorder is the mock recorder for MockAbstractCleaner
type MockAbstractCleanerMockRecorder struct {
	mock *MockAbstractCleaner
}

// NewMockAbstractCleaner creates a new mock instance
func NewMockAbstractCleaner(ctrl *gomock.Controller) *MockAbstractCleaner {
	mock := &MockAbstractCleaner{ctrl: ctrl}
	mock.recorder = &MockAbstractCleanerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAbstractCleaner) EXPECT() *MockAbstractCleanerMockRecorder {
	return m.recorder
}

// RetrieveCandidateImages mocks base method
func (m *MockAbstractCleaner) RetrieveCandidateImages() ([]*ec2.Image, error) {
	ret := m.ctrl.Call(m, "RetrieveCandidateImages")
	ret0, _ := ret[0].([]*ec2.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveCandidateImages indicates an expected call of RetrieveCandidateImages
func (mr *MockAbstractCleanerMockRecorder) RetrieveCandidateImages() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveCandidateImages", reflect.TypeOf((*MockAbstractCleaner)(nil).RetrieveCandidateImages))
}

// DeleteImage mocks base method
func (m *MockAbstractCleaner) DeleteImage(arg0 *ec2.Image) error {
	ret := m.ctrl.Call(m, "DeleteImage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImage indicates an expected call of DeleteImage
func (mr *MockAbstractCleanerMockRecorder) DeleteImage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockAbstractCleaner)(nil).DeleteImage), arg0)
}

// IsUsed mocks base method
func (m *MockAbstractCleaner) IsUsed(arg0 *ec2.Image) *Used {
	ret := m.ctrl.Call(m, "IsUsed", arg0)
	ret0, _ := ret[0].(*Used)
	return ret0
}

// IsUsed indicates an expected call of IsUsed
func (mr *MockAbstractCleanerMockRecorder) IsUsed(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUsed", reflect.TypeOf((*MockAbstractCleaner)(nil).IsUsed), arg0)
}
