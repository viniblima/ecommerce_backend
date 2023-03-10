package handlers

import (
	"errors"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

func GetCategoryByID(id string) (models.Category, error) {
	var c models.Category

	var err error

	dbResult := database.DB.Db.Where("ID = ?", id).First(&c)
	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("category not found")
	}

	return c, err
}
