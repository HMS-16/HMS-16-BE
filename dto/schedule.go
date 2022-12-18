package dto

import "HMS-16-BE/model"

type SchedulePatient struct {
	Schedule  model.Schedules `json:"schedule"`
	Condition Condition       `json:"condition"`
	Diagnose  Diagnose        `json:"diagnose"`
}

type SchedulePatientCards struct {
	PatientId       string `json:"patient_id"`
	Date            string `json:"date"`
	Shift           string `json:"shift"`
	NoMedicalRecord string `json:"no_medical_record"`
	Name            string `json:"name"`
	Doctor          string `json:"doctor"`
	Status          string `json:"status"`
}

type SchedulePatientInUser struct {
	PatientId string `json:"patient_id"`
	Date      string `json:"date"`
	Shift     string `json:"shift"`
	Name      string `json:"name"`
	Doctor    string `json:"doctor"`
	Nurse     string `json:"nurse"`
	Status    string `json:"status"`
}

type PatientDetail struct {
	PatientId string            `json:"patient_id"`
	Patient   model.Patients    `json:"patient"`
	Schedule  []SchedulePatient `json:"schedule"`
	Status    string            `json:"status"`
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
		PatientId:       p.Id,
		Date:            s.Date,
		Shift:           TimeShift(s.TimeId),
		NoMedicalRecord: p.Id,
		Name:            p.Name,
		Doctor:          d.Name,
		Status:          StatusSchedule(s.Status),
	}
}

func SchedulePatientInUserDTO(s *model.Schedules, p *model.Patients, d, n *User) *SchedulePatientInUser {
	return &SchedulePatientInUser{
		PatientId: p.Id,
		Date:      s.Date,
		Shift:     TimeShift(s.TimeId),
		Name:      p.Name,
		Doctor:    d.Name,
		Nurse:     n.Name,
		Status:    StatusSchedule(s.Status),
	}
}

func PatientDetailDTO(p *model.Patients, sp *[]SchedulePatient) *PatientDetail {
	return &PatientDetail{
		PatientId: p.Id,
		Patient:   *p,
		Schedule:  *sp,
		Status:    StatusSchedule(p.Status),
	}
}