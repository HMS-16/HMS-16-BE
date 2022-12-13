package user

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/user"
)

type UserUsecase interface {
	Create(user model.Users) error
	Login(email string) (model.Users, error)
	GetAll() ([]dto.User, error)
	GetById(id string) (dto.User, error)
	Update(user model.Users) error
	Delete(id string) error
}

type userUsecase struct {
	user user.UserRepository
}

func NewUserUsecase(u user.UserRepository) *userUsecase {
	return &userUsecase{u}
}

func (u *userUsecase) Create(user model.Users) error {
	return u.user.Create(user)
}

func (u *userUsecase) Login(email string) (model.Users, error) {
	return u.user.Login(email)
}

func (u *userUsecase) GetAll() ([]dto.User, error) {
	return u.user.GetAll()
}

func (u *userUsecase) GetById(id string) (dto.User, error) {
	return u.user.GetById(id)
}

func (u *userUsecase) Update(user model.Users) error {
	return u.user.Update(user)
}

func (u *userUsecase) Delete(id string) error {
	return u.user.Delete(id)
}
