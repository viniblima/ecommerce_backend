package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CustomUserList struct {
	gorm.Model
	ID       string `sql:"type:uuid;primary_key;"`
	Name     string `validate:"required,min=4,max=15"`
	UserID   string
	Products []Product `gorm:"many2many:list_productts" validate:"required"`
}

func (c *CustomUserList) BeforeCreate(db *gorm.DB) (err error) {
	c.ID = uuid.NewV4().String()
	return
}
