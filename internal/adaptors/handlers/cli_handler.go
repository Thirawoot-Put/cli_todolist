package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nWelcome to Task Manager")
	fmt.Println("Command: add <task name>, list, done <task id>, exit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		command := strings.Fields(input)

		if len(command) == 0 {
			fmt.Println("Please enter a command")
		}

		switch command[0] {

		case "add":
			if len(command) < 2 {
				fmt.Println("Usage: add <task name>")
				continue
			}

			taskName := strings.Join(command[1:], " ")

			err := h.usecase.AddTasks(taskName)
			if err != nil {
				fmt.Println("Failed to add task: %w", err)
			}

			fmt.Println("Add task success!")

		case "list":
			tasks, err := h.usecase.ReadTasks()
			if err != nil {
				fmt.Printf("Error to get tasks: %v\n", err)
				continue
			} else {
				for _, task := range tasks {
					fmt.Printf("ID: %d, Name: %s, Done: %t \n", task.ID, task.Name, task.Done)
				}
			}

		case "exit":
			fmt.Println(`Goodbye`)
			return

		default:
			fmt.Println(`Invalid option, Please try again`)
			continue
		}
	}
}
