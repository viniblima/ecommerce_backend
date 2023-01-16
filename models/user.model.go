package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `validate:"required,min=4,max=15"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
