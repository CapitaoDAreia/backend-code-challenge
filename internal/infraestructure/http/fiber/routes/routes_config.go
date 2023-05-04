package routes

import (
	"backend-challenge-api/internal/infraestructure/http/fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, controllers *controllers.APIControllers) {
	router.Get("/expression", controllers.GetExpressions)
}
