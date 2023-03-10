package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CategoryProduct struct {
	gorm.Model
	ID         string `sql:"type:uuid;primary_key;"`
	ProductID  string `gorm:"foreignKey: ProductID" validate:"required"`
	CategoryID string `gorm:"foreignKey: CategoryID" validate:"required"`
}

func (c *CategoryProduct) BeforeCreate(db *gorm.DB) (err error) {
	c.ID = uuid.NewV4().String()
	return
}
