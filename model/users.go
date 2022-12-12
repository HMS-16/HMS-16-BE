package model

import (
	_ "github.com/go-playground/validator/v10"
	"time"
)

type Users struct {
	Id        string    `json:"id" gorm:"primary key" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username" validate:"required,min=3,max=32" gorm:"unique"`
	Password  string    `json:"password" validate:"required,min=6"`
	Email     string    `json:"email" validate:"required,email" gorm:"unique"`
	Role      int       `json:"role" validate:"required,gte=0,lte=2"`
	STRNum    string    `json:"str_num" validate:"required"`
}
