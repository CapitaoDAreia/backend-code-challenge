package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type APIControllers struct {
}

func NewAPIControllers() *APIControllers {
	return &APIControllers{}
}

func (controller *APIControllers) GetExpressions(ctx *fiber.Ctx) error {
	return ctx.JSON("Expressions")
}
