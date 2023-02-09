package handlers

import (
	"errors"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

type ProductHandler interface {
	GetHighlights()
	GetAllProducts()
	GetProductByID(string)
	CreateProduct(models.Product)
}

func GetHighlights() []models.Product {
	var products []models.Product
	database.DB.Db.Where("highlight = ?", true).Find(&products)
	return products
}

func GetAllProducts() []map[string]interface{} {
	var products []models.Product
	database.DB.Db.Find(&products)
	var list []map[string]interface{}
	for i := 0; i < len(products); i++ {
		product := products[i]

		ds, errDs := GetDiscountbyProductID(product.ID)

		l := map[string]interface{}{
			"ID":                      product.ID,
			"Name":                    product.Name,
			"Price":                   product.Price,
			"Quantity":                product.Quantity,
			"MaxQuantityInstallments": product.MaxQuantityInstallments,
			"Highlight":               product.Highlight,
			"Discount":                ds,
		}

		if errDs != nil {
			l["Discount"] = nil
		}

		list = append(list, l)
	}
	return list
}

func GetProductByID(id string) (models.Product, error) {
	var product models.Product

	var err error

	dbResult := database.DB.Db.Where("ID = ?", id).First(&product)
	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("user not found")
	}
	return product, err
}

func CreateProduct(product *models.Product) *gorm.DB {
	return database.DB.Db.Create(&product)
}
