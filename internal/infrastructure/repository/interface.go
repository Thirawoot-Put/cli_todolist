package repository

import "task_manager/internal/domain"

type TaskRepository interface {
	SaveTasks(tasks []domain.Task) error
	LoadTasks() ([]domain.Task, error)
}
