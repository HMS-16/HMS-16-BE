package shift

import (
	"HMS-16-BE/model"
	"database/sql"
)

type ShiftRepository interface {
	GetAllUserId() ([]string, error)
	GetAllByUserId(id string) ([]model.Shifts, error)
	GetById(id string) (model.Shifts, error)
	GetIdByShift(date, time uint) ([]string, error)
	Create(shift model.Shifts) error
	Update(shift model.Shifts) error
	Delete(id string) error
}

type shiftRepository struct {
	db *sql.DB
}

func NewShiftRepository(db *sql.DB) *shiftRepository {
	return &shiftRepository{db}
}

func (s *shiftRepository) GetAllUserId() ([]string, error) {
	query := `SELECT DISTINCT user_id FROM shifts ORDER BY day_id ASC`
	row, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var users []string
	defer row.Close()
	for row.Next() {
		var user string
		err = row.Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *shiftRepository) GetAllByUserId(id string) ([]model.Shifts, error) {
	query := `SELECT * FROM shifts WHERE user_id = ? ORDER BY day_id ASC`
	row, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var shifts []model.Shifts
	defer row.Close()
	for row.Next() {
		var shift model.Shifts
		err = row.Scan(&shift.ID, &shift.CreatedAt, &shift.UpdatedAt, &shift.DeletedAt, &shift.UserId,
			&shift.DayId, &shift.TimeId)
		if err != nil {
			return nil, err
		}
		shifts = append(shifts, shift)
	}

	return shifts, nil
}

func (s *shiftRepository) GetById(id string) (model.Shifts, error) {
	query := `SELECT * FROM shifts WHERE id = ?`
	row, err := s.db.Query(query, id)
	if err != nil {
		return model.Shifts{}, err
	}
	defer row.Close()
	var shift model.Shifts
	for row.Next() {
		err = row.Scan(&shift.ID, &shift.CreatedAt, &shift.UpdatedAt, &shift.DeletedAt, &shift.UserId,
			&shift.DayId, &shift.TimeId)
		if err != nil {
			return model.Shifts{}, err
		}
	}

	return shift, nil
}

func (s *shiftRepository) GetIdByShift(date, time uint) ([]string, error) {
	query := `SELECT user_id FROM shifts WHERE day_id = ?, time_id = ?`
	row, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var users []string
	defer row.Close()
	for row.Next() {
		var user string
		err = row.Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *shiftRepository) Create(shift model.Shifts) error {
	query := `INSERT INTO shifts VALUES (?,?,?,?,?,?,?)`
	_, err := s.db.Exec(query, shift.ID, shift.CreatedAt, shift.UpdatedAt, shift.DeletedAt, shift.UserId,
		shift.DayId, shift.TimeId)
	if err != nil {
		return err
	}

	return nil
}

func (s *shiftRepository) Update(shift model.Shifts) error {
	query := `UPDATE shifts SET user_id = ?, day_id = ?, time_id = ? WHERE id = ?`
	_, err := s.db.Exec(query, shift.UserId, shift.DayId, shift.TimeId, shift.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *shiftRepository) Delete(id string) error {
	query := `DELETE FROM shifts WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
