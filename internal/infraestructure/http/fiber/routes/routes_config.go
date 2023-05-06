package routes

import (
	"backend-challenge-api/internal/infraestructure/http/fiber/controllers"
	"backend-challenge-api/internal/infraestructure/http/fiber/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, controllers *controllers.APIControllers) {
	unprotectedRoutes := router.Group("")
	unprotectedRoutes.Post("/auth", controllers.Authenticate)

	protectedRoutes := router.Group("")
	protectedRoutes.Use(middlewares.Authenticate)

	protectedRoutes.Post("/expressions", controllers.RegisterExpression)
	protectedRoutes.Get("/expressions", controllers.GetExpressions)
	protectedRoutes.Get("/expressions/:expressionID", controllers.CalculateExpression)
	protectedRoutes.Patch("/expressions/:expressionID", controllers.UpdateExpression)
	protectedRoutes.Delete("/expressions/:expressionID", controllers.DeleteExpression)

}
