package model

import (
	"HMS-16-BE/dto"
	"time"
)

type Patients struct {
	Id        string    `json:"id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" validate:"required"`
	POB       string    `json:"pob" validate:"required"`
	DOB       string    `json:"dob" validate:"required"`
	Gender    string    `json:"gender" validate:"required"`
	Married   bool      `json:"married"`
	PhoneNum  string    `json:"phone_num"`
	Email     string    `json:"email"`
	Address   string    `json:"address" validate:"required"`
	District  string    `json:"district" validate:"required"`
	City      string    `json:"city" validate:"required"`
	Province  string    `json:"province" validate:"required"`
	AdminId   string    `json:"admin_id" validate:"required"`
}

func (p *Patients) ToDTO() *dto.Patients {
	return &dto.Patients{
		Id:        p.Id,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Name:      p.Name,
		POB:       p.POB,
		DOB:       p.DOB,
		Gender:    p.Gender,
		Married:   p.Married,
		PhoneNum:  p.PhoneNum,
		Email:     p.Email,
		AdminId:   p.AdminId,
		Age:       p.GetAge(),
	}
}

func (p *Patients) GetAge() int {
	today := time.Now()
	birthdate, _ := time.Parse("01/02/2006", p.DOB) // format: MM/DD/YYY
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}
