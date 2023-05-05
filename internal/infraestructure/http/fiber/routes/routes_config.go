package routes

import (
	"backend-challenge-api/internal/infraestructure/http/fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, controllers *controllers.APIControllers) {
	router.Post("/expressions", controllers.RegisterExpression)
	router.Get("/expressions", controllers.GetExpressions)
	router.Get("/expressions/:expressionID", controllers.CalculateExpression)
	router.Patch("/expressions/:expressionID", controllers.UpdateExpression)
	router.Delete("/expressions/:expressionID", controllers.DeleteExpression)
}
