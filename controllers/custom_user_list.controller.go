package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func CreateCustomUserList(c *fiber.Ctx) error {
	validate := validator.New()
	list := new(models.CustomUserList)

	if str, ok := c.Locals("userID").(string); ok {

		list.UserID = handlers.GetUserByID(str).ID
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := c.BodyParser(list); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.BodyParser(&list)

	var errors []string

	if err := validate.Struct(list); err != nil {
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

	handlers.CreateUserList(list)

	return c.Status(201).JSON(list)
}

func GetMyLists(c *fiber.Ctx) error {
	if str, ok := c.Locals("userID").(string); ok {

		id := handlers.GetUserByID(str).ID

		lists := handlers.GetMyLists(id)

		return c.Status(200).JSON(lists)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}
}
