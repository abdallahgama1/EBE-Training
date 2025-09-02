package repository

import (
	"Todo_List/models"
	"errors"
)

var todos = []models.Todo{}
var nextID = 1

func GetTodos() []models.Todo {
	return todos
}

func GetTodoByID(id int) (*models.Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func AddTodo(todo models.Todo) models.Todo {
	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	return todo
}

func UpdateTodo(id int, updated models.Todo) (*models.Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = updated.Title
			todos[i].Completed = updated.Completed
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func DeleteTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
