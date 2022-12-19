package model

import "gorm.io/gorm"

type Diagnoses struct {
	gorm.Model
	DoctorId     string `json:"doctor_id" validate:"required"`
	Diagnose     string `json:"diagnose" validate:"required"`
	Prescription string `json:"prescription" validate:"required"`
	ScheduleId   uint   `json:"schedule_id" validate:"required"`
	Status       bool   `json:"status" gorm:"default:false"`
}
