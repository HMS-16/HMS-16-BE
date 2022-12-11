package dto

import (
	"HMS-16-BE/model"
	"strings"
)

type Time struct {
	Id   uint   `json:"id"`
	Time string `json:"time"`
}

type Day struct {
	Id  uint   `json:"id"`
	Day string `json:"day"`
}

type ShiftUsers struct {
	UserId  string     `json:"user_id"`
	Session []Sessions `json:"session"`
}

type Sessions struct {
	Id     uint   `json:"id"`
	DayId  uint   `json:"day_id"`
	Day    string `json:"day"`
	TimeId uint   `json:"time_id"`
	Time   string `json:"time"`
}

func ShiftSessionDTO(id uint, day *model.Days, time *model.Times) *Sessions {
	return &Sessions{
		Id:     id,
		DayId:  day.ID,
		Day:    day.Day,
		TimeId: time.ID,
		Time:   strings.Join([]string{time.Start, time.End}, " - "),
	}
}

func ShiftDTO(id string, sessions []Sessions) *ShiftUsers {
	return &ShiftUsers{
		UserId:  id,
		Session: sessions,
	}
}

func TimeDTO(time *model.Times) *Time {
	return &Time{
		Id:   time.ID,
		Time: strings.Join([]string{time.Start, time.End}, " - "),
	}
}

func DayDTO(day *model.Days) *Day {
	return &Day{
		Id:  day.ID,
		Day: day.Day,
	}
}
