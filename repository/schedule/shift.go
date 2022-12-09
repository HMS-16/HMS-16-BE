package schedule

import (
	"HMS-16-BE/model"
	"database/sql"
)

type ShiftRepository interface {
	GetAllByUserId(id string) ([]model.Shifts, error)
	GetById(id string) (model.Shifts, error)
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

func (s *shiftRepository) GetAllByUserId(id string) ([]model.Shifts, error) {
	query := `SELECT * FROM shifts WHERE user_id = ?`
	row, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var shifts []model.Shifts
	defer row.Close()
	for row.Next() {
		var shift model.Shifts
		err = row.Scan(&shift.ID, &shift.CreatedAt, &shift.UpdatedAt, &shift.DeletedAt, &shift.Date,
			&shift.Start, &shift.End, &shift.UserID)
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
		err = row.Scan(&shift.ID, &shift.CreatedAt, &shift.UpdatedAt, &shift.DeletedAt, &shift.Date,
			&shift.Start, &shift.End, &shift.UserID)
		if err != nil {
			return model.Shifts{}, err
		}
	}

	return shift, nil
}

func (s *shiftRepository) Create(shift model.Shifts) error {
	query := `INSERT INTO shifts VALUES (?,?,?,?,?,?,?,?,?)`
	_, err := s.db.Exec(query, shift.ID, shift.CreatedAt, shift.UpdatedAt, shift.DeletedAt, shift.Date,
		shift.Start, shift.End, shift.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (s *shiftRepository) Update(shift model.Shifts) error {
	query := `UPDATE shifts SET start = ?, end = ? WHERE id = ?`
	_, err := s.db.Exec(query, shift.Start, shift.End, shift.ID)
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
