package model

import "time"

type Doctors struct {
	StrNum        string    `json:"strNum" validate:"required"`
	UserId        string    `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name" validate:"required"`
	POB           string    `json:"pob" validate:"required"`
	DOB           string    `json:"dob" validate:"required"`
	Gender        string    `json:"gender" validate:"required"`
	Married       bool      `json:"married"`
	PhoneNum      string    `json:"phone_num"`
	Email         string    `json:"email"`
	Address       string    `json:"address" validate:"required"`
	District      string    `json:"district" validate:"required"`
	City          string    `json:"city" validate:"required"`
	Province      string    `json:"province" validate:"required"`
	EntryYear     string    `json:"entry_year" validate:"required"`
	DoctorYear    string    `json:"doctor_year" validate:"required"`
	LastEducation string    `json:"last_education" validate:"required"`
	UrlImage      string    `json:"url_image"`
	Specialist    string    `json:"specialist" validate:"required"`
}
