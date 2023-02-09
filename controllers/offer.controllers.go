package controllers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func CreateOffer(c *fiber.Ctx) error {

	payload := models.PayloadDiscount{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.BodyParser(&payload)

	//

	var input models.Offer
	// input.Products = products

	input.EndTime = payload.EndTime

	if time.Now().After(payload.EndTime) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Date cannot be before the current date",
		})
	}

	validate := validator.New()
	var errors []string

	if err := validate.Struct(input); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			errors = append(errors, validationError.Error())
		}
	}

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors[0],
		})
	}
	database.DB.Db.Create(&input)

	products := handlers.CreateDiscountLists(input.ID, payload.List)

	return c.Status(201).JSON(fiber.Map{
		"ID":        input.ID,
		"EndTime":   input.EndTime,
		"Products":  products,
		"CreatedAt": input.CreatedAt,
	})
}

func GetAllOffers(c *fiber.Ctx) error {
	offers := handlers.GetAllOffers()
	return c.Status(fiber.StatusOK).JSON(offers)
}
