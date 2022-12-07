package model

import "time"

type Guardians struct {
	Id           string    `json:"id" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name" validate:"required"`
	Relationship string    `json:"relationship" validate:"required"`
	PhoneNum     string    `json:"phone_num" validate:"required"`
	Email        string    `json:"email" validate:"required"`
	Address      string    `json:"address" validate:"required"`
	District     string    `json:"district" validate:"required"`
	City         string    `json:"city" validate:"required"`
	Province     string    `json:"province" validate:"required"`
	PatientId    string    `json:"patient_id"`
}
