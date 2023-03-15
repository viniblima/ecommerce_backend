package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func MakePurchase(c *fiber.Ctx) error {

	var user models.User

	if str, ok := c.Locals("userID").(string); ok {

		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		user = result

	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	type Payload struct {
		PaymentMethodPayload struct {
			Name         uint32 `json:"Name" validate:"required"`
			Installments uint32 `json:"Installments" validate:"required"`
		} `json:"PaymentMethod" validate:"required"`
		ProductsPayload []struct {
			ID       string `json:"ID" validate:"required"`
			Quantity uint32 `json:"Quantity" validate:"required"`
		} `json:"Products" validate:"required"`
	}

	payload := Payload{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.BodyParser(&payload)

	errors := handlers.ValidatePayload(payload)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors[0],
		})
	}

	var ps []models.Product

	amount := 0

	for i := 0; i < len(payload.ProductsPayload); i++ {
		pl := payload.ProductsPayload[i]
		//id := payload.ProductsPayload[i].ID

		p, err := handlers.GetProductByID(pl.ID)

		if err != nil || (p.Quantity-int(pl.Quantity)) < 0 {
			errors = append(errors, "Product not found or unavailable in this quantity")
		} else {
			ps = append(ps, p)

			d, e := handlers.GetDiscountbyProductID(p.ID)

			if e == nil {
				amount += int(d.PriceWithDiscount)
			} else {
				amount += int(p.Price)
			}

		}
	}

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors[0],
		})
	}
	card := models.Card{
		CardNumber:     "2307449183305478",
		CustomerName:   "Vinícius Branco",
		Holder:         "VINICIUS B LIMA",
		ExpirationDate: "04/2026",
		SecurityCode:   "275",
		Brand:          "mastercard",
	}

	purchase, err := handlers.IntegratePurchase(user, models.CreditCard, card, false, payload.PaymentMethodPayload.Installments, uint32(amount))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	fmt.Println(purchase)
	for i := 0; i < len(payload.ProductsPayload); i++ {
		pl := payload.ProductsPayload[i]
		p, _ := handlers.GetProductByID(pl.ID)
		p.Quantity -= int(pl.Quantity)
		handlers.CreateRelationProductPurchase(purchase.ID, pl.ID)
		handlers.UpdateProduct(p)
	}

	return c.Status(200).JSON(purchase)
}

func GetMyPurchases(c *fiber.Ctx) error {
	var user models.User

	if str, ok := c.Locals("userID").(string); ok {

		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		user = result

	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	result := handlers.GetMyPurchases(user.ID, c.Query("page"))

	return c.Status(200).JSON(result)
}
