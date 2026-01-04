package services

import (
	"context"
	"task-project/internal/models"
	"task-project/internal/repositories"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string) (*models.Task, error)
	// UpdateTask
	// DeleteTask
	// GetAllTasks
	// GetTaskByID
}

type taskService struct {
	r repositories.TaskRepository
}

func NewTaskService(r repositories.TaskRepository) TaskService {
	return &taskService{
		r: r,
	}
}

func (s *taskService) CreateTask(ctx context.Context, title, description string) (*models.Task, error) {
	task := &models.Task{
		Title:       title,
		Description: description,
	}
	err := s.r.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
