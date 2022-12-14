package profile

import (
	"HMS-16-BE/model"
	"database/sql"
)

type NurseRepository interface {
	GetAll() ([]model.Nurses, error)
	GetById(id string) (model.Nurses, error)
	Create(nurse model.Nurses) error
	Update(nurse model.Nurses) error
	Delete(id string) error
}

type nurseRepository struct {
	db *sql.DB
}

func NewNurseRepository(db *sql.DB) *nurseRepository {
	return &nurseRepository{db}
}

func (d *nurseRepository) GetAll() ([]model.Nurses, error) {
	query := `SELECT * FROM nurses`
	row, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var nurses []model.Nurses
	for row.Next() {
		var nurse model.Nurses
		err = row.Scan(&nurse.StrNum, &nurse.UserId, &nurse.CreatedAt, &nurse.UpdatedAt, &nurse.Name,
			&nurse.Competency, &nurse.POB, &nurse.DOB, &nurse.Gender, &nurse.Married, &nurse.PhoneNum,
			&nurse.Email, &nurse.Address, &nurse.District, &nurse.City, &nurse.Province, &nurse.GraduationYear,
			&nurse.ExpYear, &nurse.LastEducation, &nurse.UrlImage)
		if err != nil {
			return nil, err
		}
		nurses = append(nurses, nurse)
	}

	return nurses, nil
}

func (d *nurseRepository) GetById(id string) (model.Nurses, error) {
	query := `SELECT * FROM nurses WHERE str_num = ?`
	row, err := d.db.Query(query, id)
	if err != nil {
		return model.Nurses{}, err
	}
	defer row.Close()
	var nurse model.Nurses
	for row.Next() {
		err = row.Scan(&nurse.StrNum, &nurse.UserId, &nurse.CreatedAt, &nurse.UpdatedAt, &nurse.Name,
			&nurse.Competency, &nurse.POB, &nurse.DOB, &nurse.Gender, &nurse.Married, &nurse.PhoneNum,
			&nurse.Email, &nurse.Address, &nurse.District, &nurse.City, &nurse.Province, &nurse.GraduationYear,
			&nurse.ExpYear, &nurse.LastEducation, &nurse.UrlImage)
		if err != nil {
			return model.Nurses{}, err
		}
	}

	return nurse, nil
}

func (d *nurseRepository) Create(nurse model.Nurses) error {
	query := `INSERT INTO nurses VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := d.db.Exec(query, nurse.StrNum, nurse.UserId, nurse.CreatedAt, nurse.UpdatedAt, nurse.Name,
		nurse.Competency, nurse.POB, nurse.DOB, nurse.Gender, nurse.Married, nurse.PhoneNum,
		nurse.Email, nurse.Address, nurse.District, nurse.City, nurse.Province, nurse.GraduationYear,
		nurse.ExpYear, nurse.LastEducation, nurse.UrlImage)
	if err != nil {
		return err
	}

	return nil
}

func (d *nurseRepository) Update(nurse model.Nurses) error {
	query := `UPDATE nurses SET updated_at = ?, name = ?, pob = ?, dob = ?, gender = ?, married = ?, 
                   competency = ?, phone_num = ?, email = ?, address = ?, district = ?, city = ?, province = ?, graduation_year = ?, 
                   exp_year = ?, last_education = ?, url_image = ? WHERE str_num = ?`
	_, err := d.db.Exec(query, nurse.UpdatedAt, nurse.Name, nurse.POB, nurse.DOB,
		nurse.Gender, nurse.Married, nurse.Competency, nurse.PhoneNum, nurse.Email, nurse.Address,
		nurse.District, nurse.City, nurse.Province, nurse.GraduationYear,
		nurse.ExpYear, nurse.LastEducation, nurse.UrlImage, nurse.StrNum)
	if err != nil {
		return err
	}

	return nil
}

func (d *nurseRepository) Delete(id string) error {
	query := `DELETE FROM nurses WHERE str_num = ?`
	_, err := d.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
