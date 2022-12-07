package dto

import (
	"time"
)

type Patients struct {
	Id        string      `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Name      string      `json:"name"`
	POB       string      `json:"pob"`
	DOB       string      `json:"DOB"`
	Gender    string      `json:"gender"`
	Married   bool        `json:"married"`
	PhoneNum  string      `json:"phone_num"`
	Email     string      `json:"email"`
	Address   string      `json:"address"`
	District  string      `json:"district"`
	City      string      `json:"city"`
	Province  string      `json:"province"`
	Status    bool        `json:"status"`
	AdminId   string      `json:"admin_id"`
	Age       int         `json:"age"`
	Guardians []Guardians `json:"guardians"`
}
