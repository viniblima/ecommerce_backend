package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	ID        string `sql:"type:uuid;primary_key;" `
	ProductID string `gorm:"foreignKey: ProductID" validate:"required"`
	UserID    string `gorm:"foreignKey: UserID" validate:"required"`
}

func (p *Like) BeforeCreate(db *gorm.DB) (err error) {
	p.ID = uuid.NewV4().String()
	return
}
