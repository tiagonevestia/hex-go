package ports

import "github.com/tiagonevestia/hex-go/core/domain"

type ToDoRepository interface {
	Get(id string) (*domain.ToDo, error)
}
