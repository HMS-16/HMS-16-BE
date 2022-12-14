package profile

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/profile"
)

type DoctorUsecase interface {
	GetAll() ([]dto.Doctor, error)
	GetAllCards() ([]dto.DoctorCards, error)
	GetById(id string) (dto.Doctor, error)
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

func (d *doctorUsecase) GetAll() ([]dto.Doctor, error) {
	doctors, err := d.doctor.GetAll()
	var doctorsDTO []dto.Doctor
	for _, doctor := range doctors {
		doctorsDTO = append(doctorsDTO, *dto.DoctorDTO(&doctor))
	}
	return doctorsDTO, err
}

func (d *doctorUsecase) GetAllCards() ([]dto.DoctorCards, error) {
	doctors, err := d.doctor.GetAll()
	var doctorsDTO []dto.DoctorCards
	for _, doctor := range doctors {
		doctorsDTO = append(doctorsDTO, *dto.DoctorCardDTO(&doctor))
	}
	return doctorsDTO, err
}

func (d *doctorUsecase) GetById(id string) (dto.Doctor, error) {
	doctor, err := d.doctor.GetById(id)
	return *dto.DoctorDTO(&doctor), err
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
