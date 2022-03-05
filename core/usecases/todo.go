package usecases

import (
	"os/exec"

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

func (t *toDoUseCase) List() ([]domain.ToDo, error) {
	todos, err := t.todoRepo.List()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *toDoUseCase) Create(title, description string) (*domain.ToDo, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return nil, err
	}

	todo := domain.NewTodo(
		string(out),
		title,
		description,
	)

	todo, err = t.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
