package model

import "time"

type Nurses struct {
	StrNum         string    `json:"str_num" validate:"required" gorm:"unique"`
	UserId         string    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name" validate:"required"`
	Competency     string    `json:"competency" validate:"required"`
	POB            string    `json:"pob" validate:"required"`
	DOB            string    `json:"dob" validate:"required"`
	Gender         string    `json:"gender" validate:"required"`
	Married        bool      `json:"married"`
	PhoneNum       string    `json:"phone_num"`
	Email          string    `json:"email"`
	Address        string    `json:"address" validate:"required"`
	District       string    `json:"district" validate:"required"`
	City           string    `json:"city" validate:"required"`
	Province       string    `json:"province" validate:"required"`
	GraduationYear int       `json:"graduation_year" validate:"required"`
	ExpYear        int       `json:"exp_year" validate:"required"`
	LastEducation  string    `json:"last_education" validate:"required"`
	UrlImage       string    `json:"url_image"`
}
