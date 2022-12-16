package model

import "gorm.io/gorm"

type Conditions struct {
	gorm.Model
	NurseId         string `json:"nurse_id"`
	Height          string `json:"height"`
	Weight          string `json:"weight"`
	BloodPressure   string `json:"blood_pressure"`
	SugarAnalysis   string `json:"sugar_analysis"`
	BodyTemperature string `json:"body_temperature"`
	HeartRate       string `json:"heart_rate"`
	BreathRate      string `json:"breath_rate"`
	Cholesterol     string `json:"cholesterol"`
	Note            string `json:"note"`
	ScheduleId      uint   `json:"schedule_id"`
	Status          bool   `json:"status" gorm:"default:false"`
}
