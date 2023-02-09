package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ProductDiscount struct {
	gorm.Model
	ID         string `sql:"type:uuid;primary_key;"`
	ProductID  string `validate:"required"`
	DiscountID string `validate:"required"`
}

func (p *ProductDiscount) BeforeCreate(db *gorm.DB) (err error) {
	p.ID = uuid.NewV4().String()
	return
}
