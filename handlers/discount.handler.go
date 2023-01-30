package handlers

import (
	"math"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func CreateDiscountLists(ds []models.DiscountsJson) []models.Discount {
	var discounts []models.Discount

	for i := 0; i < len(ds); i++ {
		var discount models.Discount

		discount.Product = GetProductByID(ds[i].ID)
		discount.PercentDiscount = roundFloat(ds[i].Percent, 2)
		discount.PriceWithDiscount = roundFloat(float64(discount.Product.Price)*(1-discount.PercentDiscount), 2)
		database.DB.Db.Create(&discount)

		discounts = append(discounts, discount)
	}

	return discounts
}
