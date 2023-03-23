package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/viniblima/go_pq/routes"
)

func setupRoutes(app *fiber.App) {
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	api := app.Group("/api")

	setupV1Routes(api)

}

func setupV1Routes(api fiber.Router) {
	v1 := api.Group("/v1")

	routes.SetupUserRoutes(v1)
	routes.SetupTokenRoutes(v1)
	routes.SetupProductsRoutes(v1)
	routes.SetupCategoryRoutes(v1)
	routes.SetupOfferRoutes(v1)
	routes.SetupCustomUserListRoutes(v1)
	routes.SetupPurchaseRoutes(v1)
	routes.SetupLikeRoutes(v1)
}
