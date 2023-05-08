package repositories

import "backend-challenge-api/internal/domain/entities"

type ExpressionsRepository interface {
	RegisterExpression(expression *entities.Expression) (*entities.Expression, error)
	GetExpressions() ([]entities.Expression, error)
	GetExpressionById(id uint64) (entities.Expression, error)
	UpdateExpression(expression entities.Expression, expressionString string)
	DeleteExpressionById(id uint64, expression entities.Expression)
}
