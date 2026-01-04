package services

import (
	"context"
	"task-project/internal/models"
	"task-project/internal/repositories"
)

type TaskService interface {
	CreateTask(ctx context.Context, title, description string) (*models.Task, error)
	DoneTask(ctx context.Context, taskID int) error
	DeleteTask(ctx context.Context, taskID int) error
	GetAllTasks(ctx context.Context) ([]*models.Task, error)
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

func (s *taskService) DeleteTask(ctx context.Context, taskID int) error {
	err := s.r.DeleteTask(ctx, taskID)

	return err
}

func (s *taskService) DoneTask(ctx context.Context, taskID int) error {
	err := s.r.DoneTask(ctx, taskID)

	return err
}

func (s *taskService) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	tasks, err := s.r.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
