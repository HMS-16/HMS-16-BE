package dto

import (
	"HMS-16-BE/model"
	"strings"
	"time"
)

type PatientCards struct {
	Date   string `json:"date"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Status bool   `json:"status"`
}

type Patients struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	POB           string `json:"pob"`
	DOB           string `json:"DOB"`
	Gender        string `json:"gender"`
	Married       bool   `json:"married"`
	BloodType     string `json:"blood_type"`
	PhoneNum      string `json:"phone_num"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Province      string `json:"province"`
	FamilyName    string `json:"family_name"`
	Relationship  string `json:"relationship"`
	FamilyContact string `json:"family_contact"`
	Status        bool   `json:"status"`
}

func PatientCardsDTO(p *model.Patients) *PatientCards {
	return &PatientCards{
		Date:   p.CreatedAt.String(),
		Id:     p.Id,
		Name:   p.Name,
		Age:    GetAge(p),
		Gender: p.Gender,
		Status: p.Status,
	}
}

func PatientDTO(p *model.Patients) *Patients {
	return &Patients{
		Id:            p.Id,
		Name:          p.Name,
		POB:           p.POB,
		DOB:           p.DOB,
		Gender:        p.Gender,
		Married:       p.Married,
		PhoneNum:      p.PhoneNum,
		BloodType:     p.BloodType,
		Email:         p.Email,
		Address:       strings.Join([]string{p.Address, p.District}, ", "),
		City:          p.City,
		Province:      p.Province,
		FamilyName:    p.NameFamily,
		Relationship:  p.RelationshipFamily,
		FamilyContact: p.PhoneNumFamily,
		Status:        p.Status,
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
