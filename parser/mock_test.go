// Code generated by MockGen. DO NOT EDIT.
// Source: parser.go

// Package parser is a generated GoMock package.
package parser

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockdecoder is a mock of decoder interface
type Mockdecoder struct {
	ctrl     *gomock.Controller
	recorder *MockdecoderMockRecorder
}

// MockdecoderMockRecorder is the mock recorder for Mockdecoder
type MockdecoderMockRecorder struct {
	mock *Mockdecoder
}

// NewMockdecoder creates a new mock instance
func NewMockdecoder(ctrl *gomock.Controller) *Mockdecoder {
	mock := &Mockdecoder{ctrl: ctrl}
	mock.recorder = &MockdecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockdecoder) EXPECT() *MockdecoderMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *Mockdecoder) Decode(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockdecoderMockRecorder) Decode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*Mockdecoder)(nil).Decode), arg0)
}