package input

import "task_manager/internal/domain"

type TaskInputPort interface {
	AddTasks(tasks []domain.Task) error
	ReadTasks() ([]domain.Task, error)
}
