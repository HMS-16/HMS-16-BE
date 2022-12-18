package dto

import (
	"HMS-16-BE/model"
	"time"
)

type User struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"username"`
	Email     string    `json:"email"`
	Role      int       `json:"role"`
}

func UserDTO(user *model.Users) *User {
	return &User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
	}
}
