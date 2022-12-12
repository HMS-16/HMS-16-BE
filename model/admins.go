package model

import (
	"time"
)

type Admins struct {
	ID        string    `json:"id" gorm:"primary key" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username" validate:"required,min=3,max=32"`
	Password  string    `json:"password" validate:"required,min=6"`
	PhoneNum  string    `json:"phone_num" validate:"required,min=10"`
}
