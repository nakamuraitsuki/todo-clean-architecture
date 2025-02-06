package sqlite3

import (
	"practice/domain/entity"
	"practice/domain/repository"
	"time"

	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
	conn *sqlx.DB
}

func NewTodoRepository(conn *sqlx.DB) repository.ITodoRepository {
	return TodoRepository{conn: conn}
}

type TodoModel struct {
	ID          int        `db:"id"`
	Title       string     `db:"title"`
	CompletedAt *time.Time `db:"completed_at"`
	CreatedAt   time.Time  `db:"created_at"`
}

func (tr TodoRepository) GetTodos() ([]entity.Todo, error) {
	var todos []TodoModel
	err := tr.conn.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	var result []entity.Todo
	for _, todo := range todos {
		result = append(result, entity.NewTodo(todo.Title, todo.CompletedAt, todo.CreatedAt))
	}

	return result, nil
}

func (tr TodoRepository) CreateTodo(params repository.CreateTodoParams) (entity.Todo, error) {
	createdAt := time.Now()
	_, err := tr.conn.Exec("INSERT INTO todos (title, created_at) VALUES (?, ?)", params.Title, createdAt)
	if err != nil {
		return entity.Todo{}, err
	}

	return entity.NewTodo(params.Title, nil, createdAt), nil
}

func MigrateTodo(conn *sqlx.DB) error {
    _, err := conn.Exec(`
        CREATE TABLE IF NOT EXISTS todos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            completed_at TIMESTAMP,
            created_at TIMESTAMP NOT NULL
        )
    `)
    return err
}