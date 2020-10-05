// Code generated by MockGen. DO NOT EDIT.
// Source: domain/usecase/loan/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	entity "github.com/eminetto/clean-architecture-go-v2/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Borrow mocks base method
func (m *MockUseCase) Borrow(u *entity.User, b *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Borrow", u, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Borrow indicates an expected call of Borrow
func (mr *MockUseCaseMockRecorder) Borrow(u, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Borrow", reflect.TypeOf((*MockUseCase)(nil).Borrow), u, b)
}

// Return mocks base method
func (m *MockUseCase) Return(b *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Return", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Return indicates an expected call of Return
func (mr *MockUseCaseMockRecorder) Return(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Return", reflect.TypeOf((*MockUseCase)(nil).Return), b)
}
