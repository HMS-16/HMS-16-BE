package outpatientSession

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/outpatientSession"
	"HMS-16-BE/repository/user"
	"github.com/go-playground/validator/v10"
)

type ConditionUsecase interface {
	Create(condition model.Conditions, patientId string) error
	GetById(id uint) (dto.Condition, error)
	GetAllByPatient(patientId string) ([]dto.Condition, error)
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

func (c *conditionUsecase) Create(condition model.Conditions, patientId string) error {
	date := condition.CreatedAt.Format("01/02/2006")

	var err error
	condition.ScheduleId, err = c.schedule.GetIdByPatient(patientId, date)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(&condition)
	if err != nil {
		return err
	}

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

func (c *conditionUsecase) GetAllByPatient(patientId string) ([]dto.Condition, error) {
	conditions, err := c.condition.GetAllByPatient(patientId)
	if err != nil {
		return nil, err
	}
	var conditionsDTO []dto.Condition
	for _, condition := range conditions {
		schedule, err := c.schedule.GetById(condition.ScheduleId)
		if err != nil {
			return nil, err
		}
		user, err := c.user.GetById(condition.NurseId)
		if err != nil {
			return nil, err
		}
		conditionsDTO = append(conditionsDTO, *dto.ConditionDTO(&condition, &schedule, &user))
	}
	return conditionsDTO, nil
}
