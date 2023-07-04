package outpatientSession

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"HMS-16-BE/repository/outpatientSession"
	"HMS-16-BE/repository/patient"
	"HMS-16-BE/repository/shift"
	"HMS-16-BE/repository/user"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type ScheduleUsecase interface {
	Create(schedule model.Schedules) error
	GetAll() ([]dto.SchedulePatientInUser, error)
	GetAllCardByDay(date string) ([]dto.SchedulePatientCards, error) //website
	GetAllByDay(date string) ([]dto.SchedulePatientInUser, error)    //mobile
	//GetAllByDoctor(doctorId string) ([]model.Schedules, error)
	GetByScheduleId(id uint) (model.Schedules, error)
	GetAllByPatient(patientId string) ([]dto.SchedulePatientInUser, error)
	GetDetailByPatient(patientId string) (dto.PatientDetail, error)
	Update(schedule model.Schedules) error
	UpdateDoctor(schedule model.Schedules) error
	UpdateNurse(schedule model.Schedules) error
	UpdateStatus(id uint) error
	Delete(id uint) error
}

type scheduleUsecase struct {
	schedule  outpatientSession.ScheduleRepository
	condition outpatientSession.ConditionRepository
	diagnose  outpatientSession.DiagnoseRepository
	patient   patient.PatientRepository
	user      user.UserRepository
	shift     shift.ShiftRepository
}

func NewScheduleUsecase(
	s outpatientSession.ScheduleRepository,
	c outpatientSession.ConditionRepository,
	d outpatientSession.DiagnoseRepository,
	p patient.PatientRepository,
	u user.UserRepository,
	h shift.ShiftRepository) *scheduleUsecase {
	return &scheduleUsecase{s, c, d, p, u, h}
}

func (s *scheduleUsecase) Create(schedule model.Schedules) error {
	date, err := time.Parse("01/02/2006", schedule.Date)
	if err != nil {
		return err
	}
	day := date.Weekday().String()
	dayId := dto.DayShift(day)
	usersId, err := s.shift.GetIdByShift(dayId, schedule.TimeId)
	fmt.Println(day, dayId, usersId)
	if err != nil {
		return err
	}
	for _, userId := range usersId {
		user, err := s.user.GetById(userId)
		if err != nil {
			return err
		}
		if user.Role == 1 {
			schedule.DoctorId = user.StrNum
		} else if user.Role == 2 {
			schedule.NurseId = user.StrNum
		}
		fmt.Println(user, schedule.DoctorId, schedule.NurseId)
	}
	validate := validator.New()
	err = validate.Struct(&schedule)
	if err != nil {
		return err
	}
	return s.schedule.Create(schedule)

}

func (s *scheduleUsecase) GetAll() ([]dto.SchedulePatientInUser, error) {
	schedules, err := s.schedule.GetAll()
	if err != nil {
		return nil, err
	}
	var schedulesDTO []dto.SchedulePatientInUser
	for _, schedule := range schedules {
		patient, err := s.patient.GetById(schedule.PatientId)
		doctor, err := s.user.GetById(schedule.DoctorId)
		nurse, err := s.user.GetById(schedule.NurseId)
		if err != nil {
			return nil, err
		}
		schedulesDTO = append(schedulesDTO, *dto.SchedulePatientInUserDTO(&schedule, &patient, &doctor, &nurse))
	}
	return schedulesDTO, err
}

func (s *scheduleUsecase) GetAllCardByDay(date string) ([]dto.SchedulePatientCards, error) {
	schedules, err := s.schedule.GetAllByDay(date)
	if err != nil {
		return nil, err
	}
	var schedulesDTO []dto.SchedulePatientCards
	for _, schedule := range schedules {
		patient, err := s.patient.GetById(schedule.PatientId)
		doctor, err := s.user.GetById(schedule.DoctorId)
		if err != nil {
			return nil, err
		}
		schedulesDTO = append(schedulesDTO, *dto.SchedulePatientCardsDTO(&schedule, &patient, &doctor))
	}
	return schedulesDTO, err
}

func (s *scheduleUsecase) GetAllByDay(date string) ([]dto.SchedulePatientInUser, error) {
	schedules, err := s.schedule.GetAllByDay(date)
	if err != nil {
		return nil, err
	}
	var schedulesDTO []dto.SchedulePatientInUser
	for _, schedule := range schedules {
		patient, err := s.patient.GetById(schedule.PatientId)
		doctor, err := s.user.GetById(schedule.DoctorId)
		nurse, err := s.user.GetById(schedule.NurseId)
		if err != nil {
			return nil, err
		}
		schedulesDTO = append(schedulesDTO, *dto.SchedulePatientInUserDTO(&schedule, &patient, &doctor, &nurse))
	}
	return schedulesDTO, err
}

func (s *scheduleUsecase) GetByScheduleId(id uint) (model.Schedules, error) {
	schedule, err := s.schedule.GetById(id)
	if err != nil {
		return model.Schedules{}, err
	}
	return schedule, nil
}

func (s *scheduleUsecase) GetAllByPatient(patientId string) ([]dto.SchedulePatientInUser, error) {
	schedules, err := s.schedule.GetAllByPatient(patientId)
	if err != nil {
		return nil, err
	}
	var schedulesDTO []dto.SchedulePatientInUser
	for _, schedule := range schedules {
		patient, err := s.patient.GetById(schedule.PatientId)
		doctor, err := s.user.GetById(schedule.DoctorId)
		nurse, err := s.user.GetById(schedule.NurseId)
		if err != nil {
			return nil, err
		}
		schedulesDTO = append(schedulesDTO, *dto.SchedulePatientInUserDTO(&schedule, &patient, &doctor, &nurse))
	}
	return schedulesDTO, err
}

func (s *scheduleUsecase) GetDetailByPatient(patientId string) (dto.PatientDetail, error) {
	schedules, err := s.schedule.GetAllByPatient(patientId)
	if err != nil {
		return dto.PatientDetail{}, err
	}
	var schedulePatients []dto.SchedulePatient
	for _, schedule := range schedules {
		var schedulePatient dto.SchedulePatient
		schedulePatient.Schedule = schedule
		nurse, err := s.user.GetById(schedule.NurseId)
		if err != nil {
			return dto.PatientDetail{}, err
		}
		condition, err := s.condition.GetBySchedule(schedule.ID)
		if err != nil {
			return dto.PatientDetail{}, err
		}
		schedulePatient.Condition = *dto.ConditionDTO(&condition, &schedule, &nurse)
		doctor, err := s.user.GetById(schedule.DoctorId)
		if err != nil {
			return dto.PatientDetail{}, err
		}
		diagnose, err := s.diagnose.GetBySchedule(schedule.ID)
		if err != nil {
			return dto.PatientDetail{}, err
		}
		schedulePatient.Diagnose = *dto.DiagnoseDTO(&diagnose, &schedule, &doctor)
		schedulePatients = append(schedulePatients, schedulePatient)
	}
	patient, err := s.patient.GetById(patientId)
	if err != nil {
		return dto.PatientDetail{}, err
	}
	patientDetail := *dto.PatientDetailDTO(&patient, &schedulePatients)
	return patientDetail, nil
}

func (s *scheduleUsecase) Update(schedule model.Schedules) error {
	return s.schedule.Update(schedule)
}

func (s *scheduleUsecase) UpdateDoctor(schedule model.Schedules) error {
	return s.schedule.UpdateDoctor(schedule)
}

func (s *scheduleUsecase) UpdateNurse(schedule model.Schedules) error {
	return s.schedule.UpdateNurse(schedule)
}

func (s *scheduleUsecase) UpdateStatus(id uint) error {
	return s.schedule.UpdateStatus(id)
}

func (s *scheduleUsecase) Delete(id uint) error {
	return s.schedule.Delete(id)
}
