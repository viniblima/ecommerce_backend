package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
)

func CreateLikeProduct(c *fiber.Ctx) error {
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

	print(userID)

	type Payload struct {
		ID string `json:"ID" validate:"required"`
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

	l, err := handlers.IsProductLiked(userID, payload.ID)

	if err == nil {
		_, errDelete := handlers.DeleteLike(l.ID)

		if errDelete != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": errDelete,
			})
		}

		return c.Status(200).JSON(fiber.Map{"Message": "Deleted"})

	} else {
		create, errCreate := handlers.CreateLike(userID, payload.ID)

		if errCreate != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": errCreate,
			})
		}

		return c.Status(200).JSON(create)
	}

}
