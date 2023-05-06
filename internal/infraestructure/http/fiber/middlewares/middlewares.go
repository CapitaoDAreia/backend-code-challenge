package middlewares

import (
	"backend-challenge-api/internal/infraestructure/http/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing Authorization header"})
	}

	token := strings.Replace(authHeader, "Bearer ", "", 1)

	if err := auth.ValidateToken(token); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Next()
}
