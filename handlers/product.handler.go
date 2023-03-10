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

func GetHighlights(userID string) []map[string]interface{} {
	var products []models.Product
	database.DB.Db.Where("highlight = ?", true).Find(&products)
	var list []map[string]interface{}
	for i := 0; i < len(products); i++ {
		product := products[i]

		ds, errDs := GetDiscountbyProductID(product.ID)
		cs, _ := GetAllCategoriesOfProduct(product.ID)
		_, errLike := IsProductLiked(userID, product.ID)

		l := map[string]interface{}{
			"ID":                      product.ID,
			"Name":                    product.Name,
			"Price":                   product.Price + 0.00,
			"Quantity":                product.Quantity,
			"MaxQuantityInstallments": product.MaxQuantityInstallments,
			"Highlight":               product.Highlight,
			"Like":                    errLike == nil,
			"Discount":                ds,
			"Categories":              cs,
		}

		if errDs != nil {
			l["Discount"] = nil
		}

		if len(cs) < 1 {
			l["Categories"] = make([]map[string]interface{}, 0)
		}

		list = append(list, l)
	}

	return list
}
func GetLikedProducts(userID string) []map[string]interface{} {
	var likes []models.Like

	var list []map[string]interface{}

	database.DB.Db.Where("user_id = ?", userID).Find(&likes)
	fmt.Println("len(likes)")
	fmt.Println(len(likes))
	for i := 0; i < len(likes); i++ {
		like := likes[i]

		p, _ := GetProductByID(like.ProductID)
		cs, _ := GetAllCategoriesOfProduct(like.ProductID)
		ds, errDs := GetDiscountbyProductID(p.ID)

		m := map[string]interface{}{
			"Discount":                ds,
			"ID":                      p.ID,
			"Name":                    p.Name,
			"Price":                   p.Price + 0.00,
			"Quantity":                p.Quantity,
			"MaxQuantityInstallments": p.MaxQuantityInstallments,
			"Highlight":               p.Highlight,
			"Categories":              cs,
		}

		if errDs != nil {
			m["Discount"] = nil
		}

		if len(cs) < 1 {
			m["Categories"] = make([]map[string]interface{}, 0)
		}

		list = append(list, m)
	}
	return list
}

func GetAllProducts(page string, userID string, categories []string) map[string]interface{} {
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
	if len(categories) > 0 {
		var ids []string
		for i := 0; i < len(categories); i++ {
			c, _ := GetCategoryByID(categories[i])
			rls, errRLS := GetRelationByCategoryID(c.ID)

			if errRLS == nil {
				for j := 0; j < len(rls); j++ {
					ids = append(ids, rls[j].ProductID)
				}

			}
		}

		database.DB.Db.Offset(offset).Limit(limit).Where("id in ?", ids).Find(&products)
	} else {
		database.DB.Db.Offset(offset).Limit(limit).Find(&products)
	}

	var list []map[string]interface{}
	for i := 0; i < len(products); i++ {
		product := products[i]

		ds, errDs := GetDiscountbyProductID(product.ID)
		cs, _ := GetAllCategoriesOfProduct(product.ID)
		_, errLike := IsProductLiked(userID, product.ID)

		l := map[string]interface{}{
			"ID":                      product.ID,
			"Name":                    product.Name,
			"Price":                   product.Price + 0.00,
			"Quantity":                product.Quantity,
			"MaxQuantityInstallments": product.MaxQuantityInstallments,
			"Highlight":               product.Highlight,
			"Like":                    errLike == nil,
			"Discount":                ds,
			"Categories":              cs,
		}

		if errDs != nil {
			l["Discount"] = nil
		}

		if len(cs) < 1 {
			l["Categories"] = make([]map[string]interface{}, 0)
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

func CreateProduct(product *models.Product, categories []map[string]interface{}) (map[string]interface{}, error) {

	var err error

	dbResult := database.DB.Db.Create(&product)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Relation not found")
	}

	rl := CreateListRelationCategoryProduct(categories, product.ID)

	m := map[string]interface{}{
		"CreatedAt":               product.CreatedAt,
		"ID":                      product.ID,
		"Name":                    product.Name,
		"Price":                   product.Price,
		"Quantity":                product.Quantity,
		"MaxQuantityInstallments": product.MaxQuantityInstallments,
		"Highlight":               product.Highlight,
		"Categories":              rl,
	}

	return m, err
}

func UpdateProduct(p models.Product) models.Product {

	database.DB.Db.Save(&p)

	return p
}
