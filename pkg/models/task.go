package models

// Status represents the status of a task
type Status string

const (
	Created   Status = "created"
	Pending   Status = "pending"
	Completed Status = "completed"
)

// Task represents a task with an ID, title, and status
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status Status `json:"status"`
}

// new task with default format
func NewTask(id int, title string) *Task {
	return &Task{
		ID:     id,
		Title:  title,
		Status: Created, // Default status
	}
}
