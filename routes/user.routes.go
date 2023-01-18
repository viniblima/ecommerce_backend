package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/controllers"
)

func SetupUserRoutes(api fiber.Router) {
	user_routes := api.Group("/user")

	user_routes.Post("/signup", controllers.SignUp)
	user_routes.Post("/signin", controllers.SignIn)
}
