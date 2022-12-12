package model

import (
	_ "github.com/go-playground/validator/v10"
	"time"
)

type Users struct {
	Id        string    `json:"id" gorm:"primary key" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username" validate:"required,min=3,max=32"`
	Password  string    `json:"password" validate:"required,min=6"`
	Email     string    `json:"email" validate:"required,email"`
	PhoneNum  string    `json:"phone_num" validate:"required,min=10"`
	Role      int       `json:"role" validate:"required,gte=0,lte=2"`
}
