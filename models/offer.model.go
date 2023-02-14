package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Offer struct {
	gorm.Model
	ID      string    `sql:"type:uuid;primary_key;"`
	EndTime time.Time `validate:"required"`
}

func (offer *Offer) BeforeCreate(db *gorm.DB) (err error) {

	newId := uuid.NewV4().String()
	offer.ID = newId
	return
}
