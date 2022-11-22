package usecase

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/admin"
)

type AdminUsecase interface {
	Create(admin model.Admins) error
	Login(username, password string) (dto.Admin, error)
	GetById(id string) (dto.Admin, error)
	Update(admin model.Admins) error
	Delete(id string) error
}

type adminUsecase struct {
	admin repository.AdminRepository
}

func NewAdminUsecase(a repository.AdminRepository) *adminUsecase {
	return &adminUsecase{a}
}

func (a *adminUsecase) Create(admin model.Admins) error {
	err := a.admin.Create(admin)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminUsecase) Login(username, password string) (dto.Admin, error) {
	admin, err := a.admin.Login(username, password)
	if err != nil {
		return dto.Admin{}, err
	}
	return *admin.ToDTO(), nil
}

func (a *adminUsecase) GetById(id string) (dto.Admin, error) {
	admin, err := a.admin.GetById(id)
	if err != nil {
		return dto.Admin{}, err
	}
	return *admin.ToDTO(), nil
}

func (a *adminUsecase) Update(admin model.Admins) error {
	err := a.admin.Update(admin)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminUsecase) Delete(id string) error {
	err := a.admin.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
