package main

import (
	"task_manager/internal/adaptors/handlers"
	"task_manager/internal/adaptors/repositories"
	usecase "task_manager/internal/use_case"
)

func main() {
	r := repositories.NewJsonRepository()
	u := usecase.NewTaskUseCase(r)
	c := handlers.NewCliHandler(u)

	c.CliTaskManager()
}
