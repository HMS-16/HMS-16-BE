package user

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"database/sql"
)

type UserRepository interface {
	Create(user model.Users) error
	Login(email string) (model.Users, error)
	GetAll() ([]dto.User, error)
	GetById(id string) (dto.User, error)
	Update(user model.Users) error
	UpdatePassword(user model.Users) error
	Delete(id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user model.Users) error {
	query := `INSERT INTO users(id, created_at, updated_at, name, password, email, role, str_num) VALUES (?,?,?,?,?,?,?,?)`
	_, err := u.db.Exec(query, user.Id, user.CreatedAt, user.UpdatedAt, user.Name, user.Password, user.Email,
		user.Role, user.STRNum)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Login(email string) (model.Users, error) {
	query := `SELECT id, created_at, updated_at, name, password, email, role, str_num FROM users WHERE email = ?`
	row, err := u.db.Query(query, email)
	if err != nil {
		return model.Users{}, err
	}

	var user model.Users
	defer row.Close()
	row.Next()
	err = row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.Password, &user.Email,
		&user.Role, &user.STRNum)
	if err != nil {
		return model.Users{}, err
	}

	return user, nil
}

func (u *userRepository) GetAll() ([]dto.User, error) {
	query := `SELECT str_num, created_at, updated_at, name, email, role FROM users WHERE role = 1 OR role = 2`
	row, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []dto.User
	defer row.Close()
	for row.Next() {
		var user dto.User
		err = row.Scan(&user.StrNum, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) GetById(id string) (dto.User, error) {
	query := `SELECT str_num, created_at, updated_at, name, email, role FROM users WHERE str_num = ? AND (role = 1 OR role = 2)`
	row, err := u.db.Query(query, id)
	if err != nil {
		return dto.User{}, err
	}

	var user dto.User
	defer row.Close()
	row.Next()
	err = row.Scan(&user.StrNum, &user.CreatedAt, &user.UpdatedAt, &user.Name, &user.Email, &user.Role)
	if err != nil {
		return dto.User{}, err
	}

	return user, nil
}

func (u *userRepository) Update(user model.Users) error {
	query := `UPDATE users SET updated_at = ?, name = ?, email = ? WHERE str_num = ?`
	_, err := u.db.Exec(query, user.UpdatedAt, user.Name, user.Email, user.STRNum)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) UpdatePassword(user model.Users) error {
	query := `UPDATE users SET updated_at = ?, password = ? WHERE str_num = ?`
	_, err := u.db.Exec(query, user.UpdatedAt, user.Password, user.STRNum)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE str_num = ?`
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
