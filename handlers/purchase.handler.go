package handlers

import (
	"os"

	"github.com/viniblima/go_pq/cielo"
	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
)

func IntegratePurchase(
	u models.User,
	p models.PaymentMethod,
	c models.Card,
	save bool,
	i uint32,
	amount uint32,
) (*cielo.Sale, error) {
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

	purchase := new(models.Purchase)

	purchase.UserID = u.ID

	purchase.PaymentMethod = p

	database.DB.Db.Create(&purchase)

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

		resp, err := CieloClient.Authorization(&sale)
		if err != nil {
			e = err
		}

		return resp, e
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

		resp, err := CieloClient.Authorization(&sale)
		if err != nil {
			e = err
		}

		return resp, e
	}

}

/*
* 1 - Criar a compra
* 2 - Criar as relacoes
* 3 - Diminuir a quantidade de produtos
 */
// func MakePurchase(ps []models.Product) (map[string]interface{}, error) {
// 	IntegratePurchase()
// }
