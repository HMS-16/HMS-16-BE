package model

import (
	"time"
)

type Patients struct {
	Id        string    `json:"id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" validate:"required"`
	POB       string    `json:"pob" validate:"required"`
	DOB       string    `json:"dob" validate:"required"`
	Gender    string    `json:"gender" validate:"required"`
	Married   bool      `json:"married"`
	PhoneNum  string    `json:"phone_num"`
	Email     string    `json:"email"`
	Address   string    `json:"address" validate:"required"`
	District  string    `json:"district" validate:"required"`
	City      string    `json:"city" validate:"required"`
	Province  string    `json:"province" validate:"required"`
	Status    bool      `json:"status" default:"false"`
	AdminId   string    `json:"admin_id" validate:"required"`
}
