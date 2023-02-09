package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DiscountsJson struct {
	ID      string  `json:"id"`
	Percent float64 `json:"percent"`
}

type ItemDiscount struct {
	ID      string  `json:"ID"`
	Percent float64 `json:"Percent" validate:"required"`
}

type PayloadDiscount struct {
	List    []ItemDiscount `json:"Discounts"`
	EndTime time.Time      `json:"EndTime"`
}

type Discount struct {
	gorm.Model
	ID                string  `sql:"type:uuid;primary_key;"`
	PriceWithDiscount float64 `validate:"required"`
	PercentDiscount   float64 `validate:"required"`
}

func (d *Discount) BeforeCreate(db *gorm.DB) (err error) {
	newId := uuid.NewV4().String()
	d.ID = newId
	return
}
