package handlers

import (
	"errors"
	"fmt"
	"strconv"

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

func GetAllProducts(page string) map[string]interface{} {
	var products []models.Product

	if page == "" {
		page = "1"
	}
	offset := 10
	limit := 10

	int, errOffeset := strconv.Atoi(page)

	if errOffeset == nil {
		offset = (int - 1) * offset
	}

	fmt.Println("offset")
	fmt.Println(offset)

	database.DB.Db.Offset(offset - 1).Limit(limit).Find(&products)
	var list []map[string]interface{}
	for i := 0; i < len(products); i++ {
		product := products[i]

		ds, errDs := GetDiscountbyProductID(product.ID)

		l := map[string]interface{}{
			"ID":                      product.ID,
			"Name":                    product.Name,
			"Price":                   product.Price + 0.00,
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

	if list == nil {
		list = make([]map[string]interface{}, 0)
	}

	newMap := map[string]interface{}{
		"End":      len(list) < 10,
		"Products": list,
	}
	return newMap
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
