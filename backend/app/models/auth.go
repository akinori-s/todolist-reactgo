package models

type Auth struct {
	ID				int		`json:"id"`
	FirstName		string	`json:"firstName"`
	LastName		string	`json:"lastName"`
	Email			string	`json:"email"`
	Password		string	`json:"password, omitempty"`
	PasswordConfirmation	string	`json:"passwordConfirmation, omitempty"`
	PasswordHash	string	`json:"-, omitempty"`
}
