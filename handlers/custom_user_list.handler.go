package handlers

import (
	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

func CreateUserList(l *models.CustomUserList) *gorm.DB {
	return database.DB.Db.Create(&l)
}

func GetMyLists(id string) []models.CustomUserList {
	var list []models.CustomUserList
	database.DB.Db.Where("user_id = ?", id).Preload("Products").Find(&list)

	return list
}
