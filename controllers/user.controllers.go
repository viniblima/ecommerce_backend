package controllers

import (
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

	error := database.DB.Db.Where("email = ?", input.Email).First(&user).Error

	password := input.Password

	checked := handlers.CheckHash(user.Password, password)

	if error != nil || !checked {
		return c.Status(401).SendString("Email or password wrong")
	}

	s, err := handlers.GenerateJWT(user.ID)

	if err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}
	// c.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    s,
	// 	Expires:  time.Now().Add(7 * 24 * time.Hour),
	// 	HTTPOnly: true,
	// 	SameSite: "lax",
	// })

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
		},
	}

	return c.Status(201).JSON(newMap)
}
