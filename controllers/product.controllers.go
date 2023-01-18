package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/models"
)

func GetAllCategories(c *fiber.Ctx) error {
	var categories models.Category

	database.DB.Db.Where("ID = ?", "").First(categories)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"categories": categories,
	})
}
