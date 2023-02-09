package controllers

import (
	// "fmt"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func CreateCustomUserList(c *fiber.Ctx) error {
	list := new(models.CustomUserList)

	if str, ok := c.Locals("userID").(string); ok {

		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		list.UserID = result.ID
		fmt.Println(list.UserID)
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

	errors := handlers.ValidatePayload(list)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errors[0],
		})
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

		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		lists := handlers.GetMyLists(result.ID)

		return c.Status(200).JSON(lists)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}
}

func GetRelations(c *fiber.Ctx) error {
	if str, ok := c.Locals("userID").(string); ok {
		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		fmt.Println(result)
		lists := handlers.GetAllRelationsProductUserList()
		return c.Status(200).JSON(lists)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}
}

func AddProductToList(c *fiber.Ctx) error {
	type ProductJson struct {
		ID string `json:"ID" validate:"required"`
	}
	type Payload struct {
		Products         []ProductJson `json:"Products" validate:"required"`
		CustomUserListID string        `json:"CustomUserListID" validate:"required"`
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
	var products []models.Product
	for i := 0; i < len(payload.Products); i++ {
		id := payload.Products[i].ID
		product, err := handlers.GetProductByID(id)

		if err != nil {
			products = append(products, product)
		}

	}

	//var list models.CustomUserList
	list, err := handlers.AddProductToList(payload.CustomUserListID, products)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "List not found",
		})
	}
	// fmt.Println(err)

	return c.Status(fiber.StatusOK).JSON(list)
}
