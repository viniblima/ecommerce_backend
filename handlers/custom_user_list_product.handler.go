package handlers

import (
	"errors"
	"fmt"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

func GetRelationProductUserList(ListID string, productID string) (models.CustomUserListProducts, error) {
	relation := models.CustomUserListProducts{}
	var err error

	dbResult := database.DB.Db.Where("custom_user_list_id = ? AND product_id = ?", ListID, productID).First(&relation)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Relation not found")
	}

	return relation, err
}

func GetAllRelationsProductUserList() []models.CustomUserListProducts {
	var relation []models.CustomUserListProducts

	database.DB.Db.Find(&relation)

	return relation
}

func CreateRelationProductUserList(ListID string, productID string) models.CustomUserListProducts {
	relation := new(models.CustomUserListProducts)
	relation.CustomUserListID = ListID
	relation.ProductID = productID

	database.DB.Db.Create(&relation)

	return *relation
}

func AddProductToList(id string, ps []models.Product) (map[string]interface{}, error) {
	var relations []models.Product

	_, err := GetListByID(id)

	if err != nil {
		err = errors.New("List not found")
		data := map[string]interface{}{}
		fmt.Println("List not found")
		return data, err
	}

	fmt.Println(len(ps))
	fmt.Println("len(ps)")
	for i := 0; i < len(ps); i++ {
		p := ps[i]
		_, err := GetRelationProductUserList(id, p.ID)

		if err != nil {
			//CreateRelationProductUserList(id, p.ID)
			CreateRelationProductUserList(id, p.ID)

			pr, _ := GetProductByID(p.ID)
			relations = append(relations, pr)
		} else {
			pr, _ := GetProductByID(p.ID)
			relations = append(relations, pr)
		}
	}

	data := map[string]interface{}{
		"CustomUserListID": id,
		"Products":         relations,
	}

	return data, err
}
