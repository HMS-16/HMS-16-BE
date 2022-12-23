package outpatientSession

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/outpatientSession"
	"HMS-16-BE/repository/user"
	"github.com/go-playground/validator/v10"
)

type DiagnoseUsecase interface {
	Create(diagnose model.Diagnoses, patientId string) error
	GetById(id uint) (dto.Diagnose, error)
	GetAllByPatient(patientId string) ([]dto.Diagnose, error)
}

type diagnoseUsecase struct {
	diagnose outpatientSession.DiagnoseRepository
	schedule outpatientSession.ScheduleRepository
	user     user.UserRepository
}

func NewDiagnoseUseCase(d outpatientSession.DiagnoseRepository, s outpatientSession.ScheduleRepository,
	u user.UserRepository) *diagnoseUsecase {
	return &diagnoseUsecase{d, s, u}
}

func (d *diagnoseUsecase) Create(diagnose model.Diagnoses, patientId string) error {
	date := diagnose.CreatedAt.Format("01/02/2006")

	var err error
	diagnose.ScheduleId, err = d.schedule.GetIdByPatient(patientId, date)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(&diagnose)
	if err != nil {
		return err
	}
	return d.diagnose.Create(diagnose)
}

func (d *diagnoseUsecase) GetById(id uint) (dto.Diagnose, error) {
	diagnose, err := d.diagnose.GetById(id)
	if err != nil {
		return dto.Diagnose{}, err
	}
	schedule, err := d.schedule.GetById(diagnose.ScheduleId)
	if err != nil {
		return dto.Diagnose{}, err
	}
	user, err := d.user.GetById(diagnose.DoctorId)
	if err != nil {
		return dto.Diagnose{}, err
	}
	return *dto.DiagnoseDTO(&diagnose, &schedule, &user), nil
}

func (d *diagnoseUsecase) GetAllByPatient(patientId string) ([]dto.Diagnose, error) {
	diagnoses, err := d.diagnose.GetAllByPatient(patientId)
	if err != nil {
		return nil, err
	}
	var diagnosesDTO []dto.Diagnose
	for _, diagnose := range diagnoses {
		schedule, err := d.schedule.GetById(diagnose.ScheduleId)
		if err != nil {
			return nil, err
		}
		user, err := d.user.GetById(diagnose.DoctorId)
		if err != nil {
			return nil, err
		}
		diagnosesDTO = append(diagnosesDTO, *dto.DiagnoseDTO(&diagnose, &schedule, &user))
	}
	return diagnosesDTO, nil
}
