package models

type Todo struct {
	ID		int		`json:"id"`
	UserID	int		`json:"userId", omitempty`
	Task	string	`json:"text"`
	Status	bool	`json:"completed"`
}
