package main

import (
	"backend-challenge-api/internal/infraestructure/http/fiber/controllers"
	"backend-challenge-api/internal/infraestructure/http/fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())

	api := app.Group("")

	getExpressionController := controllers.NewAPIControllers()

	routes.SetupRoutes(api, getExpressionController)

	err := app.Listen(":4000")

	if err != nil {
		panic(err)
	}

}
