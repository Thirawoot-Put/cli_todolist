package usecase

import "task_manager/internal/domain"

type TaskUseCase interface {
	AddTasks(tasks []domain.Task) error
	ReadTasks() ([]domain.Task, error)
}
