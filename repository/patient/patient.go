package patient

import (
	"HMS-16-BE/model"
	"database/sql"
)

type PatientRepository interface {
	GetAll() ([]model.Patients, error)
	GetById(id string) (model.Patients, error)
	Create(patient model.Patients) error
	Update(patient model.Patients) error
	UpdateEndCase(id string) error
	Delete(id string) error
}

type patientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) *patientRepository {
	return &patientRepository{db}
}

func (p *patientRepository) GetAll() ([]model.Patients, error) {
	query := `SELECT * FROM patients ORDER BY created_at ASC`
	row, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	var patients []model.Patients
	defer row.Close()
	for row.Next() {
		var patient model.Patients
		err = row.Scan(&patient.Id, &patient.CreatedAt, &patient.UpdatedAt, &patient.Name, &patient.POB,
			&patient.DOB, &patient.Gender, &patient.Married, &patient.BloodType, &patient.PhoneNum, &patient.Email,
			&patient.Address, &patient.District, &patient.City, &patient.Province, &patient.Status, &patient.AdminId)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

func (p *patientRepository) GetById(id string) (model.Patients, error) {
	query := `SELECT * FROM patients WHERE id = ?`
	row, err := p.db.Query(query, id)
	if err != nil {
		return model.Patients{}, err
	}
	defer row.Close()
	var patient model.Patients
	for row.Next() {
		err = row.Scan(&patient.Id, &patient.CreatedAt, &patient.UpdatedAt, &patient.Name, &patient.POB,
			&patient.DOB, &patient.Gender, &patient.Married, &patient.BloodType, &patient.PhoneNum, &patient.Email,
			&patient.Address, &patient.District, &patient.City, &patient.Province, &patient.Status, &patient.AdminId)
		if err != nil {
			return model.Patients{}, err
		}
	}
	return patient, nil
}

func (p *patientRepository) Create(patient model.Patients) error {
	query := `INSERT INTO patients VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := p.db.Exec(query, patient.Id, patient.CreatedAt, patient.UpdatedAt, patient.Name, patient.POB,
		patient.DOB, patient.Gender, patient.Married, patient.BloodType, patient.PhoneNum, patient.Email,
		patient.Address, patient.District, patient.City, patient.Province, patient.Status, patient.AdminId)
	if err != nil {
		return err
	}
	return nil
}

func (p *patientRepository) Update(patient model.Patients) error {
	query := `UPDATE patients SET updated_at = ?, name = ?, pob = ?, dob = ?, gender = ?, married = ?, blood_type = ?, 
            phone_num = ?, email = ?, address = ?, district = ?, city = ?, province = ? WHERE id = ?`
	_, err := p.db.Exec(query, patient.UpdatedAt, patient.Name, patient.POB, patient.DOB, patient.Gender, patient.Married,
		patient.BloodType, patient.PhoneNum, patient.Email, patient.Address, patient.District, patient.City,
		patient.Province, patient.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *patientRepository) UpdateEndCase(id string) error {
	query := "UPDATE patients SET status = TRUE WHERE id = ?"
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *patientRepository) Delete(id string) error {
	query := `DELETE FROM patients WHERE id = ?`
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
