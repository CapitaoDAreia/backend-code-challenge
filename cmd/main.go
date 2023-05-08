package main

import (
	"backend-challenge-api/internal/application/services"
	"backend-challenge-api/internal/infraestructure/configuration"
	"backend-challenge-api/internal/infraestructure/database/postgres"
	"backend-challenge-api/internal/infraestructure/database/postgres/repositories"
	"backend-challenge-api/internal/infraestructure/http/fiber/controllers"
	"backend-challenge-api/internal/infraestructure/http/fiber/routes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	configuration.GenerateSecretKey()
}

func main() {
	db := postgres.Connect()

	apiRepository := repositories.NewExpressionsRepository(db)
	apiService := services.NewExpressionsServices(*apiRepository)
	expressionsController := controllers.NewAPIControllers(apiService)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cors.New())
	app.Use(recover.New())

	routes.SetupRoutes(app, expressionsController)

	err := app.Listen(":4000")

	if err != nil {
		panic(err)
	}
}
