package repositories

import (
	"todolist/app/models"
	"database/sql"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

func (r *TodoRepository) GetTodosByUserID(userID int) ([]models.Todo, error) {
	query := "SELECT ID, TASK, STATUS FROM todos WHERE USER_ID=$1"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) AddTodo(todo *models.Todo) error {
	query := "INSERT INTO todos (USER_ID, Task, Status) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, todo.UserID, todo.Task, todo.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) DeleteTodo(userID int, id int) error {
	query := "DELETE FROM todos WHERE USER_ID=$1 and id=$2"
	_, err := r.DB.Exec(query, userID, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) UpdateTodo(todo *models.Todo, userID int) error {
	query := "UPDATE todos SET Task=$1, Status=$2 WHERE id=$3 and USER_ID=$4"
	_, err := r.DB.Exec(query, todo.Task, todo.Status, todo.ID, userID)
	if err != nil {
		return err
	}
	return nil
}
