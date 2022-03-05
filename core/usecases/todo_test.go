package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagonevestia/hex-go/core/ports"
	mocks_ports "github.com/tiagonevestia/hex-go/core/ports/mocks"
)

func todoUseCase() ports.ToDoUseCase {
	todoRepo := mocks_ports.NewToDoRepositoryMemory()
	return NewToDoUseCase(todoRepo)
}

func TestNewToDoUseCase(t *testing.T) {
	todoUseCase := todoUseCase()
	assert.NotNil(t, todoUseCase)
}

func TestGet(t *testing.T) {
	todoUseCase := todoUseCase()

	todo, err := todoUseCase.Get("1")
	assert.Nil(t, err)
	assert.Equal(t, "1", todo.ID)

	todo, err = todoUseCase.Get("4")
	assert.NotNil(t, err)
}

func TestList(t *testing.T) {
	todoUseCase := todoUseCase()

	todos, err := todoUseCase.List()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(todos))
}

func TestCreate(t *testing.T) {
	todoUseCase := todoUseCase()

	todo, err := todoUseCase.Create("Fourth", "Fourth description")
	assert.Nil(t, err)
	assert.Equal(t, "4", todo.ID)

	todo, err = todoUseCase.Get("4")
	assert.Nil(t, err)
	assert.Equal(t, "Fourth", todo.Title)
	assert.Equal(t, "Fourth description", todo.Description)
}
