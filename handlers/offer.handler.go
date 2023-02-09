package handlers

import (
	"errors"
	"fmt"
	"time"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

type OfferHandler interface {
	GetAllOffers() []models.Offer
	CreateOffer(models.Offer) *gorm.DB
}

func GetAllOffers() []map[string]interface{} {
	var offers []models.Offer
	database.DB.Db.Where("end_time > ?", time.Now()).Find(&offers)
	var list []map[string]interface{}
	for i := 0; i < len(offers); i++ {

		offer := offers[i]

		var relations []models.OfferProduct

		database.DB.Db.Where("offer_id = ?", offer.ID).Find(&relations)
		var rlMap []map[string]interface{}
		for j := 0; j < len(relations); j++ {

			relation := relations[j]

			fmt.Println(relation.ProductID)
			product, err := GetProductByID(relation.ProductID)

			if err == nil {
				ds, _ := GetDiscountbyProductID(product.ID)

				rl := map[string]interface{}{
					"Product": map[string]interface{}{
						"ID":                      product.ID,
						"Name":                    product.Name,
						"Price":                   product.Price,
						"MaxQuantityInstallments": product.MaxQuantityInstallments,
						"Discount":                ds,
					},
				}

				rlMap = append(rlMap, rl)
				//list = append(list, rl)
			}
		}

		l := map[string]interface{}{
			"Offer":    offer,
			"Products": relations,
		}
		list = append(list, l)
		// 	offer := offers[i]
		// 	for j := 0; j < len(offer.Products); j++ {
		// 		product := offer.Products[j]
		// 		//p, _ := GetProductByID(discount.ProductID)

		// 		ds := map[string]interface{}{
		// 			"Product": product,
		// 		}
		// 		list = append(list, ds)
		// 		//discount.Product, _ = GetProductByID(discount.ProductID)
		// 	}
	}
	return list
}

func CreateOffer(o models.Offer) *gorm.DB {

	return database.DB.Db.Create(&o)
}

func GetOfferByID(id string) (models.Offer, error) {
	offer := models.Offer{}
	var err error
	dbResult := database.DB.Db.Where("id = ?", id).First(&offer)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Relation not found")
	}

	return offer, err
}
