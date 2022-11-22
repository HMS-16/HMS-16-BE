package model

import (
	"HMS-16-BE/dto"
	"time"
)

type Admins struct {
	ID        string    `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	PhoneNum  string    `json:"phone_num"`
}

func (a *Admins) ToDTO() *dto.Admin {
	return &dto.Admin{
		ID:        a.ID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Username:  a.Username,
		PhoneNum:  a.PhoneNum,
	}
}
