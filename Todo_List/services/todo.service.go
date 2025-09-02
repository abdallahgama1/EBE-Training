package services

import (
	"Todo_List/models"
	"Todo_List/repository"
	"errors"
)

func GetAllTodos() []models.Todo {
	return repository.GetTodos()
}

func GetTodoByID(id int) (*models.Todo, error) {
	return repository.GetTodoByID(id)
}

func CreateTodo(todo models.Todo) (models.Todo, error) {
	if todo.Title == "" {
		return models.Todo{}, errors.New("title cannot be empty")
	}
	return repository.AddTodo(todo), nil
}

func UpdateTodo(id int, todo models.Todo) (*models.Todo, error) {
	if todo.Title == "" {
		return nil, errors.New("title cannot be empty")
	}
	return repository.UpdateTodo(id, todo)
}

func DeleteTodo(id int) error {
	return repository.DeleteTodo(id)
}
