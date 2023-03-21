package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `sql:"type:uuid;primary_key;"`
	Name     string `validate:"required,min=4,max=15"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	IsAdmin  bool
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	user.ID = uuid.NewV4().String()
	return
}
