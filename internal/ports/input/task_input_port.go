package input

import "task_manager/internal/domain"

type TaskInputPort interface {
	AddTasks(task string) error
	ReadTasks() ([]domain.Task, error)
	TriggerTask(id int) error
}
