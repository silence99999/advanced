package model

const (
	StatusPending   = "PENDING"
	StatusRunning   = "IN_PROGRESS"
	StatusCompleted = "COMPLETED"
)

type Task struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}
