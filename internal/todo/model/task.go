// internal/todo/task.go
package model

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
