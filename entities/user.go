package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Username string
	Password string
	Email    string
}