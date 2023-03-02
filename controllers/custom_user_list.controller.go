package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func CreateCustomUserList(c *fiber.Ctx) error {

	type Payload struct {
		Name     string `json:"Name" validate:"required"`
		Products []struct {
			ID string `json:"ID" validate:"required"`
		} `json:"Products"`
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

	list := new(models.CustomUserList)
	c.BodyParser(&list)

	if str, ok := c.Locals("userID").(string); ok {

		result, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		list.UserID = result.ID

	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	fmt.Println(payload.Products != nil)
	handlers.CreateUserList(list)

	if payload.Products != nil {

		for i := 0; i < len(payload.Products); i++ {
			p := payload.Products[i]

			handlers.CreateRelationProductUserList(list.ID, p.ID)
		}
	}

	result := handlers.GetListBydIDAndProductRelations(list.UserID, list.ID)

	return c.Status(201).JSON(result)
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

		if lists == nil {
			lists = make([]map[string]interface{}, 0)
		}
		return c.Status(200).JSON(lists)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}
}

func GetRelations(c *fiber.Ctx) error {
	if str, ok := c.Locals("userID").(string); ok {
		_, err := handlers.GetUserByID(str)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}

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

	fmt.Println("len relations")
	fmt.Println(len(payload.Products))
	for i := 0; i < len(payload.Products); i++ {
		id := payload.Products[i].ID
		product, err := handlers.GetProductByID(id)

		if err == nil {
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
	//

	return c.Status(fiber.StatusOK).JSON(list)
}
