package repositories

import (
	"database/sql"
	"errors"
	"go-rest-api/internal/model"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	user.Id = uuid.New()

	query := `INSERT INTO users (id, email, full_name, password) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, user.Id, user.Email, user.FullName, user.Password)
	if err != nil {
		return errors.New("failed to create user: " + err.Error())
	}

	return nil
}

func (r *UserRepository) GetUserById(userId uuid.UUID) (*model.User, error) {
	query := `SELECT id, email, full_name, password FROM users WHERE id = $1`

	row := r.DB.QueryRow(query, userId)

	var user model.User
	err := row.Scan(&user.Id, &user.Email, &user.FullName, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("failed to get user: " + err.Error())
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	query := `UPDATE users SET email = $1, full_name = $2, password = $3 WHERE id = $4`

	res, err := r.DB.Exec(query, user.Email, user.FullName, user.Password, user.Id)
	if err != nil {
		return errors.New("failed to update user: " + err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("no user updated")
	}

	return nil
}

func (r *UserRepository) RemoveUser(userId uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	res, err := r.DB.Exec(query, userId)
	if err != nil {
		return errors.New("failed to delete user: " + err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("no user deleted")
	}

	return nil
}
