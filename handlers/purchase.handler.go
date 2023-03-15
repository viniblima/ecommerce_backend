package handlers

import (
	"errors"
	"os"
	"strconv"

	"github.com/viniblima/go_pq/cielo"
	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
	"gorm.io/gorm"
)

func IntegratePurchase(u models.User, p models.PaymentMethod, c models.Card, save bool, i uint32, amount uint32) (models.Purchase, error) {
	var e error
	CieloClient, err := cielo.NewClient(os.Getenv("MERCHANT_ID"), os.Getenv("MERCHANT_KEY"), cielo.SandboxEnvironment)
	if err != nil {
		e = err
	}
	CieloClient.Log = os.Stdout

	cc := cielo.CreditCard{
		CardNumber:     c.CardNumber,
		Holder:         c.Holder,
		CustomerName:   c.CustomerName,
		ExpirationDate: c.ExpirationDate,
		SecurityCode:   c.SecurityCode,
		SaveCard:       save,
		Brand:          c.Brand,
	}

	storedCC, err := CieloClient.CreateTokenizeCard(&cc)

	if err != nil {
		e = err
	}

	purchase := CreatePurchase(u.ID, p)

	if p == models.CreditCard {
		sale := cielo.Sale{
			MerchantOrderID: purchase.ID,
			Customer: &cielo.Customer{
				Email: u.Email,
				Name:  u.Name,
			},
			Payment: &cielo.Payment{
				Installments: i,
				Amount:       amount,
				Type:         "CreditCard",
				// Pensar em alguma descricao
				SoftDescriptor:   u.ID,
				ServiceTaxAmount: 0,
				CreditCard: &cielo.CreditCard{
					CardToken: storedCC.CardToken, SecurityCode: c.SecurityCode,
				},
			},
		}

		_, err := CieloClient.Authorization(&sale)
		if err != nil {
			e = err
		} else {
			newPurchase, _ := GetPurchaseByID(purchase.ID)

			newPurchase.Success = true
			UpdatePurchase(newPurchase)
		}

		purchase, _ := GetPurchaseByID(purchase.ID)
		return purchase, e
	} else {
		sale := cielo.Sale{
			MerchantOrderID: purchase.ID,
			Customer: &cielo.Customer{
				Name:         u.Name,
				Identity:     "44255170835",
				IdentityType: "CPF",
			},
			Payment: &cielo.Payment{
				Amount: amount,
				Type:   "Pix",
			},
		}

		_, err := CieloClient.Authorization(&sale)
		if err != nil {
			e = err
		} else {
			newPurchase, _ := GetPurchaseByID(purchase.ID)

			newPurchase.Success = true
			UpdatePurchase(newPurchase)
		}

		purchase, _ := GetPurchaseByID(purchase.ID)
		return purchase, e
	}
}

func GetMyPurchases(id string, page string) map[string]interface{} {
	var ps []models.Purchase

	if page == "" {
		page = "1"
	}
	offset := 10
	limit := 10

	int, errOffeset := strconv.Atoi(page)

	if errOffeset == nil {
		offset = (int - 1) * offset
	}

	var result []map[string]interface{}
	database.DB.Db.Offset(offset).Limit(limit).Where("user_id = ?", id).Find(&ps)

	for i := 0; i < len(ps); i++ {
		p := ps[i]

		r := map[string]interface{}{
			"CreatedAt": p.CreatedAt,
			"ID":        p.ID,
			"Success":   p.Success,
		}
		var rl []models.PurchaseProduct

		database.DB.Db.Where("purchase_id = ?", p.ID).Find(&rl)

		var prl []map[string]interface{}

		for j := 0; j < len(rl); j++ {
			product, _ := GetProductByID(rl[j].ProductID)
			ds, errDs := GetDiscountbyProductID(product.ID)
			cs, _ := GetAllCategoriesOfProduct(product.ID)
			_, errLike := IsProductLiked(id, product.ID)
			m := map[string]interface{}{
				"Highlight":               product.Highlight,
				"ID":                      product.ID,
				"Like":                    errLike == nil,
				"MaxQuantityInstallments": product.MaxQuantityInstallments,
				"Name":                    product.Name,
				"Price":                   product.Price,
				"Quantity":                product.Quantity,
				"Discount":                ds,
				"Categories":              cs,
			}

			if errDs != nil {
				m["Discount"] = nil
			}
			if len(cs) < 1 {
				m["Categories"] = make([]models.Category, 0)
			}

			purchaseProduct := map[string]interface{}{
				"Product": m,
				"Status":  rl[j].Status,
				"ID":      rl[j].PurchaseID,
			}
			prl = append(prl, purchaseProduct)
		}

		r["PurchaseProducts"] = prl

		if len(prl) < 1 {
			r["PurchaseProducts"] = make([]map[string]interface{}, 0)
		}

		result = append(result, r)
	}
	newMap := map[string]interface{}{
		"End":       len(result) < 10,
		"Purchases": result,
	}

	if len(result) < 1 {
		newMap["Purchases"] = make([]map[string]interface{}, 0)
	}

	return newMap
}

func CreatePurchase(userID string, m models.PaymentMethod) models.Purchase {
	var purchase models.Purchase

	purchase.UserID = userID

	purchase.PaymentMethod = m

	database.DB.Db.Create(&purchase)

	return purchase
}

func UpdatePurchase(p models.Purchase) models.Purchase {
	database.DB.Db.Save(&p)

	return p
}

func GetPurchaseByID(id string) (models.Purchase, error) {
	var purchase models.Purchase

	var err error

	dbResult := database.DB.Db.Where("ID = ?", id).First(&purchase)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		err = errors.New("Purchase not found")
	}

	return purchase, err
}

func CreateRelationProductPurchase(purchaseID string, productID string) models.PurchaseProduct {
	var rl models.PurchaseProduct

	rl.PurchaseID = purchaseID
	rl.ProductID = productID

	database.DB.Db.Create(&rl)

	return rl
}

/*
* 1 - Criar a compra
* 2 - Criar as relacoes
* 3 - Diminuir a quantidade de produtos
 */
// func MakePurchase(ps []models.Product) (map[string]interface{}, error) {
// 	IntegratePurchase()
// }
