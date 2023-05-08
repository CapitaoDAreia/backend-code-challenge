package services

import (
	expressionErrors "backend-challenge-api/internal/application/services/errors"
	"backend-challenge-api/internal/domain/entities"
	"backend-challenge-api/internal/infraestructure/database/postgres/repositories"
	"fmt"
	"log"
	"strconv"

	"github.com/Knetic/govaluate"
	"github.com/gofiber/fiber/v2"
)

type ExpressionServices interface {
	RegisterExpression(expression *entities.Expression) (entities.Expression, error)
	GetExpressions() ([]entities.Expression, error)
	GetExpressionById(id uint64) (entities.Expression, error)
	UpdateExpression(parsedId uint64, expressionString *string) (entities.Expression, error)
	DeleteExpressionById(id uint64) error
	CalculateExpression(parsedId uint64, ctx *fiber.Ctx) (interface{}, error)
}

type expressionService struct {
	expressionRepository repositories.ExpressionsRepository
}

func NewExpressionsServices(expressionRepository repositories.ExpressionsRepository) *expressionService {
	return &expressionService{
		expressionRepository,
	}
}

func (service *expressionService) RegisterExpression(expression *entities.Expression) (entities.Expression, error) {
	createdExpression, err := service.expressionRepository.RegisterExpression(expression)
	if err != nil {
		log.Println("error3")
		return entities.Expression{}, err
	}

	return createdExpression, nil
}

func (service *expressionService) GetExpressions() ([]entities.Expression, error) {
	expressions, err := service.expressionRepository.GetExpressions()
	if err != nil {
		return []entities.Expression{}, err
	}

	return expressions, nil
}

func (service *expressionService) GetExpressionById(id uint64) (entities.Expression, error) {
	returnedExpression, err := service.expressionRepository.GetExpressionById(id)
	if err != nil {
		return entities.Expression{}, err
	}

	return returnedExpression, nil
}

func (service *expressionService) UpdateExpression(parsedId uint64, expressionString *string) (entities.Expression, error) {
	existentExpression, err := service.GetExpressionById(parsedId)
	if err != nil {
		return entities.Expression{}, err
	}

	updatedExpression, err := service.expressionRepository.UpdateExpression(existentExpression, *expressionString)
	if err != nil {
		return entities.Expression{}, err
	}

	return updatedExpression, nil
}

func (service *expressionService) DeleteExpressionById(id uint64) error {

	existentExpression, err := service.GetExpressionById(id)
	if err != nil {
		return err
	}

	err = service.expressionRepository.DeleteExpressionById(id, existentExpression)
	if err != nil {
		return err
	}

	return nil
}

func (service *expressionService) CalculateExpression(parsedId uint64, ctx *fiber.Ctx) (interface{}, error) {

	retrievedExpression, err := service.expressionRepository.GetExpressionById(parsedId)
	if err != nil {
		return nil, err
	}

	evaluableExpression, err := govaluate.NewEvaluableExpression(retrievedExpression.ExpressionString)
	if err != nil {
		return nil, err
	}

	params := make(map[string]interface{})
	for _, v := range evaluableExpression.Vars() {
		val := ctx.Query(v)
		if val != "" {
			parsedValue, err := strconv.ParseBool(val)
			if err != nil {
				return nil, err
			}
			params[v] = parsedValue
		} else {
			message := fmt.Sprintf("missing variable: %s", v)
			expressionErrorResponse := expressionErrors.NewExpressionErrors(message)
			return nil, expressionErrorResponse
		}
	}

	result, err := evaluableExpression.Evaluate(params)

	if err != nil {
		return nil, err
	}

	return result, nil
}
