package handlers

import (
	"errors"
	"fmt"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

func CreateUserList(l *models.CustomUserList) *gorm.DB {
	return database.DB.Db.Create(&l)
}

func GetMyLists(id string) []map[string]interface{} {
	var list []models.CustomUserList
	var result []map[string]interface{}

	database.DB.Db.Model(&models.CustomUserList{}).Where("user_id = ?", id).Find(&list)

	for i := 0; i < len(list); i++ {
		//relations := GetRelationsByListID(list[i].ID)
		var relations []models.CustomUserListProducts
		database.DB.Db.Where("custom_user_list_id = ?", list[i].ID).Find(&relations)
		fmt.Println("tamanho das relacoes")
		fmt.Println(len(relations))
		var products []models.Product
		fmt.Println(len(relations))
		for j := 0; j < len(relations); j++ {
			products = append(products, GetProductByID(relations[j].ProductID))
		}
		local := map[string]interface{}{
			"CustomUserList": list[i],
			"Products":       products,
		}
		result = append(result, local)
	}
	return result
}

func GetListByID(id string) (models.CustomUserList, error) {
	var list models.CustomUserList
	var err error
	dbResult := database.DB.Db.Where("id = ?", id).First(&list)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("user not found")
	}

	return list, err
}
