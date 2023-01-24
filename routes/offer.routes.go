package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viniblima/go_pq/controllers"
	"github.com/viniblima/go_pq/handlers"
)

func SetupOfferRoutes(api fiber.Router) {
	offer_routes := api.Group("/offer", handlers.VerifyJWT)

	offer_routes.Post("/", controllers.CreateOffer)
	offer_routes.Get("/", controllers.GetAllOffers)
}
