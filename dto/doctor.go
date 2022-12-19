package dto

import (
	"HMS-16-BE/model"
	"strconv"
	"strings"
)

type DoctorCards struct {
	StrNum   string `json:"str_num"`
	Name     string `json:"name"`
	UrlImage string `json:"url_image"`
}

type Doctor struct {
	StrNum         string `json:"str_num" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Competency     string `json:"competency" validate:"required"`
	POB            string `json:"pob" validate:"required"`
	DOB            string `json:"dob" validate:"required"`
	Gender         string `json:"gender" validate:"required"`
	Married        bool   `json:"married"`
	PhoneNum       string `json:"phone_num"`
	Email          string `json:"email"`
	LastEducation  string `json:"last_education" validate:"required"`
	GraduationYear int    `json:"graduation_year" validate:"required"`
	ExpYear        string `json:"exp_year" validate:"required"`
	Address        string `json:"address" validate:"required"`
	District       string `json:"district" validate:"required"`
	City           string `json:"city" validate:"required"`
	Province       string `json:"province" validate:"required"`
	UrlImage       string `json:"url_image"`
}

func DoctorDTO(d *model.Doctors) *Doctor {
	return &Doctor{
		StrNum:         d.StrNum,
		Name:           d.Name,
		POB:            d.POB,
		Competency:     d.Competency,
		DOB:            d.DOB,
		Gender:         d.Gender,
		Married:        d.Married,
		PhoneNum:       d.PhoneNum,
		Email:          d.Email,
		LastEducation:  d.LastEducation,
		GraduationYear: d.GraduationYear,
		ExpYear:        strings.Join([]string{strconv.Itoa(d.ExpYear), "year"}, " "),
		Address:        d.Address,
		District:       d.District,
		City:           d.City,
		Province:       d.Province,
		UrlImage:       d.UrlImage,
	}
}

func DoctorCardDTO(d *model.Doctors) *DoctorCards {
	return &DoctorCards{
		StrNum:   d.StrNum,
		Name:     d.Name,
		UrlImage: d.UrlImage,
	}
}
