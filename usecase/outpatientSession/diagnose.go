package outpatientSession

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/outpatientSession"
	"HMS-16-BE/repository/user"
)

type DiagnoseUsecase interface {
	Create(diagnose model.Diagnoses) error
	GetById(id uint) (dto.Diagnose, error)
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

func (d *diagnoseUsecase) Create(diagnose model.Diagnoses) error {
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
