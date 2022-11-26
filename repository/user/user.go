package user

import (
	"HMS-16-BE/dto"
	"HMS-16-BE/model"
	"database/sql"
)

type UserRepository interface {
	Create(user model.Users) error
	Login(username string) (model.Users, error)
	GetAll() ([]dto.User, error)
	GetById(id string) (dto.User, error)
	Update(user model.Users) error
	Delete(id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user model.Users) error {
	query := `INSERT INTO users VALUES (?,?,?,?,?,?,?,?)`
	_, err := u.db.Exec(query, user.Id, user.CreatedAt, user.UpdatedAt, user.Username, user.Password, user.Email,
		user.PhoneNum, user.Role)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Login(username string) (model.Users, error) {
	query := `SELECT * FROM users WHERE username = ?`
	row, err := u.db.Query(query, username)
	if err != nil {
		return model.Users{}, err
	}

	var user model.Users
	defer row.Close()
	row.Next()
	err = row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Password, &user.Email,
		&user.PhoneNum, &user.Role)
	if err != nil {
		return model.Users{}, err
	}

	return user, nil
}

func (u *userRepository) GetAll() ([]dto.User, error) {
	query := `SELECT id, created_at, updated_at, username, email, phone_num, role FROM users`
	row, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []dto.User
	defer row.Close()
	for row.Next() {
		var user dto.User
		err = row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Email, &user.PhoneNum,
			&user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) GetById(id string) (dto.User, error) {
	query := `SELECT id, created_at, updated_at, username, email, phone_num, role FROM users WHERE id = ?`
	row, err := u.db.Query(query, id)
	if err != nil {
		return dto.User{}, err
	}

	var user dto.User
	defer row.Close()
	row.Next()
	err = row.Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Username, &user.Email, &user.PhoneNum,
		&user.Role)
	if err != nil {
		return dto.User{}, err
	}

	return user, nil
}

func (u *userRepository) Update(user model.Users) error {
	query := `INSERT users SET username = ?, email = ?, phone_num = ? WHERE id = ?`
	_, err := u.db.Exec(query, user.Username, user.Email, user.PhoneNum, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
