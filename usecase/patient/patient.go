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
	patient  patient.PatientRepository
	guardian patient.GuardianRepositoy
}

func NewPatientUsecase(p patient.PatientRepository, g patient.GuardianRepositoy) *patientUsecase {
	return &patientUsecase{p, g}
}

func (p *patientUsecase) GetAll() ([]dto.Patients, error) {
	var patientsDTO []dto.Patients
	patients, err := p.patient.GetAll()
	for _, patient := range patients {
		dto := *patient.ToDTO()
		dto.Guardians, err = p.guardian.GetAllByPatientId(patient.Id)
		patientsDTO = append(patientsDTO, dto)
	}
	return patientsDTO, err
}

func (p *patientUsecase) GetById(id string) (dto.Patients, error) {
	patient, err := p.patient.GetById(id)
	dto := *patient.ToDTO()
	dto.Guardians, err = p.guardian.GetAllByPatientId(patient.Id)
	return dto, err
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
