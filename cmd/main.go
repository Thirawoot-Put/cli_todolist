package main

import (
	"task_manager/internal/adaptors/handlers"
	"task_manager/internal/adaptors/repositories"
	"task_manager/internal/services"
)

func main() {
	r := repositories.NewJsonRepository()
	u := services.NewTaskService(r)
	c := handlers.NewCliHandler(u)

	c.CliTaskManager()
}
