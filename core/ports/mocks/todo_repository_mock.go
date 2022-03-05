package mocks_ports

import (
	"errors"

	"github.com/tiagonevestia/hex-go/core/domain"
	"github.com/tiagonevestia/hex-go/core/ports"
)

type toDoRepositoryMemory struct {
	todos []domain.ToDo
}

func NewToDoRepositoryMemory() ports.ToDoRepository {
	return &toDoRepositoryMemory{
		todos: []domain.ToDo{
			{ID: "1", Title: "First", Description: "First description"},
			{ID: "2", Title: "Second", Description: "Second description"},
			{ID: "3", Title: "Third", Description: "Third description"},
		},
	}
}

func (t *toDoRepositoryMemory) Get(id string) (*domain.ToDo, error) {
	for _, todo := range t.todos {
		if todo.ID == id {
			return &todo, nil
		}
	}

	return nil, errors.New("Item not found")
}

func (t *toDoRepositoryMemory) List() ([]domain.ToDo, error) {
	return t.todos, nil
}

func (t *toDoRepositoryMemory) Create(todo *domain.ToDo) (*domain.ToDo, error) {
	newTodo := todo
	newTodo.ID = "4"
	t.todos = append(t.todos, *newTodo)
	return todo, nil
}
