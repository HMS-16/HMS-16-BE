package dto

import (
	"HMS-16-BE/model"
	"time"
)

type User struct {
	StrNum    string    `json:"str_num"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      int       `json:"role"`
}

func UserDTO(user *model.Users) *User {
	return &User{
		StrNum:    user.STRNum,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
	}
}
