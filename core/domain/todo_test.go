package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	todo := NewTodo("1", "title", "description")

	assert.Equal(t, "1", todo.ID)
	assert.Equal(t, "title", todo.Title)
	assert.Equal(t, "description", todo.Description)
}

func TestTodoString(t *testing.T) {
	todo := NewTodo("1", "title", "description")

	assert.Equal(t, "title description", todo.String())
}
