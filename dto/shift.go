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

func TimeShift(id uint) string {
	arr := []string{
		"no data",
		"09.00 - 12.00",
		"13.00 - 16.00",
		"18.00 - 21.00",
	}
	return arr[id]
}

func DayShift(name string) uint {
	arr := map[string]uint{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}
	return arr[name]
}
