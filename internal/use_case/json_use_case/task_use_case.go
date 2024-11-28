package jsonusecase

import (
	"fmt"
	"task_manager/internal/domain"
	"task_manager/internal/infrastructure/repository"
	usecase "task_manager/internal/use_case"
)

type JsonUseCaseImpl struct {
	repo repository.TaskRepository
}

func NewJsonUseCase(repo repository.TaskRepository) usecase.TaskUseCase {
	return &JsonUseCaseImpl{
		repo: repo,
	}
}

func (u *JsonUseCaseImpl) AddTasks(tasks []domain.Task) error {
	if err := u.repo.SaveTasks(tasks); err != nil {
		return fmt.Errorf(`%s`, err)
	}

	return nil
}

func (u *JsonUseCaseImpl) ReadTasks() ([]domain.Task, error) {
	data, err := u.repo.LoadTasks()
	if err != nil {
		return nil, fmt.Errorf(`%s`, err)
	}

	return data, nil
}
