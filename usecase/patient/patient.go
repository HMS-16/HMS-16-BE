package patient

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/patient"
)

type PatientUsecase interface {
	GetAll() ([]dto.Patients, error)
	GetAllCards() ([]dto.PatientCards, error)
	GetById(id string) (dto.Patients, error)
	Create(patient model.Patients) error
	Update(patient model.Patients) error
	UpdateEndCase(id string) error
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
		guardian, _ := p.guardian.GetByPatientId(patient.Id)
		dto := dto.PatientDTO(&patient, &guardian)
		patientsDTO = append(patientsDTO, *dto)
	}
	return patientsDTO, err
}

func (p *patientUsecase) GetAllCards() ([]dto.PatientCards, error) {
	var patientsDTO []dto.PatientCards
	patients, err := p.patient.GetAll()
	for _, patient := range patients {
		dto := dto.PatientCardsDTO(&patient)
		patientsDTO = append(patientsDTO, *dto)
	}
	return patientsDTO, err
}

func (p *patientUsecase) GetById(id string) (dto.Patients, error) {
	patient, err := p.patient.GetById(id)
	guardian, err := p.guardian.GetByPatientId(id)
	dto := dto.PatientDTO(&patient, &guardian)
	return *dto, err
}

func (p *patientUsecase) Create(patient model.Patients) error {
	err := p.patient.Create(patient)
	return err
}

func (p *patientUsecase) Update(patient model.Patients) error {
	err := p.patient.Update(patient)
	return err
}

func (p *patientUsecase) UpdateEndCase(id string) error {
	err := p.patient.UpdateEndCase(id)
	return err
}

func (p *patientUsecase) Delete(id string) error {
	err := p.patient.Delete(id)
	return err
}
