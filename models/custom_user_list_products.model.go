package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

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
