// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package repository_mock is a generated GoMock package.
package repository_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repository "github.com/rogaliiik/bookstore/pkg/repository"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockRepository) CreateBook() *repository.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook")
	ret0, _ := ret[0].(*repository.Book)
	return ret0
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockRepositoryMockRecorder) CreateBook() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockRepository)(nil).CreateBook))
}

// DeleteBook mocks base method.
func (m *MockRepository) DeleteBook(Id int) repository.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", Id)
	ret0, _ := ret[0].(repository.Book)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockRepositoryMockRecorder) DeleteBook(Id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockRepository)(nil).DeleteBook), Id)
}

// GetAllBooks mocks base method.
func (m *MockRepository) GetAllBooks() []repository.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBooks")
	ret0, _ := ret[0].([]repository.Book)
	return ret0
}

// GetAllBooks indicates an expected call of GetAllBooks.
func (mr *MockRepositoryMockRecorder) GetAllBooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBooks", reflect.TypeOf((*MockRepository)(nil).GetAllBooks))
}

// GetBookById mocks base method.
func (m *MockRepository) GetBookById(Id int) *repository.Book {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookById", Id)
	ret0, _ := ret[0].(*repository.Book)
	return ret0
}

// GetBookById indicates an expected call of GetBookById.
func (mr *MockRepositoryMockRecorder) GetBookById(Id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookById", reflect.TypeOf((*MockRepository)(nil).GetBookById), Id)
}
