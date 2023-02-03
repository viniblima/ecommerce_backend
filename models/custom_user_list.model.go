package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CustomUserList struct {
	gorm.Model
	ID     string `sql:"type:uuid;primary_key;"`
	Name   string `validate:"required,min=4,max=15"`
	UserID string
}

func (c *CustomUserList) BeforeCreate(db *gorm.DB) (err error) {
	c.ID = uuid.NewV4().String()
	return
}

type CustomUserListProducts struct {
	gorm.Model
	ID               string `sql:"type:uuid;primary_key;"`
	CustomUserListID string `gorm:"foreignKey: CustomUserListID" validate:"required"`
	ProductID        string `gorm:"foreignKey: ProductID" validate:"required"`
}

func (c *CustomUserListProducts) BeforeCreate(db *gorm.DB) (err error) {
	c.ID = uuid.NewV4().String()
	return
}
