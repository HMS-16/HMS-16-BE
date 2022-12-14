package admin

import (
	"HMS-16-BE/model"
	"database/sql"
)

type AdminRepository interface {
	Create(admin model.Admins) error
	Login(username string) (model.Admins, error)
	GetById(id string) (model.Admins, error)
	Update(admin model.Admins) error
	Delete(id string) error
}

type adminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) *adminRepository {
	return &adminRepository{db}
}

func (a *adminRepository) Create(admin model.Admins) error {
	query := `INSERT INTO admins(id, created_at, updated_at, username, password, phone_num, email, name) VALUES (?,?,?,?,?,?,?,?)`
	_, err := a.db.Exec(query, admin.ID, admin.CreatedAt, admin.UpdatedAt, admin.Username, admin.Password,
		admin.PhoneNum, admin.Email, admin.Name)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminRepository) Login(username string) (model.Admins, error) {
	var admin model.Admins

	query := `SELECT id, created_at, updated_at, username, password, phone_num, email, name FROM admins WHERE username = ?`
	row, err := a.db.Query(query, username)
	if err != nil {
		return model.Admins{}, err
	}

	defer row.Close()
	row.Next()
	err = row.Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt, &admin.Username,
		&admin.Password, &admin.PhoneNum, &admin.Email, &admin.Name)
	if err != nil {
		return model.Admins{}, err
	}

	return admin, nil
}

func (a *adminRepository) GetById(id string) (model.Admins, error) {
	var admin model.Admins

	query := `SELECT id, created_at, updated_at, username, password, phone_num, email, name FROM admins WHERE id = ?`
	row, err := a.db.Query(query, id)
	if err != nil {
		return model.Admins{}, err
	}

	defer row.Close()
	row.Next()
	err = row.Scan(&admin.ID, &admin.CreatedAt, &admin.UpdatedAt, &admin.Username,
		&admin.Password, &admin.PhoneNum, &admin.Email, &admin.Name)
	if err != nil {
		return model.Admins{}, err
	}

	return admin, nil
}

func (a *adminRepository) Update(admin model.Admins) error {
	query := `UPDATE admins SET updated_at = ?, phone_num = ?, email = ?, name = ? WHERE id = ?`
	_, err := a.db.Exec(query, admin.UpdatedAt, admin.PhoneNum, admin.Email, admin.Name, admin.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a *adminRepository) Delete(id string) error {
	query := `DELETE FROM admins WHERE id = ?`
	_, err := a.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
