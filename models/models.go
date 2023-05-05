package models

import (
	"gorm.io/gorm"
)

type Expression struct {
	gorm.Model
	ID               uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	ExpressionString string `json:"expression" gorm:"text;not null;default:null"`
}
