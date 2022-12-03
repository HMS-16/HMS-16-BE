package dto

import (
	"time"
)

type Admin struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	PhoneNum  string    `json:"phone_num"`
}
