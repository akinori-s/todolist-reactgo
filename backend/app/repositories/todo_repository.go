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

func (r *TodoRepository) GetTodos() ([]models.Todo, error) {
	query := "SELECT * FROM todos"
	rows, err := r.DB.Query(query)
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
	query := "INSERT INTO todos (Task, Status) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, todo.Task, todo.Status)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) DeleteTodo(id int) error {
	query := "DELETE FROM todos WHERE id=$1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) UpdateTodo(todo *models.Todo, id int) error {
	query := "UPDATE todos SET Task=$1, Status=$2 WHERE id=$3"
	_, err := r.DB.Exec(query, todo.Task, todo.Status, id)
	if err != nil {
		return err
	}
	return nil
}
