package dto

import (
	"HMS-16-BE/model"
	"time"
)

type Admin struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	PhoneNum  string    `json:"phone_num"`
}

func AdminDTO(a *model.Admins) *Admin {
	return &Admin{
		ID:        a.ID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Username:  a.Username,
		PhoneNum:  a.PhoneNum,
	}
}
