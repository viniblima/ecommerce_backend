package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func GetHighlights(c *fiber.Ctx) error {

	products := handlers.GetHighlights()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Highlights": products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var input models.Product
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.BodyParser(&input)

	database.DB.Db.Create(&product)

	return c.Status(201).JSON(product)
}
