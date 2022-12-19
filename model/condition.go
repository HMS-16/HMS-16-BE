package model

import "gorm.io/gorm"

type Conditions struct {
	gorm.Model
	NurseId         string `json:"nurse_id" validate:"required"`
	Height          string `json:"height" validate:"required"`
	Weight          string `json:"weight" validate:"required"`
	BloodPressure   string `json:"blood_pressure" validate:"required"`
	SugarAnalysis   string `json:"sugar_analysis" validate:"required"`
	BodyTemperature string `json:"body_temperature" validate:"required"`
	HeartRate       string `json:"heart_rate" validate:"required"`
	BreathRate      string `json:"breath_rate" validate:"required"`
	Cholesterol     string `json:"cholesterol" validate:"required"`
	Note            string `json:"note" validate:"required"`
	ScheduleId      uint   `json:"schedule_id" validate:"required"`
	Status          bool   `json:"status" gorm:"default:false"`
}
