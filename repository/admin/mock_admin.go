package admin

import (
	"HMS-16-BE/model"
	"github.com/stretchr/testify/mock"
)

type MockAdminRepository struct {
	mock.Mock
}

func (m *MockAdminRepository) Create(admin model.Admins) error {
	ret := m.Called(admin)
	return ret.Error(0)
}

func (m *MockAdminRepository) Login(username string) (model.Admins, error) {
	ret := m.Called(username)
	return ret.Get(0).(model.Admins), ret.Error(1)
}

func (m *MockAdminRepository) GetById(id string) (model.Admins, error) {
	ret := m.Called(id)
	return ret.Get(0).(model.Admins), ret.Error(1)
}

func (m *MockAdminRepository) Update(admin model.Admins) error {
	ret := m.Called(admin)
	return ret.Error(0)
}

func (m *MockAdminRepository) Delete(id string) error {
	ret := m.Called(id)
	return ret.Error(0)
}
