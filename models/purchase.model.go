package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	ID            string        `sql:"type:uuid;primary_key;"`
	PaymentMethod PaymentMethod `validate:"required"`
	UserID        string        `validate:"required"`
	Success       bool
}

func (p *Purchase) BeforeCreate(db *gorm.DB) (err error) {

	newId := uuid.NewV4().String()
	p.Success = false
	p.ID = newId
	return
}
