package usecase

import (
	"fmt"
	"task_manager/internal/domain"
	"task_manager/internal/ports/input"
	"task_manager/internal/ports/output"
)

type TaskUseCaseImpl struct {
	repo output.TaskOutputPort
}

func NewTaskUseCase(repo output.TaskOutputPort) input.TaskInputPort {
	return &TaskUseCaseImpl{
		repo: repo,
	}
}

func (u *TaskUseCaseImpl) AddTasks(tasks []domain.Task) error {
	if err := u.repo.SaveTasks(tasks); err != nil {
		return fmt.Errorf(`%s`, err)
	}

	return nil
}

func (u *TaskUseCaseImpl) ReadTasks() ([]domain.Task, error) {
	data, err := u.repo.LoadTasks()
	if err != nil {
		return nil, fmt.Errorf(`%s`, err)
	}

	return data, nil
}
