package routes

import (
	"net/http"
	"task-project/internal/handlers"
)

func SetupRoutes(taskHandler handlers.TaskHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /tasks", taskHandler.CreateTask)
	mux.HandleFunc("DELETE /tasks", taskHandler.DeleteTask)
	mux.HandleFunc("PUT /tasks", taskHandler.DoneTask)
	mux.HandleFunc("GET /tasks", taskHandler.GetAllTasks)

	return mux
}
