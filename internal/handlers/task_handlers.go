package handlers

import (
	"encoding/json"
	"net/http"
	"task-project/internal/models"
	"task-project/internal/services"
)

type TaskHandler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	DoneTask(w http.ResponseWriter, r *http.Request)
	GetAllTasks(w http.ResponseWriter, r *http.Request)
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
		return
	}

	task, err := h.s.CreateTask(r.Context(), req.Title, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *taskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var taskID int

	if err := json.NewDecoder(r.Body).Decode(&taskID); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.s.DeleteTask(r.Context(), taskID)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskID)
}

func (h *taskHandler) DoneTask(w http.ResponseWriter, r *http.Request) {
	var taskID int

	if err := json.NewDecoder(r.Body).Decode(&taskID); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.s.DoneTask(r.Context(), taskID)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskID)
}

func (h *taskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.s.GetAllTasks(r.Context())
	if err != nil {
		http.Error(w, "Something going wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
