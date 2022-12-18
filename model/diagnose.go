package model

import "gorm.io/gorm"

type Diagnoses struct {
	gorm.Model
	DoctorId     string `json:"doctor_id"`
	Diagnose     string `json:"diagnose"`
	Prescription string `json:"prescription"`
	ScheduleId   uint   `json:"schedule_id"`
	Status       bool   `json:"status" gorm:"default:false"`
}
