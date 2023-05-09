package mocks

import (
	"backend-challenge-api/internal/domain/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

type ExpressionServicesMock struct {
	mock.Mock
}

func NewExpressionServiceMock() *ExpressionServicesMock {
	return &ExpressionServicesMock{}
}

func (s *ExpressionServicesMock) RegisterExpression(expression *entities.Expression) (entities.Expression, error) {
	args := s.Called(expression)

	return args.Get(0).(entities.Expression), args.Error(1)
}

func (s *ExpressionServicesMock) GetExpressions() ([]entities.Expression, error) {
	args := s.Called()

	return args.Get(0).([]entities.Expression), args.Error(1)
}

func (s *ExpressionServicesMock) UpdateExpression(parsedId uint64, expressionString *string) (entities.Expression, error) {
	args := s.Called(parsedId, expressionString)

	return args.Get(0).(entities.Expression), args.Error(1)
}

func (s *ExpressionServicesMock) DeleteExpressionById(id uint64) error {
	args := s.Called(id)

	return args.Error(0)
}

func (s *ExpressionServicesMock) CalculateExpression(parsedId uint64, ctx *fiber.Ctx) (interface{}, error) {
	args := s.Called(parsedId, ctx)

	return args.Get(0).(interface{}), args.Error(1)
}

func (s *ExpressionServicesMock) GetExpressionById(id uint64) (entities.Expression, error) {
	args := s.Called(id)

	return args.Get(0).(entities.Expression), args.Error(1)
}
