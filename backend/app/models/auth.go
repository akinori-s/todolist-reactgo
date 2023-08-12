package models

type Auth struct {
	ID				int		`json:"id"`
	Username		string	`json:"username"`
	Email			string	`json:"email"`
	PasswordHash	string	`json:"-"`
}
