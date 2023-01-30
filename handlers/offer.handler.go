package handlers

import (
	"time"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

type OfferHandler interface {
	GetAllOffers() []models.Offer
	CreateOffer(models.Offer) *gorm.DB
}

func GetAllOffers() []models.Offer {
	var offers []models.Offer
	database.DB.Db.Where("end_time > ?", time.Now()).Preload("Discounts").Find(&offers)
	return offers
}

func CreateOffer(o models.Offer) *gorm.DB {

	return database.DB.Db.Create(&o)
}
