// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/entity/animal.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	entity "github.com/TulioGuaraldoB/pizoo/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockIAnimalRepository is a mock of IAnimalRepository interface.
type MockIAnimalRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAnimalRepositoryMockRecorder
}

// MockIAnimalRepositoryMockRecorder is the mock recorder for MockIAnimalRepository.
type MockIAnimalRepositoryMockRecorder struct {
	mock *MockIAnimalRepository
}

// NewMockIAnimalRepository creates a new mock instance.
func NewMockIAnimalRepository(ctrl *gomock.Controller) *MockIAnimalRepository {
	mock := &MockIAnimalRepository{ctrl: ctrl}
	mock.recorder = &MockIAnimalRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAnimalRepository) EXPECT() *MockIAnimalRepositoryMockRecorder {
	return m.recorder
}

// CreateAnimal mocks base method.
func (m *MockIAnimalRepository) CreateAnimal(animal *entity.Animal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnimal", animal)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAnimal indicates an expected call of CreateAnimal.
func (mr *MockIAnimalRepositoryMockRecorder) CreateAnimal(animal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnimal", reflect.TypeOf((*MockIAnimalRepository)(nil).CreateAnimal), animal)
}

// DeleteAnimal mocks base method.
func (m *MockIAnimalRepository) DeleteAnimal(animalId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnimal", animalId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAnimal indicates an expected call of DeleteAnimal.
func (mr *MockIAnimalRepositoryMockRecorder) DeleteAnimal(animalId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnimal", reflect.TypeOf((*MockIAnimalRepository)(nil).DeleteAnimal), animalId)
}

// GetAllAnimals mocks base method.
func (m *MockIAnimalRepository) GetAllAnimals() ([]entity.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAnimals")
	ret0, _ := ret[0].([]entity.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAnimals indicates an expected call of GetAllAnimals.
func (mr *MockIAnimalRepositoryMockRecorder) GetAllAnimals() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAnimals", reflect.TypeOf((*MockIAnimalRepository)(nil).GetAllAnimals))
}

// GetAnimalById mocks base method.
func (m *MockIAnimalRepository) GetAnimalById(animalId uint) (*entity.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnimalById", animalId)
	ret0, _ := ret[0].(*entity.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnimalById indicates an expected call of GetAnimalById.
func (mr *MockIAnimalRepositoryMockRecorder) GetAnimalById(animalId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnimalById", reflect.TypeOf((*MockIAnimalRepository)(nil).GetAnimalById), animalId)
}
