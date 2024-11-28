package handlers

import (
	"fmt"
	"task_manager/internal/domain"
	"task_manager/internal/ports/input"
)

type CliHandler struct {
	usecase input.TaskInputPort
}

func NewCliHandler(uc input.TaskInputPort) *CliHandler {
	return &CliHandler{
		usecase: uc,
	}
}

func (h *CliHandler) CliTaskManager() {
	fmt.Println("\nTask Manager")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Exit")
	fmt.Print("Choose an option: ")

	var choice int
	if _, err := fmt.Scanln(&choice); err != nil {
		fmt.Printf(`Error to receive choice: %v\n`, err)
	}

	if choice == 1 {
		var name string
		if _, err := fmt.Scanln(&choice); err != nil {
			fmt.Printf(`Error to receive task name: %v\n`, err)
		}

		err := h.usecase.AddTasks(domain.Task{ID: 1, Name: name, Done: false})

		if err != nil {
			fmt.Printf(`Error to add task: %v\n`, err)
		} else {
			fmt.Println(`Add task success!`)
		}
	} else if choice == 2 {
		// get tasks
		tasks, err := h.usecase.ReadTasks()
		if err != nil {
			fmt.Printf(`Error to get tasks: %v\n`, err)
		} else {
			for _, task := range tasks {
				fmt.Printf(`ID: %d, Name: %s, Done: %t \n`, task.ID, task.Name, task.Done)
			}
		}
	} else if choice == 3 {
		fmt.Println(`Goodbye`)
	} else {
		fmt.Println(`Invalid option, Please try again`)
	}
}
