package api

import (
	"assignment2/internal/model"
	"assignment2/internal/store"
)

func CountStats(repo *store.Repository[string, *model.Task]) model.Stats {
	stats := model.Stats{}

	for _, task := range repo.All() {
		switch task.Status {
		case model.StatusPending:
			stats.Submitted++
		case model.StatusRunning:
			stats.InProgress++
		case model.StatusCompleted:
			stats.Completed++
		}
	}

	return stats
}
