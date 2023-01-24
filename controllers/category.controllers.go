package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
)

func GetAllCategories(c *fiber.Ctx) error {
	var categories []models.Category

	database.DB.Db.Find(&categories)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Categories": categories,
	})
}

func CreateCategory(c *fiber.Ctx) error {
	var input models.Category
	category := new(models.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.BodyParser(&input)

	database.DB.Db.Create(&category)

	return c.Status(201).JSON(fiber.Map{
		"ID":        category.ID,
		"CreatedAt": category.CreatedAt,
		"UpdatedAt": category.UpdatedAt,
		"Name":      category.Name,
	})
}
