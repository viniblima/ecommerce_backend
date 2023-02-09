package handlers

import (
	"errors"
	"fmt"
	"math"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func GetDiscountbyProductID(productID string) (models.Discount, error) {

	var relation models.ProductDiscount

	var discount models.Discount

	var err error

	resultRelation := database.DB.Db.Where("product_id = ?", productID).Last(&relation)

	fmt.Println("resultRelation")
	fmt.Println(relation.ID)
	if errors.Is(resultRelation.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Discount not found")
	}

	if err != nil {
		return discount, nil
	}

	dbResult := database.DB.Db.Where("id = ?", relation.DiscountID).First(&discount)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Discount not found")
	}

	return discount, err
}

func DeleteDiscount(discountID string) models.Discount {
	var discount models.Discount

	var relation models.ProductDiscount

	var offerProduct models.OfferProduct

	database.DB.Db.Where("discount_id = ?", discountID).Delete(&relation)

	database.DB.Db.Where("id = ?", discountID).Delete(&discount)

	database.DB.Db.Where("product_id = ?", relation.ProductID).Delete(&offerProduct)

	return discount
}

func GetAllDiscounts() []models.Discount {
	var discounts []models.Discount

	database.DB.Db.Find(&discounts)

	return discounts
}

func GetAllProductDiscounts() []models.ProductDiscount {
	var discounts []models.ProductDiscount

	database.DB.Db.Find(&discounts)

	return discounts
}

func CreateDiscountLists(offerID string, ps []models.ItemDiscount) []map[string]interface{} {

	var products []map[string]interface{}

	for i := 0; i < len(ps); i++ {
		fmt.Println("teste")
		//var discount models.Discount
		fmt.Println(ps[i].ID)
		product, err := GetProductByID(ps[i].ID)

		foundProduct := false

		for j := 0; j < len(ps); j++ {
			if product.ID == ps[j].ID {
				foundProduct = true
			}
		}

		if err == nil {
			fmt.Println("1")
			fmt.Println(product.Name)
			actualDiscount, errActualDiscount := GetDiscountbyProductID(product.ID)
			if errActualDiscount == nil || foundProduct {
				fmt.Println("2")
				DeleteDiscount(actualDiscount.ID)
			}
			newOfferRelation := new(models.OfferProduct)

			newOfferRelation.OfferID = offerID
			newOfferRelation.ProductID = product.ID

			database.DB.Db.Create(&newOfferRelation)

			newDiscount := new(models.Discount)
			newDiscount.PercentDiscount = roundFloat(ps[i].Percent, 2)
			newDiscount.PriceWithDiscount = roundFloat(float64(product.Price)*(1-ps[i].Percent), 2)

			database.DB.Db.Create(&newDiscount)

			newProductDiscountRelation := new(models.ProductDiscount)
			newProductDiscountRelation.ProductID = ps[i].ID
			newProductDiscountRelation.DiscountID = newDiscount.ID

			database.DB.Db.Create(&newProductDiscountRelation)

			m := map[string]interface{}{
				"Product": map[string]interface{}{
					"ID":                      product.ID,
					"Name":                    product.Name,
					"Price":                   product.Price,
					"MaxQuantityInstallments": product.MaxQuantityInstallments,
					"Discount":                newDiscount,
				},
			}

			products = append(products, m)
		}

	}

	return products
}
