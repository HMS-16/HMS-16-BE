package shift

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/shift"
)

type DayUsecase interface {
	GetAll() ([]dto.Day, error)
	GetById(id uint) (dto.Day, error)
	Create(day model.Days) error
	Update(day model.Days) error
	Delete(id uint) error
}

type dayUsecase struct {
	day shift.DayRepository
}

func NewDayUsecase(t shift.DayRepository) *dayUsecase {
	return &dayUsecase{t}
}

func (t *dayUsecase) GetAll() ([]dto.Day, error) {
	days, err := t.day.GetAll()
	var daysDTO []dto.Day
	for _, day := range days {
		dayDTO := dto.DayDTO(&day)
		daysDTO = append(daysDTO, *dayDTO)
	}

	return daysDTO, err
}

func (t *dayUsecase) GetById(id uint) (dto.Day, error) {
	day, err := t.day.GetById(id)
	return *dto.DayDTO(&day), err
}

func (t *dayUsecase) Create(days model.Days) error {
	return t.day.Create(days)
}

func (t *dayUsecase) Update(days model.Days) error {
	return t.day.Update(days)
}

func (t *dayUsecase) Delete(id uint) error {
	return t.day.Delete(id)
}
