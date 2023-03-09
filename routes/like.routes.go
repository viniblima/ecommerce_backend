package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/handlers"
)

func SetupLikeRoutes(api fiber.Router) {
	like_routes := api.Group("/like", handlers.VerifyJWT)

	like_routes.Post("", handlers.VerifyJWT, controllers.CreateLikeProduct)
	like_routes.Get("", handlers.VerifyJWT, controllers.LikedProducts)
}
