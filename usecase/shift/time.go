package shift

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/shift"
)

type TimeUsecase interface {
	GetAll() ([]dto.Time, error)
	GetById(id uint) (dto.Time, error)
	Create(time model.Times) error
	Update(time model.Times) error
	Delete(id uint) error
}

type timeUsecase struct {
	time shift.TimeRepository
}

func NewtimeUsecase(t shift.TimeRepository) *timeUsecase {
	return &timeUsecase{t}
}

func (t *timeUsecase) GetAll() ([]dto.Time, error) {
	times, err := t.time.GetAll()
	var timesDTO []dto.Time
	for _, time := range times {
		timeDTO := dto.TimeDTO(&time)
		timesDTO = append(timesDTO, *timeDTO)
	}

	return timesDTO, err
}

func (t *timeUsecase) GetById(id uint) (dto.Time, error) {
	time, err := t.time.GetById(id)
	return *dto.TimeDTO(&time), err
}

func (t *timeUsecase) Create(times model.Times) error {
	return t.time.Create(times)
}

func (t *timeUsecase) Update(times model.Times) error {
	return t.time.Update(times)
}

func (t *timeUsecase) Delete(id uint) error {
	return t.time.Delete(id)
}
