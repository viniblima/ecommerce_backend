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

func GetListBydIDAndProductRelations(userID string, listID string) map[string]interface{} {
	var list models.CustomUserList
	database.DB.Db.Model(&models.CustomUserList{}).Where("user_id = ? AND id = ?", userID, listID).Find(&list)

	var relations []models.CustomUserListProducts

	database.DB.Db.Where("custom_user_list_id = ?", list.ID).Find(&relations)

	var products []models.Product

	fmt.Println("len relations")
	fmt.Println(len(relations))
	for j := 0; j < len(relations); j++ {

		product, err := GetProductByID(relations[j].ProductID)

		if err == nil {
			products = append(products, product)
		}

	}

	if products == nil {
		products = make([]models.Product, 0)
	}

	local := map[string]interface{}{
		"CustomUserList": map[string]interface{}{
			"ID":        list.ID,
			"CreatedAt": list.CreatedAt,
			"UpdatedAt": list.UpdatedAt,
			"Name":      list.Name,
		},
		"Products": products,
	}

	return local
}

func GetMyLists(id string) []map[string]interface{} {
	var list []models.CustomUserList
	var result []map[string]interface{}

	database.DB.Db.Model(&models.CustomUserList{}).Where("user_id = ?", id).Find(&list)

	for i := 0; i < len(list); i++ {
		//relations := GetRelationsByListID(list[i].ID)
		var relations []models.CustomUserListProducts
		database.DB.Db.Where("custom_user_list_id = ?", list[i].ID).Find(&relations)

		var products []map[string]interface{}

		fmt.Println("len relations")
		fmt.Println(len(relations))
		for j := 0; j < len(relations); j++ {

			product, err := GetProductByID(relations[j].ProductID)

			if err == nil {
				_, errLike := IsProductLiked(id, product.ID)

				m := map[string]interface{}{
					"ID": product.ID,
					// "CreatedAt": "2023-02-22T18:40:51.657414Z",
					// "UpdatedAt": "2023-02-22T18:40:51.657414Z",
					"Name":                    product.Name,
					"Price":                   product.Price,
					"Quantity":                product.Quantity,
					"MaxQuantityInstallments": product.MaxQuantityInstallments,
					"Highlight":               product.Highlight,
					"Like":                    errLike == nil,
				}
				products = append(products, m)
			}

		}

		if products == nil {
			products = make([]map[string]interface{}, 0)
		}
		local := map[string]interface{}{
			"CustomUserList": map[string]interface{}{
				"ID":        list[i].ID,
				"CreatedAt": list[i].CreatedAt,
				"UpdatedAt": list[i].UpdatedAt,
				"Name":      list[i].Name,
			},
			"Products": products,
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
