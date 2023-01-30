package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func GetHighlights(c *fiber.Ctx) error {

	products := handlers.GetHighlights()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Highlights": products,
	})
}

func GetAllProducts(c *fiber.Ctx) error {

	products := handlers.GetAllProducts()
	return c.Status(fiber.StatusOK).JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var input models.Product
	validate := validator.New()
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.BodyParser(&input)
	var errors []string

	if err := validate.Struct(product); err != nil {
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

	handlers.CreateProduct(product)

	return c.Status(201).JSON(product)
}
