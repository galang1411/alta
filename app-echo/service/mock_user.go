package service

import (
	"app-echo/domain/model"
)

type MockUserService struct {
	data []model.User
}

func NewMockUserService() *MockUserService {
	return &MockUserService{
		data: []model.User{},
	}
}

func (us *MockUserService) Create(user model.User) (model.User, error) {
	us.data = append(us.data, user)
	return user, nil
}

func (us *MockUserService) GetAll() ([]model.User, error) {
	return us.data, nil
}
