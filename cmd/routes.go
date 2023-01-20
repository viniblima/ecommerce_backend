package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/routes"
)

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")

	setupV1Routes(api)

}

func setupV1Routes(api fiber.Router) {
	v1 := api.Group("/v1")

	routes.SetupUserRoutes(v1)
	routes.SetupProductsRoutes(v1)
	routes.SetupCategoryRoutes(v1)
}
