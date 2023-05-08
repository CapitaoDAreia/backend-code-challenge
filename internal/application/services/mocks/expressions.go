package mocks

import (
	"backend-challenge-api/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type ExpressionServicesMock struct {
	mock.Mock
}

func NewExpressionServiceMock() *ExpressionServicesMock {
	return &ExpressionServicesMock{}
}

func (s *ExpressionServicesMock) RegisterExpression(expression entities.Expression) (*entities.Expression, error) {
	args := s.Called(expression)

	return args.Get(0).(*entities.Expression), args.Error(1)
}
