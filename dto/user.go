package dto

import (
	"HMS-16-BE/model"
	"time"
)

type User struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	PhoneNum  string    `json:"phone_num"`
	Role      int       `json:"role"`
}

func UserDTO(user *model.Users) *User {
	return &User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Email:     user.Email,
		PhoneNum:  user.PhoneNum,
		Role:      user.Role,
	}
}
