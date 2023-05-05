package controllers

import (
	"backend-challenge-api/internal/domain/entities"
	"backend-challenge-api/internal/infraestructure/database/postgres/repositories"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type APIControllers struct {
	r repositories.ExpressionsRepository
}

func NewAPIControllers(r repositories.ExpressionsRepository) *APIControllers {
	return &APIControllers{
		r,
	}
}

func (controller *APIControllers) RegisterExpression(ctx *fiber.Ctx) error {
	receivedExpression := new(entities.Expression)
	if err := ctx.BodyParser(receivedExpression); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createdExpressionId, err := controller.r.RegisterExpression(receivedExpression)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(200).JSON(createdExpressionId)
}

func (controller *APIControllers) GetExpressions(ctx *fiber.Ctx) error {
	return ctx.JSON("Expressions")
}

func (controller *APIControllers) CalculateExpression(ctx *fiber.Ctx) error {
	expressionID := ctx.Params("expressionId")

	xValue := ctx.Query("x")
	yValue := ctx.Query("y")
	zValue := ctx.Query("z")

	type receivedInfoType struct {
		ID string `json:"id"`
		X  string `json:"x"`
		Y  string `json:"y"`
		Z  string `json:"z"`
	}

	var receivedInfo receivedInfoType

	receivedInfo.ID = expressionID
	receivedInfo.X = xValue
	receivedInfo.Y = yValue
	receivedInfo.Z = zValue

	fmt.Println(receivedInfo)

	return ctx.JSON(receivedInfo)
}

func (controller *APIControllers) UpdateExpression(ctx *fiber.Ctx) error {
	return ctx.JSON("Expressions")
}
