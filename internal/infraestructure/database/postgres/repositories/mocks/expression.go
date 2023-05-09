package mocks

import (
	"backend-challenge-api/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type ExpressionsRepositoryMock struct {
	mock.Mock
}

func NewExpressionsRepositoryMock() *ExpressionsRepositoryMock {
	return &ExpressionsRepositoryMock{}
}

func (r *ExpressionsRepositoryMock) RegisterExpression(expression *entities.Expression) (*entities.Expression, error) {
	args := r.Called(expression)

	return args.Get(0).(*entities.Expression), args.Error(1)
}

func (r *ExpressionsRepositoryMock) GetExpressions() ([]entities.Expression, error) {
	args := r.Called()

	return args.Get(0).([]entities.Expression), args.Error(1)
}

func (r *ExpressionsRepositoryMock) GetExpressionById(id uint64) (entities.Expression, error) {
	args := r.Called(id)

	return args.Get(0).(entities.Expression), args.Error(1)
}

func (r *ExpressionsRepositoryMock) UpdateExpression(expression entities.Expression, expressionString string) (entities.Expression, error) {
	args := r.Called(expression, expressionString)

	return args.Get(0).(entities.Expression), args.Error(1)
}

func (r *ExpressionsRepositoryMock) DeleteExpressionById(id uint64, expression entities.Expression) error {
	args := r.Called(id, expression)

	return args.Error(0)
}
