package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func GetHighlights(c *fiber.Ctx) error {
	var userID string

	if str, ok := c.Locals("userID").(string); ok {

		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		userID = result.ID
	}
	// } else {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "User not found",
	// 	})
	// }
	products := handlers.GetHighlights(userID)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Highlights": products,
	})
}

func GetProductByID(c *fiber.Ctx) error {
	var userID string
	if str, ok := c.Locals("userID").(string); ok {

		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		userID = result.ID
	}

	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	result, err := handlers.GetProductByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	} else {

		ds, errDs := handlers.GetDiscountbyProductID(result.ID)
		_, errLike := handlers.IsProductLiked(userID, result.ID)

		m := map[string]interface{}{
			"CreatedAt":               result.CreatedAt,
			"UpdatedAt":               result.UpdatedAt,
			"DeletedAt":               result.DeletedAt,
			"ID":                      result.ID,
			"Name":                    result.Name,
			"Price":                   result.Price,
			"Quantity":                result.Quantity,
			"MaxQuantityInstallments": result.MaxQuantityInstallments,
			"Highlight":               result.Highlight,
			"Discount":                ds,
			"Like":                    errLike == nil,
		}

		if errDs != nil {
			m["Discount"] = nil
		}
		return c.Status(fiber.StatusOK).JSON(m)
	}

}

func GetAllProducts(c *fiber.Ctx) error {
	var userID string

	if str, ok := c.Locals("userID").(string); ok {

		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		userID = result.ID
	}
	// } else {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "User not found",
	// 	})
	// }
	products := handlers.GetAllProducts(c.Query("page"), userID)
	return c.Status(fiber.StatusOK).JSON(products)
}

func LikedProducts(c *fiber.Ctx) error {
	fmt.Println("Teste 123123")
	var userID string

	if str, ok := c.Locals("userID").(string); ok {

		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		userID = result.ID

	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	ls := handlers.GetLikedProducts(userID)

	if ls == nil {
		ls = make([]map[string]interface{}, 0)
	}

	return c.Status(fiber.StatusOK).JSON(ls)
}

func CreateProduct(c *fiber.Ctx) error {

	type Payload struct {
		Name                    string  `json:"Name" validate:"required"`
		Price                   float64 `json:"Price" validate:"required"`
		Quantity                int     `json:"Quantity" validate:"required"`
		MaxQuantityInstallments int     `json:"MaxQuantityInstallments" validate:"required,min=1"`
		Highlight               bool    `json:"Highlight"`
		Categories              []struct {
			ID string `json:"ID" validate:"required"`
		} `json:"Categories" validate:"required"`
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

	if len(payload.Categories) < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Must have at least 1 category",
		})
	}

	product := new(models.Product)

	c.BodyParser(&product)

	var ms []map[string]interface{}

	for i := 0; i < len(payload.Categories); i++ {
		m := map[string]interface{}{
			"ID": payload.Categories[i].ID,
		}
		ms = append(ms, m)
	}

	p, errP := handlers.CreateProduct(product, ms)

	if errP == nil {
		return c.Status(201).JSON(p)
	} else {
		return c.Status(400).JSON(errP)
	}

}
