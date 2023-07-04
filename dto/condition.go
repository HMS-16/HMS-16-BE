package dto

import (
	"HMS-16-BE/model"
	"time"
)

type Condition struct {
	Id              uint   `json:"id"`
	RegisterDate    string `json:"register_date"`
	NurseId         string `json:"nurse_id"`
	Name            string `json:"name"`
	Height          string `json:"height"`
	Weight          string `json:"weight"`
	BloodPressure   string `json:"blood_pressure"`
	SugarAnalysis   string `json:"sugar_analysis"`
	BodyTemperature string `json:"body_temperature"`
	HeartRate       string `json:"heart_rate"`
	BreathRate      string `json:"breath_rate"`
	Cholesterol     string `json:"cholesterol"`
	Note            string `json:"note"`
	ScheduleId      uint   `json:"schedule_id"`
	Time            string `json:"time"`
	Status          bool   `json:"status"`
}

func ConditionDTO(c *model.Conditions, s *model.Schedules, u *User) *Condition {
	return &Condition{
		Id:              c.ID,
		RegisterDate:    c.CreatedAt.Format(time.RFC850),
		NurseId:         c.NurseId,
		Name:            u.Name,
		Height:          c.Height,
		Weight:          c.Weight,
		BloodPressure:   c.BloodPressure,
		SugarAnalysis:   c.SugarAnalysis,
		BodyTemperature: c.BodyTemperature,
		HeartRate:       c.HeartRate,
		BreathRate:      c.BreathRate,
		Cholesterol:     c.Cholesterol,
		Note:            c.Note,
		ScheduleId:      c.ScheduleId,
		Time:            TimeShift(s.TimeId),
		Status:          c.Status,
	}
}
