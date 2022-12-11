package shift

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/shift"
)

type ShiftUsecase interface {
	GetAll() ([]dto.ShiftUsers, error)
	GetAllByUserId(userId string) (dto.ShiftUsers, error)
	GetById(id string) (model.Shifts, error)
	Create(shift model.Shifts) error
	Update(shift model.Shifts) error
	Delete(id string) error
}

type shiftUsecase struct {
	shift shift.ShiftRepository
	day   shift.DayRepository
	time  shift.TimeRepository
}

func NewShiftUsecase(s shift.ShiftRepository, d shift.DayRepository, t shift.TimeRepository) *shiftUsecase {
	return &shiftUsecase{s, d, t}
}

func (s *shiftUsecase) GetAllByUserId(userId string) (dto.ShiftUsers, error) {
	shifts, err := s.shift.GetAllByUserId(userId)
	if err != nil {
		return dto.ShiftUsers{}, err
	}
	var sessions []dto.Sessions
	for _, shift := range shifts {
		day, err := s.day.GetById(shift.DayId)
		if err != nil {
			return dto.ShiftUsers{}, err
		}

		time, err := s.time.GetById(shift.TimeId)
		if err != nil {
			return dto.ShiftUsers{}, err
		}

		session := dto.ShiftSessionDTO(shift.ID, &day, &time)
		sessions = append(sessions, *session)
	}
	shiftUser := dto.ShiftDTO(userId, sessions)
	return *shiftUser, nil
}

func (s *shiftUsecase) GetAll() ([]dto.ShiftUsers, error) {
	users, err := s.shift.GetAllUserId()
	if err != nil {
		return nil, err
	}

	var shiftUsers []dto.ShiftUsers
	for _, userId := range users {
		shifts, err := s.shift.GetAllByUserId(userId)
		if err != nil {
			return nil, err
		}
		var sessions []dto.Sessions
		for _, shift := range shifts {
			day, err := s.day.GetById(shift.DayId)
			if err != nil {
				return nil, err
			}

			time, err := s.time.GetById(shift.TimeId)
			if err != nil {
				return nil, err
			}

			session := dto.ShiftSessionDTO(shift.ID, &day, &time)
			sessions = append(sessions, *session)
		}
		shiftUser := dto.ShiftDTO(userId, sessions)
		shiftUsers = append(shiftUsers, *shiftUser)
	}

	return shiftUsers, nil
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
