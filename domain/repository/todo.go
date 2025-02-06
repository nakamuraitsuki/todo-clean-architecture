package repository

import "practice/domain/entity"

type CreateTodoParams struct {
	Title string
}

type ITodoRepository interface {
	GetTodos() ([]entity.Todo, error)
	CreateTodo(params CreateTodoParams) (entity.Todo, error)
}
