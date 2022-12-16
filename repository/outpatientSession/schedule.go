package outpatientSession

import (
	"HMS-16-BE/model"
	"database/sql"
)

type ScheduleRepository interface {
	Create(schedule model.Schedules) error
	GetAllByDay(date string) ([]model.Schedules, error)
	GetAllByPatient(patientId string) ([]model.Schedules, error)
	GetAllByDoctor(doctorId string) ([]model.Schedules, error)
	GetAllById(id string) ([]model.Schedules, error)
	Update(schedule model.Schedules) error
	UpdateStatus(id string) error
	Delete(id string) error
}

type scheduleRepository struct {
	db *sql.DB
}

func NewScheduleRepository(db *sql.DB) *scheduleRepository {
	return &scheduleRepository{db}
}

func (s *scheduleRepository) Create(schedule model.Schedules) error {
	query := `INSERT INTO schedules VALUES (?,?,?,?,?,?,?,?)`
	_, err := s.db.Exec(query, schedule.ID, schedule.CreatedAt, schedule.UpdatedAt, schedule.DeletedAt,
		schedule.PatientId, schedule.DoctorId, schedule.Date, schedule.TimeId)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduleRepository) GetAllByDay(date string) ([]model.Schedules, error) {
	query := `SELECT * FROM schedules WHERE date = ?`
	row, err := s.db.Query(query, date)
	if err != nil {
		return nil, err
	}
	var schedules []model.Schedules
	defer row.Close()
	for row.Next() {
		var schedule model.Schedules
		err = row.Scan(&schedule.ID, &schedule.CreatedAt, &schedule.UpdatedAt, &schedule.DeletedAt,
			&schedule.PatientId, &schedule.DoctorId, &schedule.Date, &schedule.TimeId, &schedule.Status)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (s *scheduleRepository) GetAllByPatient(patientId string) ([]model.Schedules, error) {
	query := `SELECT * FROM schedules WHERE patient_id = ?`
	row, err := s.db.Query(query, patientId)
	if err != nil {
		return nil, err
	}
	var schedules []model.Schedules
	defer row.Close()
	for row.Next() {
		var schedule model.Schedules
		err = row.Scan(&schedule.ID, &schedule.CreatedAt, &schedule.UpdatedAt, &schedule.DeletedAt,
			&schedule.PatientId, &schedule.DoctorId, &schedule.Date, &schedule.TimeId, &schedule.Status)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (s *scheduleRepository) GetAllByDoctor(doctorId string) ([]model.Schedules, error) {
	query := `SELECT * FROM schedules WHERE doctor_id = ?`
	row, err := s.db.Query(query, doctorId)
	if err != nil {
		return nil, err
	}
	var schedules []model.Schedules
	defer row.Close()
	for row.Next() {
		var schedule model.Schedules
		err = row.Scan(&schedule.ID, &schedule.CreatedAt, &schedule.UpdatedAt, &schedule.DeletedAt,
			&schedule.PatientId, &schedule.DoctorId, &schedule.Date, &schedule.TimeId, &schedule.Status)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (s *scheduleRepository) GetAllById(id string) (model.Schedules, error) {
	query := `SELECT * FROM schedules WHERE id = ?`
	row, err := s.db.Query(query, id)
	if err != nil {
		return model.Schedules{}, err
	}
	var schedule model.Schedules
	defer row.Close()
	for row.Next() {
		err = row.Scan(&schedule.ID, &schedule.CreatedAt, &schedule.UpdatedAt, &schedule.DeletedAt,
			&schedule.PatientId, &schedule.DoctorId, &schedule.Date, &schedule.TimeId, &schedule.Status)
		if err != nil {
			return model.Schedules{}, err
		}
	}
	return schedule, nil
}

func (s *scheduleRepository) Update(schedule model.Schedules) error {
	query := `UPDATE schedules SET date = ?, time_id = ? WHERE id = ?`
	_, err := s.db.Exec(query, schedule.Date, schedule.TimeId, schedule.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduleRepository) UpdateStatus(id string) error {
	query := `UPDATE schedules SET status = true WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduleRepository) Delete(id string) error {
	query := `DELETE FROM schedules WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
