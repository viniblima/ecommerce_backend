package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   string `sql:"type:uuid;primary_key;"`
	Name string `validate:"required,min=4,max=15"`
}

func (c *Category) BeforeCreate(db *gorm.DB) (err error) {
	c.ID = uuid.NewV4().String()
	return
}
