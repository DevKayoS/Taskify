package store

import "time"

type Task struct {
	Id          int32     `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int32     `json:"priority,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskStore interface {
	CreateTask(title, description string, priority int32) (Task, error)
	GetTaskById(id int32) (Task, error)
	ListTasks() ([]Task, error)
	UpdateTask(id int32, title, description string, priority int32) (Task, error)
	DeleteTask(id int32) error
}
