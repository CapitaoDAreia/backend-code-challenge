package repositories

import (
	"backend-challenge-api/internal/domain/entities"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type ExpressionsRepository struct {
	DB *gorm.DB
}

func NewExpressionsRepository(DB *gorm.DB) *ExpressionsRepository {
	return &ExpressionsRepository{
		DB,
	}
}

func (r *ExpressionsRepository) RegisterExpression(expression *entities.Expression) (uint64, error) {
	err := r.DB.Create(expression)
	if err.Error != nil {
		return 0, err.Error
	}

	return expression.ID, nil
}
