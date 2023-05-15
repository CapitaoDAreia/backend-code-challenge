package entities

import (
	"gorm.io/gorm"
)

type Expression struct {
	gorm.Model
	ExpressionString string `json:"expression" gorm:"text;not null;default:null"`
}
