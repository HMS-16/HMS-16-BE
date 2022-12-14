package profile

import (
	"HMS-16-BE/model"
	"database/sql"
)

type DoctorRepository interface {
	GetAll() ([]model.Doctors, error)
	GetById(id string) (model.Doctors, error)
	Create(doctor model.Doctors) error
	Update(doctor model.Doctors) error
	Delete(id string) error
}

type doctorRepository struct {
	db *sql.DB
}

func NewDoctorRepository(db *sql.DB) *doctorRepository {
	return &doctorRepository{db}
}

func (d *doctorRepository) GetAll() ([]model.Doctors, error) {
	query := `SELECT * FROM doctors`
	row, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var doctors []model.Doctors
	for row.Next() {
		var doctor model.Doctors
		err = row.Scan(&doctor.StrNum, &doctor.UserId, &doctor.CreatedAt, &doctor.UpdatedAt, &doctor.Name,
			&doctor.Competency, &doctor.POB, &doctor.DOB, &doctor.Gender, &doctor.Married, &doctor.PhoneNum,
			&doctor.Email, &doctor.Address, &doctor.District, &doctor.City, &doctor.Province, &doctor.GraduationYear,
			&doctor.ExpYear, &doctor.LastEducation, &doctor.UrlImage)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (d *doctorRepository) GetById(id string) (model.Doctors, error) {
	query := `SELECT * FROM doctors WHERE str_num = ?`
	row, err := d.db.Query(query, id)
	if err != nil {
		return model.Doctors{}, err
	}
	defer row.Close()
	var doctor model.Doctors
	for row.Next() {
		err = row.Scan(&doctor.StrNum, &doctor.UserId, &doctor.CreatedAt, &doctor.UpdatedAt, &doctor.Name,
			&doctor.Competency, &doctor.POB, &doctor.DOB, &doctor.Gender, &doctor.Married, &doctor.PhoneNum,
			&doctor.Email, &doctor.Address, &doctor.District, &doctor.City, &doctor.Province, &doctor.GraduationYear,
			&doctor.ExpYear, &doctor.LastEducation, &doctor.UrlImage)
		if err != nil {
			return model.Doctors{}, err
		}
	}

	return doctor, nil
}

func (d *doctorRepository) Create(doctor model.Doctors) error {
	query := `INSERT INTO doctors VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := d.db.Exec(query, doctor.StrNum, doctor.UserId, doctor.CreatedAt, doctor.UpdatedAt, doctor.Name,
		doctor.Competency, doctor.POB, doctor.DOB, doctor.Gender, doctor.Married, doctor.PhoneNum, doctor.Email,
		doctor.Address, doctor.District, doctor.City, doctor.Province, doctor.GraduationYear,
		doctor.ExpYear, doctor.LastEducation, doctor.UrlImage)
	if err != nil {
		return err
	}

	return nil
}

func (d *doctorRepository) Update(doctor model.Doctors) error {
	query := `UPDATE doctors SET updated_at = ?, name = ?, pob = ?, dob = ?, gender = ?, married = ?, 
                   competency = ?, phone_num = ?, email = ?, address = ?, district = ?, city = ?, province = ?, graduation_year = ?, 
                   exp_year = ?, last_education = ?, url_image = ? WHERE str_num = ?`
	_, err := d.db.Exec(query, doctor.UpdatedAt, doctor.Name, doctor.POB, doctor.DOB, doctor.Gender, doctor.Married,
		doctor.Competency, doctor.PhoneNum, doctor.Email, doctor.Address, doctor.District, doctor.City, doctor.Province,
		doctor.GraduationYear, doctor.ExpYear, doctor.LastEducation, doctor.UrlImage, doctor.StrNum)
	if err != nil {
		return err
	}

	return nil
}

func (d *doctorRepository) Delete(id string) error {
	query := `DELETE FROM doctors WHERE str_num = ?`
	_, err := d.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
