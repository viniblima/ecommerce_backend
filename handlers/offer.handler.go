package handlers

import (
	"time"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
)

func GetAllOffers() []models.Offer {
	var offers []models.Offer
	database.DB.Db.Where("end_time > ?", time.Now()).Find(&offers)
	return offers
}

func DeleteOffer(id string) {
	database.DB.Db.Delete(models.Offer{}, id)
}
