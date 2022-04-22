package service

import (
	"app-echo/domain/model"

	"gorm.io/gorm"
)

type DBUserService struct {
	db *gorm.DB
}

func NewDBUserService(db *gorm.DB) DBUserService {
	return DBUserService{
		db: db,
	}
}

func (us DBUserService) Create(user model.User) (model.User, error) {
	tx := us.db.Save(&user)
	err := tx.Error
	return user, err
}

func (us DBUserService) GetAll() ([]model.User, error) {
	users := []model.User{}
	tx := us.db.Find(&users)
	err := tx.Error
	return users, err
}
