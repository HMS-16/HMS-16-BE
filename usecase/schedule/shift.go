package schedule

import (
	"HMS-16-BE/model"
	"HMS-16-BE/repository/schedule"
)

type ShiftUsecase interface {
	GetAllByUserId(id string) ([]model.Shifts, error)
	GetById(id string) (model.Shifts, error)
	Create(shift model.Shifts) error
	Update(shift model.Shifts) error
	Delete(id string) error
}

type shiftUsecase struct {
	shift schedule.ShiftRepository
}

func NewShiftUsecase(s schedule.ShiftRepository) *shiftUsecase {
	return &shiftUsecase{s}
}

func (s *shiftUsecase) GetAllByUserId(id string) ([]model.Shifts, error) {
	return s.shift.GetAllByUserId(id)
}

func (s *shiftUsecase) GetById(id string) (model.Shifts, error) {
	return s.shift.GetById(id)
}

func (s *shiftUsecase) Create(shifts model.Shifts) error {
	return s.shift.Create(shifts)
}

func (s *shiftUsecase) Update(shifts model.Shifts) error {
	return s.shift.Update(shifts)
}

func (s *shiftUsecase) Delete(id string) error {
	return s.shift.Delete(id)
}
