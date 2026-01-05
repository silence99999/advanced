package model

type Stats struct {
	Submitted  int `json:"submitted"`
	InProgress int `json:"in_progress"`
	Completed  int `json:"completed"`
}
