package main

import (
	"log"
	"net/http"
	"practice/handler"
	"practice/infrastructure/sqlite3"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	todoGroup := e.Group("/todo")
	todoRepository := sqlite3.NewTodoRepository(db)
	handler.NewTodoHandler(todoRepository).Register(todoGroup)

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":8000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
