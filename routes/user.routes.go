package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/handlers"
)

func SetupUserRoutes(api fiber.Router) {
	user_routes := api.Group("/user")

	user_routes.Post("/signup", controllers.SignUp)
	user_routes.Post("/signin", controllers.SignIn)
	user_routes.Post("/first_access", handlers.FirstAccess, controllers.SignUp)

	user_routes.Post("/signin_portal", controllers.SignInPortal)
}
