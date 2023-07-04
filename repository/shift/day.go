package shift

import (
	"HMS-16-BE/model"
	"database/sql"
)

type DayRepository interface {
	GetAll() ([]model.Days, error)
	GetById(id uint) (model.Days, error)
	Create(Day model.Days) error
	Update(Day model.Days) error
	Delete(id uint) error
}

type dayRepository struct {
	db *sql.DB
}

func NewDayRepository(db *sql.DB) *dayRepository {
	return &dayRepository{db}
}

func (d *dayRepository) GetAll() ([]model.Days, error) {
	query := `SELECT * FROM days ORDER BY id ASC`
	row, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	var Days []model.Days
	defer row.Close()
	for row.Next() {
		var Day model.Days
		err = row.Scan(&Day.ID, &Day.CreatedAt, &Day.UpdatedAt, &Day.DeletedAt, &Day.Day)
		if err != nil {
			return nil, err
		}
		Days = append(Days, Day)
	}

	return Days, nil
}

func (d *dayRepository) GetById(id uint) (model.Days, error) {
	query := `SELECT * FROM days WHERE id = ?`
	row, err := d.db.Query(query, id)
	if err != nil {
		return model.Days{}, err
	}
	defer row.Close()
	var Day model.Days
	for row.Next() {
		err = row.Scan(&Day.ID, &Day.CreatedAt, &Day.UpdatedAt, &Day.DeletedAt, &Day.Day)
		if err != nil {
			return model.Days{}, err
		}
	}

	return Day, nil
}

func (d *dayRepository) Create(Day model.Days) error {
	query := `INSERT INTO days VALUES (?,?,?,?,?)`
	_, err := d.db.Exec(query, Day.ID, Day.CreatedAt, Day.UpdatedAt, Day.DeletedAt, Day.Day)
	if err != nil {
		return err
	}

	return nil
}

func (d *dayRepository) Update(Day model.Days) error {
	query := `UPDATE days SET updated_at = ?, day = ? WHERE id = ?`
	_, err := d.db.Exec(query, Day.UpdatedAt, Day.Day, Day.ID)
	if err != nil {
		return err
	}

	return nil
}

func (d *dayRepository) Delete(id uint) error {
	query := `DELETE FROM days WHERE id = ?`
	_, err := d.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
