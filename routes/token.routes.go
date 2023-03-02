package routes

import (
	"github.com/gofiber/fiber/v2"
	//"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/handlers"
)

func SetupTokenRoutes(api fiber.Router) {
	token_routes := api.Group("/token", handlers.RefreshToken)

	token_routes.Post("/refresh", controllers.RefreshToken)
}
