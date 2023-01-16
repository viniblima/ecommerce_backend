package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/handlers"
	"github.com/viniblima/go_pq/routes"
)

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/", handlers.ListFacts)
	api.Post("/fact", handlers.CreateFact)

	setupV1Routes(api)

}

func setupV1Routes(api fiber.Router) {
	v1 := api.Group("/v1")

	routes.SetupUserRoutes(v1)
}
