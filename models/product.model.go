package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                      string  `sql:"type:uuid;primary_key;"`
	Name                    string  `validate:"required,min=4,max=15"`
	OriginalPrice           float32 `validate:"required"`
	PriceWithDiscount       float32 `validate:"required"`
	PercentDiscount         float32 `validate:"required"`
	Quantity                int     `validate:"required"`
	MaxQuantityInstallments int     `validate:"required,min=1"`
	// CategoriesID            []uint
	// Categories              []Category `gorm:"foreignKey: CategoriesID"`
}

func (p *Product) BeforeCreate(db *gorm.DB) (err error) {
	p.ID = uuid.NewV4().String()
	return
}
