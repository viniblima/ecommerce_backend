package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
)

func GetAllDiscounts(c *fiber.Ctx) error {
	discounts := handlers.GetAllDiscounts()

	return c.Status(fiber.StatusOK).JSON(discounts)
}

func GetAllDiscountsProducts(c *fiber.Ctx) error {
	discounts := handlers.GetAllProductDiscounts()

	return c.Status(fiber.StatusOK).JSON(discounts)
}
