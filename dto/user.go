package dto

import "time"

type User struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	PhoneNum  string    `json:"phone_num"`
	Role      int       `json:"role"`
}
