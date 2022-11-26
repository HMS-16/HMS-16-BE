package model

import (
	"HMS-16-BE/dto"
	"time"
)

type Users struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	PhoneNum  string    `json:"phone_num"`
	Role      int       `json:"role"`
}

func (u *Users) ToDTO() *dto.User {
	return &dto.User{
		u.Id,
		u.CreatedAt,
		u.UpdatedAt,
		u.Username,
		u.Email,
		u.PhoneNum,
		u.Role,
	}
}
