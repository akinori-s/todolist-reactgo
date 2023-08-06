package models

type Todo struct {
	ID		int		`json:"id"`
	Task	string	`json:"text"`
	Status	bool	`json:"completed"`
}
