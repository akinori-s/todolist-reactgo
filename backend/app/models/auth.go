package models

type Auth struct {
	ID				int		`json:"id", omitempty`
	FirstName		string	`json:"firstName", omitempty`
	LastName		string	`json:"lastName", omitempty`
	Email			string	`json:"email", omitempty`
	Password		string	`json:"password, omitempty"`
	PasswordConfirmation	string	`json:"passwordConfirmation, omitempty"`
	PasswordHash	string	`json:"-, omitempty"`
}
