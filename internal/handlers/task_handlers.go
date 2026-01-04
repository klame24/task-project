package handlers

import (
	"encoding/json"
	"net/http"
	"task-project/internal/models"
	"task-project/internal/services"
)

type TaskHandler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	s services.TaskService
}

func NewTaskHandler(s services.TaskService) TaskHandler {
	return &taskHandler{
		s: s,
	}
}

func (h *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	req := models.Request{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	task, err := h.s.CreateTask(r.Context(), req.Title, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
