package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"assignment2/internal/model"
	"assignment2/internal/queue"
	"assignment2/internal/store"
)

type API struct {
	repo  *store.Repository[string, *model.Task]
	queue *queue.TaskQueue
}

func New(repo *store.Repository[string, *model.Task], q *queue.TaskQueue) *API {
	return &API{
		repo:  repo,
		queue: q,
	}
}

func (a *API) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Task string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	task := &model.Task{
		ID:     uuid.NewString(),
		Task:   req.Task,
		Status: model.StatusPending,
	}

	a.repo.Set(task.ID, task)
	a.queue.Push(task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (a *API) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a.repo.All())
}

func (a *API) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	task, ok := a.repo.Get(id)
	if !ok {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (a *API) GetStatsHandler(w http.ResponseWriter, _ *http.Request) {
	stats := CountStats(a.repo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func (a *API) SlowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(4 * time.Second)
	w.Write([]byte("done"))
}
