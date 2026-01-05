package queue

import "assignment2/internal/model"

type TaskQueue struct {
	ch chan *model.Task
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		ch: make(chan *model.Task, 100),
	}
}

func (q *TaskQueue) Push(task *model.Task) {
	q.ch <- task
}

func (q *TaskQueue) Channel() <-chan *model.Task {
	return q.ch
}

func (q *TaskQueue) Close() {
	close(q.ch)
}
