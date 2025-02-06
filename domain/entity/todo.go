package entity

import "time"

type Todo struct {
	title       string
	completedAt *time.Time
	createdAt   time.Time
}

func (t Todo) IsAvailable() bool {
	return t.completedAt == nil
}