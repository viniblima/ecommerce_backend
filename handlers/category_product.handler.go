package handlers

import (
	"errors"

	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

func GetAllCategoriesOfProduct(productID string) ([]models.Category, error) {
	var rls []models.CategoryProduct
	var err error

	dbResult := database.DB.Db.Where("product_id = ?", productID).Find(&rls)
	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Relation not found")
	}

	var cs []models.Category
	if err == nil {
		for i := 0; i < len(rls); i++ {
			c, _ := GetCategoryByID(rls[i].CategoryID)
			cs = append(cs, c)
		}

	}
	return cs, err
}

func CreateRelationCategoryProduct(categoryID string, productID string) (models.CategoryProduct, error) {
	var r models.CategoryProduct
	var err error
	r.CategoryID = categoryID
	r.ProductID = productID

	dbResult := database.DB.Db.Create(&r)
	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Error on create like")
	}

	return r, err
}

func DeleteRelationCategoryProduct(id string) (models.CategoryProduct, error) {
	var r models.CategoryProduct
	var err error

	dbResult := database.DB.Db.Where("id = ?", id).Delete(&r)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Error on delete relation")
	}

	return r, err
}

func DeleteAllRelationCategoryProduct(productID string) ([]models.CategoryProduct, error) {
	var r []models.CategoryProduct
	var err error

	dbResult := database.DB.Db.Where("product_id = ?", productID).Delete(&r)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Error on delete relation")
	}

	return r, err
}

func CreateListRelationCategoryProduct(categoryIDs []map[string]interface{}, productID string) []map[string]interface{} {
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

func GetRelationByCategoryID(categoryID string) ([]models.CategoryProduct, error) {
	relations := []models.CategoryProduct{}

	var err error

	dbResult := database.DB.Db.Where("category_id = ?", categoryID).Find(&relations)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Relation not found")
	}

	return relations, err
}
