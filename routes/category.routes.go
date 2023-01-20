package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/handlers"
)

func SetupCategoryRoutes(api fiber.Router) {
	category_routes := api.Group("/category", handlers.VerifyJWT)

	category_routes.Get("/", controllers.GetAllCategories)
	category_routes.Post("/create", controllers.CreateCategory)
}
