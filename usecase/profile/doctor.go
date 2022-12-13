package profile

import (
	"HMS-16-BE/model"
	"HMS-16-BE/repository/profile"
)

type DoctorUsecase interface {
	GetAll() ([]model.Doctors, error)
	GetById(id string) (model.Doctors, error)
	Create(doctor model.Doctors) error
	Update(doctor model.Doctors) error
	Delete(id string) error
}

type doctorUsecase struct {
	doctor profile.DoctorRepository
}

func NewDoctorUsecase(d profile.DoctorRepository) *doctorUsecase {
	return &doctorUsecase{d}
}

func (d *doctorUsecase) GetAll() ([]model.Doctors, error) {
	doctors, err := d.doctor.GetAll()
	return doctors, err
}

func (d *doctorUsecase) GetById(id string) (model.Doctors, error) {
	doctor, err := d.doctor.GetById(id)
	return doctor, err
}

func (d *doctorUsecase) Create(doctor model.Doctors) error {
	err := d.doctor.Create(doctor)
	return err
}

func (d *doctorUsecase) Update(doctor model.Doctors) error {
	err := d.doctor.Update(doctor)
	return err
}

func (d *doctorUsecase) Delete(id string) error {
	err := d.doctor.Delete(id)
	return err
}
