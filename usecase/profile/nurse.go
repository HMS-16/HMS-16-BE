package profile

import (
	"HMS-16-BE/model"
	"HMS-16-BE/repository/profile"
)

type NurseUsecase interface {
	GetAll() ([]model.Nurses, error)
	GetById(id string) (model.Nurses, error)
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

func (n *nurseUsecase) GetAll() ([]model.Nurses, error) {
	nurses, err := n.nurse.GetAll()
	return nurses, err
}

func (n *nurseUsecase) GetById(id string) (model.Nurses, error) {
	nurse, err := n.nurse.GetById(id)
	return nurse, err
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
