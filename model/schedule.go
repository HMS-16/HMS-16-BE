package model

import "gorm.io/gorm"

type Schedules struct {
	gorm.Model
	PatientId string `json:"patient_id" validate:"required"`
	DoctorId  string `json:"doctor_id" validate:"required"`
	NurseId   string `json:"nurse_id" validate:"required"`
	Date      string `json:"date" validate:"required"`
	TimeId    uint   `json:"time_id" validate:"required"`
	Status    bool   `json:"status" gorm:"default:false"`
}
