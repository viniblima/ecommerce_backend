package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                      string  `sql:"type:uuid;primary_key;"`
	Name                    string  `validate:"required,min=4,max=15"`
	Price                   float32 `validate:"required"`
	Quantity                int     `validate:"required"`
	MaxQuantityInstallments int     `validate:"required,min=1"`
	Highlight               bool
	// CategoriesID            []uint
	// Categories              []Category `gorm:"foreignKey: CategoriesID"`
}

func (p *Product) BeforeCreate(db *gorm.DB) (err error) {
	p.ID = uuid.NewV4().String()
	return
}
