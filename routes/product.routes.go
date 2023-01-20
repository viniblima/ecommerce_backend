package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/handlers"
)

func SetupProductsRoutes(api fiber.Router) {
	product_routes := api.Group("/products", handlers.VerifyJWT)

	product_routes.Get("/highlights", controllers.GetHighlights)
	product_routes.Post("", controllers.CreateProduct)
}
