package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/handlers"
)

func SetupProductsRoutes(api fiber.Router) {
	product_routes := api.Group("/products")

	product_routes.Get("/highlights", controllers.GetHighlights)
	product_routes.Get("", controllers.GetAllProducts)
	product_routes.Get("/:id", controllers.GetProductByID)

	product_routes.Post("", handlers.VerifyJWT, controllers.CreateProduct)
	product_routes.Post("/category", handlers.VerifyJWT, controllers.AddCategoryToProduct)

}
