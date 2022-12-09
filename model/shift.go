package model

import (
	"gorm.io/gorm"
)

type Shifts struct {
	gorm.Model
	Date   string `json:"date" validate:"required"`
	Start  string `json:"start" validate:"required"`
	End    string `json:"end" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}
