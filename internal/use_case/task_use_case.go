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

func (u *TaskUseCaseImpl) AddTasks(task string) error {
	tasks, err := u.ReadTasks()
	if err != nil {
		return err
	}

	newTask := domain.Task{
		ID:   len(tasks) + 1,
		Name: task,
		Done: false,
	}

	tasks = append(tasks, newTask)

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

func (u *TaskUseCaseImpl) TriggerTask(id int) error {
	tasks, err := u.repo.LoadTasks()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	for i := range len(tasks) {
		if tasks[i].ID == id {
			tasks[i].Done = !tasks[i].Done
		}
	}

	err = u.repo.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Failed to update task in save period: %w", err)
	}

	return nil
}

func (u *TaskUseCaseImpl) RemoveTask(id int) error {
	tasks, err := u.repo.LoadTasks()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	i := 0
	for idx, task := range tasks {
		if task.ID == id {
			continue
		}

		tasks[i] = tasks[idx]
		i++
	}
	tasks = tasks[:i]

	err = u.repo.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Failed to update task in save period: %w", err)
	}

	return nil
}
