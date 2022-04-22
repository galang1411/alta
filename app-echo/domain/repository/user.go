package repository

import "app-echo/domain/model"

type UserRepository interface {
	Create(user model.User) (model.User, error)
	GetAll() ([]model.User, error)
}
