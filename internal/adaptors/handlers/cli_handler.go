package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fmt.Println("Command: add <task name>, list, done <task id>, rm <task id>, exit")

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
			fmt.Println("test")

			taskName := strings.Join(command[1:], " ")

			err := h.usecase.AddTasks(taskName)
			if err != nil {
				fmt.Printf("Failed to add task: %s\n", err)
				continue
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

		case "done":
			if len(command) != 2 {
				fmt.Println("Usage: done <task id>")
				continue
			}

			id, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Printf("Failed to convert data; please make sure id is integer: %s\n", err)
				continue
			}

			err = h.usecase.TriggerTask(id)
			if err != nil {
				fmt.Printf("Failed to update task id %d: %s", id, err)
				continue
			}

			fmt.Println("Update task success!")

		case "rm":
			if len(command) != 2 {
				fmt.Println("Usage: rm <task id>")
				continue
			}

			id, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Printf("Failed to convert data; please make sure id is integer: %s\n", err)
				continue
			}

			err = h.usecase.RemoveTask(id)
			if err != nil {
				fmt.Printf("Failed to update task id %d: %s", id, err)
				continue
			}

			fmt.Println("Update task success!")

		case "exit":
			fmt.Println(`Goodbye`)
			return

		default:
			fmt.Println(`Invalid option, Please try again`)
			continue
		}
	}
}
