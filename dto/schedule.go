package dto

type SchedulePatientCards struct {
	Date            string `json:"date"`
	Shift           string `json:"shift"`
	NoMedicalRecord string `json:"no_medical_record"`
	Name            string `json:"name"`
	Doctor          string `json:"doctor"`
	Status          string `json:"status"`
}
