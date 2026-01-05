package worker

import (
	"assignment2/internal/model"
	"assignment2/internal/queue"
	"assignment2/internal/store"
	"log"
	"time"
)

func Start(id int, q *queue.TaskQueue, repo *store.Repository[string, *model.Task]) {
	for task := range q.Channel() {
		task.Status = model.StatusRunning
		log.Printf("worker %d processing %s\n", id, task.ID)

		time.Sleep(2 * time.Second)

		task.Status = model.StatusCompleted
		repo.Set(task.ID, task)

		log.Printf("worker %d completed %s\n", id, task.ID)
	}
}
