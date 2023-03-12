package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/handlers"
)

func SetupPurchaseRoutes(api fiber.Router) {
	purchase_routes := api.Group("/purchase", handlers.VerifyJWT)

	purchase_routes.Post("/", controllers.MakePurchase)
	purchase_routes.Get("/", controllers.GetMyPurchases)
}
