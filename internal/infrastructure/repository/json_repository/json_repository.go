package jsonrepository

import (
	"encoding/json"
	"fmt"
	"os"
	"task_manager/internal/domain"
	"task_manager/internal/infrastructure/ports/output"
)

const filename = "task.json"

type JsonRepository struct{}

func NewJsonRepository() output.TaskOutputPort {
	return &JsonRepository{}
}

func (r *JsonRepository) SaveTasks(tasks []domain.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (r *JsonRepository) LoadTasks() ([]domain.Task, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []domain.Task{}, nil // Return an empty slice if the file doesn't exist
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var tasks []domain.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	return tasks, nil
}
