package patient

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"database/sql"
)

type GuardianRepositoy interface {
	GetAllByPatientId(id string) ([]dto.Guardians, error)
	GetById(id string) (dto.Guardians, error)
	Create(guardian model.Guardians) error
	Update(guardian model.Guardians) error
	Delete(id string) error
}

type guardianRepository struct {
	db *sql.DB
}

func NewGuardianRepository(db *sql.DB) *guardianRepository {
	return &guardianRepository{db}
}

func (g *guardianRepository) GetAllByPatientId(id string) ([]dto.Guardians, error) {
	query := `SELECT * FROM guardians WHERE patient_id = ? ORDER BY name ASC`
	row, err := g.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	var guardians []dto.Guardians
	defer row.Close()
	for row.Next() {
		var guardian dto.Guardians
		err = row.Scan(&guardian.Id, &guardian.CreatedAt, &guardian.UpdatedAt, &guardian.Name,
			&guardian.Relationship, &guardian.PhoneNum, &guardian.Email, &guardian.Address, &guardian.District,
			&guardian.City, &guardian.Province, &guardian.PatientId)
		if err != nil {
			return nil, err
		}
		guardians = append(guardians, guardian)
	}
	return guardians, nil
}

func (g *guardianRepository) GetById(id string) (dto.Guardians, error) {
	query := `SELECT * FROM guardians WHERE id = ?`
	row, err := g.db.Query(query, id)
	if err != nil {
		return dto.Guardians{}, err
	}
	var guardian dto.Guardians
	defer row.Close()
	for row.Next() {
		err = row.Scan(&guardian.Id, &guardian.CreatedAt, &guardian.UpdatedAt, &guardian.Name,
			&guardian.Relationship, &guardian.PhoneNum, &guardian.Email, &guardian.Address, &guardian.District,
			&guardian.City, &guardian.Province, &guardian.PatientId)
		if err != nil {
			return dto.Guardians{}, err
		}
	}
	return guardian, nil
}

func (g *guardianRepository) Create(guardian model.Guardians) error {
	query := `INSERT INTO guardians VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := g.db.Exec(query, guardian.Id, guardian.CreatedAt, guardian.UpdatedAt, guardian.Name,
		guardian.Relationship, guardian.PhoneNum, guardian.Email, guardian.Address, guardian.District,
		guardian.City, guardian.Province, guardian.PatientId)
	if err != nil {
		return err
	}
	return nil
}

func (g *guardianRepository) Update(guardian model.Guardians) error {
	query := `UPDATE guardians SET updated_at = ?, name = ?, relationship = ?, phone_num = ?, email = ?,
                     address = ?, district = ?, city = ?, province = ? WHERE id = ?`
	_, err := g.db.Exec(query, guardian.UpdatedAt, guardian.Name, guardian.Relationship, guardian.PhoneNum,
		guardian.Email, guardian.Address, guardian.District, guardian.City, guardian.Province, guardian.Id)
	if err != nil {
		return err
	}
	return nil
}

func (g *guardianRepository) Delete(id string) error {
	query := `DELETE FROM guardians WHERE id = ?`
	_, err := g.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
