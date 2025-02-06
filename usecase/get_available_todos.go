package usecase

import (
	"practice/domain/entity"
	"practice/domain/repository"
)

type GetAvailableTodosOutput struct {
	Todos []entity.Todo
}

type GetAvailableTodosInput struct {
	TodoRepository repository.ITodoRepository
}

func GetAvailableTodos(input GetAvailableTodosInput) (*GetAvailableTodosOutput, error) {
	todos, err := input.TodoRepository.GetTodos()
	if err != nil {
		return nil, err
	}

	availableTodos := []entity.Todo{}

	for _, todo := range todos {
		if todo.IsAvailable() {
			availableTodos = append(availableTodos, todo)
		}
	}

	return &GetAvailableTodosOutput{Todos: availableTodos}, nil
}
