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

func (r *AuthRepository) GetUserByEmail(email string) (*models.Auth, error) {
	user := new(models.Auth)

	row := r.DB.QueryRow("SELECT ID, FirstName, LastName, Email, Password_hash FROM users WHERE email=$1", email)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
	)
	return user, err
}

func (r *AuthRepository) CreateUser(user *models.Auth) error {
	query := "INSERT INTO users (FirstName, LastName, Email, Password_hash) VALUES ($1, $2, $3, $4) RETURNING ID"
	row := r.DB.QueryRow(query, user.FirstName, user.LastName, user.Email, user.PasswordHash)
	err := row.Scan(&user.ID)
	return err
}
