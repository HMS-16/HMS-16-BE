package patient

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/patient"
)

type PatientUsecase interface {
	GetAll() ([]dto.Patients, error)
	GetById(id string) (dto.Patients, error)
	Create(patient model.Patients) error
	Update(patient model.Patients) error
	Delete(id string) error
}

type patientUsecase struct {
	patient patient.PatientRepository
}

func NewPatientUsecase(p patient.PatientRepository) *patientUsecase {
	return &patientUsecase{p}
}

func (p *patientUsecase) GetAll() ([]dto.Patients, error) {
	var patientsDTO []dto.Patients
	patients, err := p.patient.GetAll()
	for _, patient := range patients {
		patientsDTO = append(patientsDTO, *patient.ToDTO())
	}
	return patientsDTO, err
}

func (p *patientUsecase) GetById(id string) (dto.Patients, error) {
	patient, err := p.patient.GetById(id)
	return *patient.ToDTO(), err
}

func (p *patientUsecase) Create(patient model.Patients) error {
	err := p.patient.Create(patient)
	return err
}

func (p *patientUsecase) Update(patient model.Patients) error {
	err := p.patient.Update(patient)
	return err
}

func (p *patientUsecase) Delete(id string) error {
	err := p.patient.Delete(id)
	return err
}
