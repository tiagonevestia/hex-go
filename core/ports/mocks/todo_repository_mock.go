package mocks_ports

import (
	"errors"

	"github.com/tiagonevestia/hex-go/core/domain"
	"github.com/tiagonevestia/hex-go/core/ports"
)

type toDoRepositoryMemory struct{}

func NewToDoRepositoryMemory() ports.ToDoRepository {
	return &toDoRepositoryMemory{}
}

func (t *toDoRepositoryMemory) Get(id string) (*domain.ToDo, error) {
	todos := []*domain.ToDo{
		{ID: "1", Title: "First", Description: "First description"},
		{ID: "2", Title: "Second", Description: "Second description"},
	}

	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return nil, errors.New("Item not found")
}
