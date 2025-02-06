package entity

import "time"

type Todo struct {
	title       string
	completedAt *time.Time
	createdAt   time.Time
}
