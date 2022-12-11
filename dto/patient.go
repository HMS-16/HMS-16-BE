package dto

import (
	"HMS-16-BE/model"
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

func PatientDTO(p *model.Patients) *Patients {
	return &Patients{
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
		Address:   p.Address,
		District:  p.District,
		City:      p.City,
		Province:  p.Province,
		Status:    p.Status,
		AdminId:   p.AdminId,
		Age:       GetAge(p),
	}
}

func GetAge(p *model.Patients) int {
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
