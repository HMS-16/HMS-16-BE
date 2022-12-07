package patient

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/patient"
)

type GuardianUsecase interface {
	GetById(id string) (dto.Guardians, error)
	Create(guardian model.Guardians) error
	Update(guardian model.Guardians) error
	Delete(id string) error
}

type guardianUsecase struct {
	guardian patient.GuardianRepositoy
}

func NewGuardianUSecase(g patient.GuardianRepositoy) *guardianUsecase {
	return &guardianUsecase{g}
}

func (g *guardianUsecase) GetById(id string) (dto.Guardians, error) {
	guardian, err := g.guardian.GetById(id)
	return guardian, err
}

func (g *guardianUsecase) Create(guardian model.Guardians) error {
	err := g.guardian.Create(guardian)
	return err
}

func (g *guardianUsecase) Update(guardian model.Guardians) error {
	err := g.guardian.Update(guardian)
	return err
}

func (g *guardianUsecase) Delete(id string) error {
	err := g.guardian.Delete(id)
	return err
}
