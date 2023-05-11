package controllers

import (
	"backend-challenge-api/internal/application/services"
	"backend-challenge-api/internal/domain/entities"
	"backend-challenge-api/internal/infraestructure/http/auth"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type APIControllers struct {
	s services.ExpressionServices
}

func NewAPIControllers(s services.ExpressionServices) *APIControllers {
	return &APIControllers{
		s,
	}
}

// Insert a new expression into database
func (controller *APIControllers) RegisterExpression(ctx *fiber.Ctx) error {
	receivedExpression := new(entities.Expression)
	if err := ctx.BodyParser(receivedExpression); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error on parsing in RegisterExpression": err.Error(),
		})
	}

	createdExpressionId, err := controller.s.RegisterExpression(receivedExpression)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on register a new expression": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdExpressionId)
}

// Retrieve all expressions from database
func (controller *APIControllers) GetExpressions(ctx *fiber.Ctx) error {
	expressions, err := controller.s.GetExpressions()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(200).JSON(expressions)
}

// Update an existent Expression
func (controller *APIControllers) UpdateExpression(ctx *fiber.Ctx) error {
	id := ctx.Params("expressionID")
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error on parse expressionID": err.Error(),
		})
	}

	expressionReceived := new(entities.Expression)
	if err := ctx.BodyParser(expressionReceived); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error on parsing in UpdateExpression": err.Error(),
		})
	}

	expressionString := &expressionReceived.ExpressionString

	updatedExpression, err := controller.s.UpdateExpression(parsedId, expressionString)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error on update expression": err.Error(),
		})
	}

	return ctx.Status(200).JSON(updatedExpression)
}

// Delete an existent Expression
func (controller *APIControllers) DeleteExpression(ctx *fiber.Ctx) error {
	id := ctx.Params("expressionID")
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on parse expressionID": err.Error(),
		})
	}

	if err := controller.s.DeleteExpressionById(parsedId); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on delete expression by ID": err.Error(),
		})
	}

	return ctx.Status(204).JSON("")
}

// Calculate an expression based on received values
func (controller *APIControllers) CalculateExpression(ctx *fiber.Ctx) error {
	id := ctx.Params("expressionID")
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on parse expressionID": err.Error(),
		})
	}

	result, err := controller.s.CalculateExpression(parsedId, ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on calculate": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

// Authenticate an user with a token to reach others endpoints
func (controller *APIControllers) Authenticate(ctx *fiber.Ctx) error {
	userToken, err := auth.GenerateToken()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on generate token": err.Error(),
		})
	}

	return ctx.Status(200).JSON(userToken)
}
