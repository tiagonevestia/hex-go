package ports

import "github.com/tiagonevestia/hex-go/core/domain"

type ToDoUseCase interface {
	Get(id string) (*domain.ToDo, error)
}
