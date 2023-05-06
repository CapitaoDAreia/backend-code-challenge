package controllers

import (
	"backend-challenge-api/internal/domain/entities"
	"backend-challenge-api/internal/infraestructure/database/postgres/repositories"
	"backend-challenge-api/internal/infraestructure/http/auth"
	"fmt"
	"strconv"

	"github.com/Knetic/govaluate"
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

// Insert a new expression into database
func (controller *APIControllers) RegisterExpression(ctx *fiber.Ctx) error {
	receivedExpression := new(entities.Expression)
	if err := ctx.BodyParser(receivedExpression); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error on parsing in RegisterExpression": err.Error(),
		})
	}

	createdExpressionId, err := controller.r.RegisterExpression(receivedExpression)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on register": err.Error(),
		})
	}

	return ctx.Status(200).JSON(createdExpressionId)
}

// Retrieve all expressions from database
func (controller *APIControllers) GetExpressions(ctx *fiber.Ctx) error {
	expressions, err := controller.r.GetExpressions()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(200).JSON(expressions)
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

	retrievedExpression, err := controller.r.GetExpressionById(parsedId)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on catch expressionID": err.Error(),
		})
	}

	evaluableExpression, err := govaluate.NewEvaluableExpression(retrievedExpression.ExpressionString)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on calculate expression": err.Error(),
		})
	}

	params := make(map[string]interface{})
	for _, v := range evaluableExpression.Vars() {
		val := ctx.Query(v)
		if val != "" {
			parsedValue, err := strconv.ParseBool(val)
			if err != nil {
				return ctx.Status(500).JSON(fiber.Map{
					"error on parse variable value": err.Error(),
				})
			}
			params[v] = parsedValue
		} else {
			return ctx.Status(400).JSON(fiber.Map{
				"error on query variable": fmt.Sprintf("Variable %s not found", v),
			})
		}
	}

	result, err := evaluableExpression.Evaluate(params)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on calculate expression": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

// Update an existent Expression
func (controller *APIControllers) UpdateExpression(ctx *fiber.Ctx) error {
	id := ctx.Params("expressionID")
	parsedId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on parse expressionID": err.Error(),
		})
	}

	expressionReceived := new(entities.Expression)

	if err := ctx.BodyParser(expressionReceived); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error on parsing in UpdateExpression": err.Error(),
		})
	}

	expression, err := controller.r.GetExpressionById(parsedId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error on found expression": err.Error(),
		})
	}

	updatedExpression, err := controller.r.UpdateExpression(expression, *&expressionReceived.ExpressionString)
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

	retrievedExpression, err := controller.r.GetExpressionById(parsedId)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on catch expressionID": err.Error(),
		})
	}

	retrievedExpressionID := uint64(retrievedExpression.ID)

	if err := controller.r.DeleteExpressionById(retrievedExpressionID, retrievedExpression); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on delete expression by ID": err.Error(),
		})
	}

	return ctx.Status(204).JSON("")
}

// Authenticate user
func (controller *APIControllers) Authenticate(ctx *fiber.Ctx) error {
	userToken, err := auth.GenerateToken()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error on generate token": err.Error(),
		})
	}

	return ctx.Status(200).JSON(userToken)
}
