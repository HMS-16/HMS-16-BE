package shift

import (
	"HMS-16-BE/model"
	"database/sql"
)

type TimeRepository interface {
	GetAll() ([]model.Times, error)
	GetById(id uint) (model.Times, error)
	Create(time model.Times) error
	Update(time model.Times) error
	Delete(id uint) error
}

type timeRepository struct {
	db *sql.DB
}

func NewTimeRepository(db *sql.DB) *timeRepository {
	return &timeRepository{db}
}

func (t *timeRepository) GetAll() ([]model.Times, error) {
	query := `SELECT * FROM times ORDER BY id ASC`
	row, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	var times []model.Times
	defer row.Close()
	for row.Next() {
		var time model.Times
		err = row.Scan(&time.ID, &time.CreatedAt, &time.UpdatedAt, &time.DeletedAt, &time.Start, &time.End)
		if err != nil {
			return nil, err
		}
		times = append(times, time)
	}

	return times, nil
}

func (t *timeRepository) GetById(id uint) (model.Times, error) {
	query := `SELECT * FROM times WHERE id = ?`
	row, err := t.db.Query(query, id)
	if err != nil {
		return model.Times{}, err
	}
	defer row.Close()
	var time model.Times
	for row.Next() {
		err = row.Scan(&time.ID, &time.CreatedAt, &time.UpdatedAt, &time.DeletedAt, &time.Start, &time.End)
		if err != nil {
			return model.Times{}, err
		}
	}

	return time, nil
}

func (t *timeRepository) Create(time model.Times) error {
	query := `INSERT INTO times VALUES (?,?,?,?,?,?)`
	_, err := t.db.Exec(query, time.ID, time.CreatedAt, time.UpdatedAt, time.DeletedAt, time.Start, time.End)
	if err != nil {
		return err
	}

	return nil
}

func (t *timeRepository) Update(time model.Times) error {
	query := `UPDATE times SET start = ?, end = ? WHERE id = ?`
	_, err := t.db.Exec(query, time.Start, time.End, time.ID)
	if err != nil {
		return err
	}

	return nil
}

func (t *timeRepository) Delete(id uint) error {
	query := `DELETE FROM times WHERE id = ?`
	_, err := t.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
