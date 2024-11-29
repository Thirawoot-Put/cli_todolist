package handlers

import (
	"fmt"
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
	for {
		fmt.Println("\nTask Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Printf("Get choice error: %s", err)
		}

		switch choice {

		case 1:
			var name string
			fmt.Println("Enter your task name")

			_, err := fmt.Scanln(&name)
			if err != nil {
				fmt.Printf("Get choice error: %s", err)
				continue
			}

			err = h.usecase.AddTasks(name)

			if err != nil {
				fmt.Printf("Error to add task: %v\n", err)
				continue
			} else {
				fmt.Println(`Add task success!`)
			}

		case 2:
			tasks, err := h.usecase.ReadTasks()
			if err != nil {
				fmt.Printf("Error to get tasks: %v\n", err)
				continue
			} else {
				for _, task := range tasks {
					fmt.Printf("ID: %d, Name: %s, Done: %t \n", task.ID, task.Name, task.Done)
				}
			}

		case 3:
			fmt.Println(`Goodbye`)
			return

		default:
			fmt.Println(`Invalid option, Please try again`)
			continue
		}
	}
}
