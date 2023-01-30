package handlers

import (
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

func GetAllProducts() []models.Product {
	var products []models.Product
	database.DB.Db.Find(&products)
	return products
}

func GetProductByID(id string) models.Product {
	var product models.Product
	database.DB.Db.Where("ID = ?", id).Find(&product)
	return product
}

func CreateProduct(product *models.Product) *gorm.DB {
	return database.DB.Db.Create(&product)
}
