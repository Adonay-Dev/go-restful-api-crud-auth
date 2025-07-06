package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserModel struct {
	DB *sql.DB
}
type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"fullName"`
	Password string    `json:"-"`
}
