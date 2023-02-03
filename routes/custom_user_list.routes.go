package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/handlers"
)

func SetupCustomUserListRoutes(api fiber.Router) {
	list_routes := api.Group("my_lists", handlers.VerifyJWT)

	list_routes.Post("", controllers.CreateCustomUserList)
	list_routes.Get("", controllers.GetMyLists)
	list_routes.Post("/add_product", controllers.AddProductToList)

}
