package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/procyon-projects/chrono"
	"github.com/viniblima/go_pq/database"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/models"
)

func CreateOffer(c *fiber.Ctx) error {

	payload := struct {
		Products []string `json:"products"`
		Name     string   `json:"name"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.BodyParser(&payload)

	var products []models.Product

	for i := 0; i < len(payload.Products); i++ {
		product := handlers.GetProductByID(payload.Products[i])
		products = append(products, product)
	}

	var offer models.Offer
	offer.Products = products

	timeToEnd := time.Now().Add(10 * time.Minute)

	offer.EndTime = timeToEnd

	database.DB.Db.Create(&offer)

	taskScheduler := chrono.NewDefaultTaskScheduler()

	task, err := taskScheduler.Schedule(func(ctx context.Context) {
		handlers.DeleteOffer(offer.ID)
	}, chrono.WithTime(timeToEnd))

	if err == nil {
		fmt.Println("Task has been scheduled successfully.")
	}
	fmt.Print(task)

	return c.Status(201).JSON(offer)
}

func GetAllOffers(c *fiber.Ctx) error {
	offers := handlers.GetAllOffers()
	return c.Status(fiber.StatusOK).JSON(offers)
}
