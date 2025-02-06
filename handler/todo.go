package handler

import (
	"net/http"
	"practice/domain/repository"
	"practice/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoRepository  repository.ITodoRepository
}

func NewTodoHandler (todoRepository repository.ITodoRepository) TodoHandler {
	return TodoHandler{todoRepository: todoRepository}
}

func (t TodoHandler) Register (g *echo.Group) {
	g.GET("/available", t.GetAvailableTodos)
}

type GetAvailableTodosResponseTodo struct {
	title string `json:"title"`
	createdAt time.Time `json:"created_at"`
}

type GetAvailableTodosResponse struct {
	todos []GetAvailableTodosResponseTodo `json:"todos"`
}

func (t TodoHandler) GetAvailableTodos(c echo.Context) error {
	input := usecase.GetAvailableTodosInput{TodoRepository: t.todoRepository}
	output, err := usecase.GetAvailableTodos(input)
	if err != nil {
		return err
	}

	res := GetAvailableTodosResponse{
		todos: []GetAvailableTodosResponseTodo{},		
	}

	for _, todo := range output.Todos {
		res.todos = append(res.todos, GetAvailableTodosResponseTodo{
			title: todo.GetTitle(),
			createdAt: todo.GetCreatedAt(),
		})
	}

	return c.JSON(http.StatusOK, res)
}