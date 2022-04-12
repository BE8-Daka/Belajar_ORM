package datastore

import (
	"orm_selasa/entities"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func (u UserDB) GetAllUser() ([]entities.User, error) {
	var users []entities.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}