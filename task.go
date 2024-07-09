package main

import (
	"job_task/handlers"
	"job_task/repositories"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	taskRepositories := repositories.TaskLogicRepository()
    h := handlers.TaskHandler(taskRepositories)
    e.POST("/task-one", h.TaskOne)
    e.POST("/task-two", h.TaskTwo)
    e.POST("/task-three", h.TaskThree)

    e.Start(":5050")
}
