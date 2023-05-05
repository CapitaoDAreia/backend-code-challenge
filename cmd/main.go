package main

import (
	"backend-challenge-api/internal/infraestructure/database/postgres"
	"backend-challenge-api/internal/infraestructure/database/postgres/repositories"
	"backend-challenge-api/internal/infraestructure/http/fiber/controllers"
	"backend-challenge-api/internal/infraestructure/http/fiber/routes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	db := postgres.Connect()

	expressionsRepository := repositories.NewExpressionsRepository(db)
	expressionsController := controllers.NewAPIControllers(*expressionsRepository)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cors.New())
	app.Use(recover.New())

	api := app.Group("")

	routes.SetupRoutes(api, expressionsController)

	err := app.Listen(":4000")

	if err != nil {
		panic(err)
	}
}
