package outpatientSession

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/outpatientSession"
	"HMS-16-BE/repository/user"
)

type ConditionUsecase interface {
	Create(condition model.Conditions) error
	GetById(id uint) (dto.Condition, error)
}

type conditionUsecase struct {
	condition outpatientSession.ConditionRepository
	schedule  outpatientSession.ScheduleRepository
	user      user.UserRepository
}

func NewConditionUsecase(c outpatientSession.ConditionRepository, s outpatientSession.ScheduleRepository,
	u user.UserRepository) *conditionUsecase {
	return &conditionUsecase{c, s, u}
}

func (c *conditionUsecase) Create(condition model.Conditions) error {
	return c.condition.Create(condition)
}

func (c *conditionUsecase) GetById(id uint) (dto.Condition, error) {
	condition, err := c.condition.GetById(id)
	if err != nil {
		return dto.Condition{}, err
	}
	schedule, err := c.schedule.GetById(condition.ScheduleId)
	if err != nil {
		return dto.Condition{}, err
	}
	user, err := c.user.GetById(condition.NurseId)
	if err != nil {
		return dto.Condition{}, err
	}
	return *dto.ConditionDTO(&condition, &schedule, &user), nil
}
