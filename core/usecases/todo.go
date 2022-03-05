package usecases

import (
	"github.com/tiagonevestia/hex-go/core/domain"
	"github.com/tiagonevestia/hex-go/core/ports"
)

type toDoUseCase struct {
	todoRepo ports.ToDoRepository
}

func NewToDoUseCase(todoRepo ports.ToDoRepository) ports.ToDoUseCase {
	return &toDoUseCase{
		todoRepo: todoRepo,
	}
}

func (t *toDoUseCase) Get(id string) (*domain.ToDo, error) {
	todo, err := t.todoRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
