package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type OfferProduct struct {
	gorm.Model
	ID        string `sql:"type:uuid;primary_key;"`
	ProductID string `validate:"required"`
	OfferID   string `validate:"required"`
}

func (offer *OfferProduct) BeforeCreate(db *gorm.DB) (err error) {

	newId := uuid.NewV4().String()
	offer.ID = newId
	return
}
