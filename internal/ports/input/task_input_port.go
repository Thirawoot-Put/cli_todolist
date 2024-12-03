package input

import "task_manager/internal/domain"

type TaskInputPort interface {
	AddTasks(task string) error
	ReadTasks() ([]domain.Task, error)
	// UpdateTaskName(id int, newTaskName string) error
}
