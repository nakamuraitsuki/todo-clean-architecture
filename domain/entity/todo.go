package entity

import "time"

type Todo struct {
	title       string
	completedAt *time.Time
	createdAt   time.Time
}

func NewTodo(title string, completedAt *time.Time, createdAt time.Time) Todo {
	return Todo{
		title:       title,
		completedAt: completedAt,
		createdAt:   createdAt,
	}
}

func (t Todo) IsAvailable() bool {
	return t.completedAt == nil
}

func (t Todo) GetTitle() string {
	return t.title
}

func (t Todo) GetCreatedAt() time.Time {
	return t.createdAt
}