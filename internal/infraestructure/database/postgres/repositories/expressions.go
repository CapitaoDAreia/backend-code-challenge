package repositories

import (
	"backend-challenge-api/internal/domain/entities"
	"log"

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

func (r *ExpressionsRepository) RegisterExpression(expression *entities.Expression) (entities.Expression, error) {
	err := r.DB.Create(expression)
	if err.Error != nil {
		log.Println("error4.")
		return entities.Expression{}, err.Error
	}

	return *expression, nil
}

func (r *ExpressionsRepository) GetExpressions() ([]entities.Expression, error) {
	var expressions []entities.Expression
	if err := r.DB.Find(&expressions).Error; err != nil {
		return nil, err
	}

	return expressions, nil
}

func (r *ExpressionsRepository) GetExpressionById(id uint64) (entities.Expression, error) {
	var expression entities.Expression
	if err := r.DB.First(&expression, id).Error; err != nil {
		return entities.Expression{}, err
	}
	return expression, nil
}

func (r *ExpressionsRepository) UpdateExpression(expression entities.Expression, expressionString string) (entities.Expression, error) {
	expression.ExpressionString = expressionString

	if err := r.DB.Save(&expression).Error; err != nil {
		return entities.Expression{}, err
	}

	return expression, nil
}

func (r *ExpressionsRepository) DeleteExpressionById(id uint64, expression entities.Expression) error {
	if err := r.DB.Delete(&expression, id).Error; err != nil {
		return err
	}
	return nil
}
