package dto

import "time"

type Guardians struct {
	Id           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name""`
	Relationship string    `json:"relationship"`
	PhoneNum     string    `json:"phone_num"`
	Email        string    `json:"email""`
	Address      string    `json:"address"`
	District     string    `json:"district"`
	City         string    `json:"city"`
	Province     string    `json:"province"`
	PatientId    string    `json:"patient_id"`
}
