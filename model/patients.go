package model

import (
	"time"
)

type Patients struct {
	Id                 string    `json:"id" validate:"required"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Name               string    `json:"name" validate:"required" gorm:"unique"`
	POB                string    `json:"pob" validate:"required"`
	DOB                string    `json:"dob" validate:"required"`
	Gender             string    `json:"gender" validate:"required"`
	Married            bool      `json:"married"`
	BloodType          string    `json:"blood_type" validate:"required"`
	PhoneNum           string    `json:"phone_num" validate:"required"`
	Email              string    `json:"email" validate:"required"`
	Address            string    `json:"address" validate:"required"`
	District           string    `json:"district" validate:"required"`
	City               string    `json:"city" validate:"required"`
	Province           string    `json:"province" validate:"required"`
	NameFamily         string    `json:"name_family" validate:"required"`
	RelationshipFamily string    `json:"relationship_family" validate:"required"`
	PhoneNumFamily     string    `json:"phone_num_family" validate:"required"`
	EmailFamily        string    `json:"email_family" validate:"required"`
	AddressFamily      string    `json:"address_family" validate:"required"`
	DistrictFamily     string    `json:"district_family" validate:"required"`
	CityFamily         string    `json:"city_family" validate:"required"`
	ProvinceFamily     string    `json:"province_family" validate:"required"`
	Status             bool      `json:"status" gorm:"default:false"`
	AdminId            string    `json:"admin_id"`
}
