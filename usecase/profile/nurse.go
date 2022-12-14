package profile

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/profile"
)

type NurseUsecase interface {
	GetAll() ([]dto.Nurse, error)
	GetAllCards() ([]dto.NurseCards, error)
	GetById(id string) (dto.Nurse, error)
	Create(nurse model.Nurses) error
	Update(nurse model.Nurses) error
	Delete(id string) error
}

type nurseUsecase struct {
	nurse profile.NurseRepository
}

func NewNurseUsecase(n profile.NurseRepository) *nurseUsecase {
	return &nurseUsecase{n}
}

func (n *nurseUsecase) GetAll() ([]dto.Nurse, error) {
	nurses, err := n.nurse.GetAll()
	var nursesDTO []dto.Nurse
	for _, nurse := range nurses {
		nursesDTO = append(nursesDTO, *dto.NurseDTO(&nurse))
	}
	return nursesDTO, err
}

func (n *nurseUsecase) GetAllCards() ([]dto.NurseCards, error) {
	nurses, err := n.nurse.GetAll()
	var nursesDTO []dto.NurseCards
	for _, nurse := range nurses {
		nursesDTO = append(nursesDTO, *dto.NurseCardsDTO(&nurse))
	}
	return nursesDTO, err
}

func (n *nurseUsecase) GetById(id string) (dto.Nurse, error) {
	nurse, err := n.nurse.GetById(id)
	return *dto.NurseDTO(&nurse), err
}

func (n *nurseUsecase) Create(nurse model.Nurses) error {
	err := n.nurse.Create(nurse)
	return err
}

func (n *nurseUsecase) Update(nurse model.Nurses) error {
	err := n.nurse.Update(nurse)
	return err
}

func (n *nurseUsecase) Delete(id string) error {
	err := n.nurse.Delete(id)
	return err
}
