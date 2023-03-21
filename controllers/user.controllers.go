package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
	"github.com/viniblima/go_pq/util"
)

func SignIn(c *fiber.Ctx) error {
	var input models.User

	err := c.BodyParser(&input)
	if err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	var user models.User
	input.Email = util.NormalizeEmail(input.Email)

	u, errorU := handlers.GetByEmail(input.Email, false)

	user = u

	fmt.Println(user.IsAdmin)
	fmt.Println(errorU)

	password := input.Password

	checked := handlers.CheckHash(user.Password, password)

	if errorU != nil || !checked {
		return c.Status(401).SendString("Email or password wrong")
	}

	s, err := handlers.GenerateJWT(user.ID)

	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Auth": s,
		"User": fiber.Map{
			"ID":        user.ID,
			"CreatedAt": user.CreatedAt,
			"UpdatedAt": user.UpdatedAt,
			"Name":      user.Name,
			"Email":     user.Email,
		},
	})
}

func SignInPortal(c *fiber.Ctx) error {
	var input models.User

	err := c.BodyParser(&input)
	if err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	var user models.User
	input.Email = util.NormalizeEmail(input.Email)

	u, errorU := handlers.GetByEmail(input.Email, true)

	user = u

	fmt.Println(user.IsAdmin == false)

	password := input.Password

	checked := handlers.CheckHash(user.Password, password)

	if errorU != nil || !checked {
		return c.Status(401).JSON(fiber.Map{"message": "Email or password wrong"})
	}

	if !user.IsAdmin {
		return c.Status(401).JSON(fiber.Map{"message": "User is not an admin"})
	}

	s, err := handlers.GenerateJWT(user.ID)

	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Auth": s,
		"User": fiber.Map{
			"ID":        user.ID,
			"CreatedAt": user.CreatedAt,
			"UpdatedAt": user.UpdatedAt,
			"Name":      user.Name,
			"Email":     user.Email,
		},
	})
}

func SignUp(c *fiber.Ctx) error {
	validate := validator.New()
	var input models.User
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.BodyParser(&input)
	input.Email = util.NormalizeEmail(input.Email)

	error := database.DB.Db.Where("email = ?", input.Email).First(&user).Error

	if error == nil {
		return c.Status(401).JSON(fiber.Map{"message": "Email already registered"})
	}

	var errors []string

	if err := validate.Struct(user); err != nil {
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
	hashed, _ := handlers.HashPassword(user.Password)

	user.Password = hashed

	user.IsAdmin = false

	if str, ok := c.Locals("firstAccess").(string); ok {
		fmt.Println(str)
		user.IsAdmin = true
	}

	database.DB.Db.Create(&user)

	s, err := handlers.GenerateJWT(user.ID)

	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	newMap := map[string]interface{}{
		"Auth": s,
		"User": map[string]interface{}{
			"ID":        user.ID,
			"CreatedAt": user.CreatedAt,
			"UpdatedAt": user.UpdatedAt,
			"Name":      user.Name,
			"Email":     user.Email,
			"IsAdmin":   user.IsAdmin,
		},
	}

	return c.Status(201).JSON(newMap)
}
