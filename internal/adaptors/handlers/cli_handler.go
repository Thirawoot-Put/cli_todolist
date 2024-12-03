package handlers

import (
	"flag"
	"fmt"
	"os"
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
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	removeCmd := flag.NewFlagSet("rm", flag.ExitOnError)

	addTaskName := addCmd.String("name", "", "Name of the task to add")
	doneId := doneCmd.Int("id", 0, "Name of the task to add")
	removeId := removeCmd.Int("id", 0, "Name of the task to add")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'list', or 'done' commands")
		return
	}

	switch os.Args[1] {

	case "add":
		if err := addCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println("Failed to parse add cmd: %w", err)
			return
		}

		if *addTaskName == "" {
			fmt.Println("Task name is required. Use -name <task>")
			return
		}

		err := h.usecase.AddTasks(*addTaskName)
		if err != nil {
			fmt.Printf("Failed to add task: %s\n", err)
			return
		}

		fmt.Println("Add task success!")

	case "list":
		if err := listCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println("Failed to parse list cmd: %w", err)
			return
		}

		tasks, err := h.usecase.ReadTasks()
		if err != nil {
			fmt.Printf("Error to get tasks: %v\n", err)
			return
		} else {
			for _, task := range tasks {
				fmt.Printf("ID: %d, Name: %s, Done: %t \n", task.ID, task.Name, task.Done)
			}
		}

	case "done":
		if err := doneCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println("Failed to parse done cmd: %w", err)
			return
		}
		if *doneId == 0 {
			fmt.Println("Usage: done -id <task id>")
			return
		}

		err := h.usecase.TriggerTask(*doneId)
		if err != nil {
			fmt.Printf("Failed to update task id %d: %s", *doneId, err)
			return
		}

		fmt.Println("Update task success!")

	case "rm":
		if err := removeCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println("Failed to parse rm cmd: %w", err)
			return
		}
		if *removeId == 0 {
			fmt.Println("Usage: rm -id <task id>")
			return
		}

		err := h.usecase.RemoveTask(*removeId)
		if err != nil {
			fmt.Printf("Failed to update task id %d: %s", *removeId, err)
			os.Exit(1)
		}

		fmt.Println("Update task success!")

	case "exit":
		fmt.Println(`Goodbye`)
		return

	default:
		fmt.Println(`Invalid option, Please try again`)
		os.Exit(1)
	}
}
