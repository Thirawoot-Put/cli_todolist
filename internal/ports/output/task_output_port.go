package output

import "task_manager/internal/domain"

type TaskOutputPort interface {
	SaveTasks(tasks []domain.Task) error
	LoadTasks() ([]domain.Task, error)
}
