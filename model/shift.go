package model

import (
	"gorm.io/gorm"
)

type Days struct {
	gorm.Model
	Day string `json:"day" validate:"required"`
}

type Times struct {
	gorm.Model
	Start string `json:"start" validate:"required"`
	End   string `json:"end" validate:"required"`
}

type Shifts struct {
	gorm.Model
	UserId string `json:"user_id" validate:"required"`
	DayId  uint   `json:"day_id" validate:"required"`
	TimeId uint   `json:"time_id" validate:"required"`
}
