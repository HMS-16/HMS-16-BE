package dto

import (
	"HMS-16-BE/model"
	"time"
)

type Diagnose struct {
	Id           uint   `json:"id"`
	RegisterDate string `json:"register_date"`
	DoctorId     string `json:"doctor_id"`
	Name         string `json:"name"`
	Diagnose     string `json:"diagnose"`
	Prescription string `json:"prescription"`
	ScheduleId   string `json:"schedule_id"`
	Time         string `json:"time"`
	Status       bool   `json:"status"`
}

func DiagnoseDTO(d *model.Diagnoses, s *model.Schedules, u *User) *Diagnose {
	return &Diagnose{
		Id:           d.ID,
		RegisterDate: d.CreatedAt.Format(time.RFC850),
		DoctorId:     d.DoctorId,
		Name:         u.Name,
		Diagnose:     d.Diagnose,
		Prescription: d.Prescription,
		ScheduleId:   d.ScheduleId,
		Time:         TimeShift(s.TimeId),
		Status:       d.Status,
	}
}
