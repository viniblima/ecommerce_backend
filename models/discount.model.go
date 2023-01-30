package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	ID        string `sql:"type:uuid;primary_key;"`
	ProductID string
	Product   Product `gorm:"foreignKey: ProductID"`

	PriceWithDiscount float64 `validate:"required"`
	PercentDiscount   float64 `validate:"required"`
}

type DiscountsJson struct {
	ID      string  `json:"id"`
	Percent float64 `json:"percent"`
}

func (d *Discount) BeforeCreate(db *gorm.DB) (err error) {
	newId := uuid.NewV4().String()
	d.ID = newId
	return
}
