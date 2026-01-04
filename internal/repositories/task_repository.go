package repositories

import (
	"context"
	"task-project/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Task) error
	DeleteTask(ctx context.Context, taskID int) error
	GetAllTasks(ctx context.Context) ([]*models.Task, error)
	// GetTaskByID(ctx context.Context, taskID int) (*models.Task, error)
	DoneTask(ctx context.Context, taskID int) error
}

type taskRepository struct {
	ctx context.Context
	db  *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) CreateTask(ctx context.Context, task *models.Task) error {
	task.Completed = false
	task.CreatedAt = time.Now()

	sqlQuery := `INSERT INTO tasks(title, description, completed, created_at)
	VALUES ($1, $2, $3, $4);`

	_, err := r.db.Exec(
		ctx,
		sqlQuery,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
	)

	return err
}

func (r *taskRepository) UpdateTask() {}

func (r *taskRepository) DeleteTask(ctx context.Context, taskID int) error {
	sqlQuery := `DELETE FROM tasks WHERE tasks.id = $1`
	_, err := r.db.Exec(
		ctx,
		sqlQuery,
		taskID,
	)

	return err
}

func (r *taskRepository) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	sqlQuery := `
		SELECT * FROM tasks;
	`

	rows, err := r.db.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*models.Task

	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, taskID int) (*models.Task, error) {
	sqlQuery := `
		SELECT * FROM tasks
		WHERE tasks.id = $1
	`

	row, err := r.db.Query(ctx, sqlQuery, taskID)
	if err != nil {
		return nil, err
	}

	task := &models.Task{}

	err = row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) DoneTask(ctx context.Context, taskID int) error {
	query := `
		UPDATE tasks
		SET completed=true
		WHERE tasks.id = $1;
	`

	_, err := r.db.Exec(ctx, query, taskID)

	return err
}
