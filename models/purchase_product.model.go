package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PurchaseProduct struct {
	gorm.Model
	ID         string `sql:"type:uuid;primary_key;"`
	PurchaseID string `validate:"required"`
}

func (p *PurchaseProduct) BeforeCreate(db *gorm.DB) (err error) {

	newId := uuid.NewV4().String()
	p.ID = newId
	return
}
