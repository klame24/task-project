package routes

import (
	"net/http"
	"task-project/internal/handlers"
)

func SetupRoutes(taskHandler handlers.TaskHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /tasks", taskHandler.CreateTask)

	return mux
}
