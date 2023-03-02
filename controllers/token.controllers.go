package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func RefreshToken(c *fiber.Ctx) error {
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

	s, err := handlers.GenerateJWT(user.ID)

	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Auth": s,
	})
}
