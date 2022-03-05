package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mocks_ports "github.com/tiagonevestia/hex-go/core/ports/mocks"
)

func TestNewToDoUseCase(t *testing.T) {
	todoRepo := mocks_ports.NewToDoRepositoryMemory()
	todoUseCase := NewToDoUseCase(todoRepo)
	assert.NotNil(t, todoUseCase)
}

func TestGet(t *testing.T) {
	todoRepo := mocks_ports.NewToDoRepositoryMemory()
	todoUseCase := NewToDoUseCase(todoRepo)

	todo, err := todoUseCase.Get("1")
	assert.Nil(t, err)
	assert.Equal(t, "1", todo.ID)

	todo, err = todoUseCase.Get("3")
	assert.NotNil(t, err)
}
