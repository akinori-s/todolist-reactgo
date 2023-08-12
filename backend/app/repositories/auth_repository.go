package repositories

import (
	"todolist/app/models"
	"database/sql"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (r *AuthRepository) GetUserByUsername(username string) (*models.Auth, error) {
	user := new(models.Auth)

	row := r.DB.QueryRow("SELECT ID, Username, Email, Password_hash FROM users WHERE username=$1", username)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)
	return user, err
}
