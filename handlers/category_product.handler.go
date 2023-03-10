package handlers

import (
	"errors"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

func CreateRelationCategoryProduct(categoryIDs []map[string]interface{}, productID string) []map[string]interface{} {
	var rls []map[string]interface{}
	for i := 0; i < len(categoryIDs); i++ {
		var item string
		item = categoryIDs[i]["ID"].(string)
		rl, errRL := GetRelationByProductIDAndCategoryID(productID, item)

		if errRL == nil {
			c, _ := GetCategoryByID(rl.CategoryID)
			rls = append(rls, map[string]interface{}{
				"ID":   c.ID,
				"Name": c.Name,
			})
		} else {
			c, errC := GetCategoryByID(item)

			if errC == nil {

				rp, errRP := GetProductByID(productID)

				if errRP == nil {
					relation := new(models.CategoryProduct)
					relation.CategoryID = c.ID
					relation.ProductID = rp.ID

					database.DB.Db.Create(&relation)

					rls = append(rls, map[string]interface{}{
						"ID":   c.ID,
						"Name": c.Name,
					})
				}
			}
		}

	}

	return rls
}

func GetRelationByProductIDAndCategoryID(productID string, categoryID string) (models.CategoryProduct, error) {
	relation := models.CategoryProduct{}

	var err error

	dbResult := database.DB.Db.Where("product_id = ? AND category_id = ?", productID, categoryID).First(&relation)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Relation not found")
	}

	return relation, err
}
