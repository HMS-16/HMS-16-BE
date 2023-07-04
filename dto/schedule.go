package dto

import "HMS-16-BE/model"

type SchedulePatient struct {
	Schedule  model.Schedules `json:"schedule"`
	Condition Condition       `json:"condition"`
	Diagnose  Diagnose        `json:"diagnose"`
}

type SchedulePatientCards struct {
	Id              uint   `json:"id"`
	Date            string `json:"date"`
	Shift           string `json:"shift"`
	NoMedicalRecord string `json:"no_medical_record"`
	Name            string `json:"name"`
	Doctor          string `json:"doctor"`
	Status          bool   `json:"status"`
	StatusString    string `json:"status_string"`
}

type SchedulePatientInUser struct {
	Id           uint   `json:"id"`
	PatientId    string `json:"patient_id"`
	Date         string `json:"date"`
	Shift        string `json:"shift"`
	Name         string `json:"name"`
	Doctor       string `json:"doctor"`
	Nurse        string `json:"nurse"`
	Status       bool   `json:"status"`
	StatusString string `json:"status_string"`
}

type PatientDetail struct {
	PatientId    string            `json:"patient_id"`
	Patient      model.Patients    `json:"patient"`
	Medical      []SchedulePatient `json:"medical"`
	Status       bool              `json:"status"`
	StatusString string            `json:"status_string"`
}

func StatusSchedule(status bool) string {
	if status {
		return "Done"
	} else {
		return "Process"
	}
}

func SchedulePatientCardsDTO(s *model.Schedules, p *model.Patients, d *User) *SchedulePatientCards {
	return &SchedulePatientCards{
		Id:              s.ID,
		Date:            s.Date,
		Shift:           TimeShift(s.TimeId),
		NoMedicalRecord: p.Id,
		Name:            p.Name,
		Doctor:          d.Name,
		Status:          s.Status,
		StatusString:    StatusSchedule(s.Status),
	}
}

func SchedulePatientInUserDTO(s *model.Schedules, p *model.Patients, d, n *User) *SchedulePatientInUser {
	return &SchedulePatientInUser{
		Id:           s.ID,
		PatientId:    p.Id,
		Date:         s.Date,
		Shift:        TimeShift(s.TimeId),
		Name:         p.Name,
		Doctor:       d.Name,
		Nurse:        n.Name,
		Status:       s.Status,
		StatusString: StatusSchedule(s.Status),
	}
}

func PatientDetailDTO(p *model.Patients, sp *[]SchedulePatient) *PatientDetail {
	return &PatientDetail{
		PatientId:    p.Id,
		Patient:      *p,
		Medical:      *sp,
		Status:       p.Status,
		StatusString: StatusSchedule(p.Status),
	}
}
