package main

import (
	"taskmanager/domain"
	"taskmanager/handler"
	"taskmanager/repository"
	"taskmanager/service"
)

func main() {
	// Step 1: Create a new in-memory repository
	taskRepo := &repository.InMemoryTaskRepository{
		Tasks: []domain.Task{},
	}

	// Step 2: Create a new task service, injecting the repository
	taskService := service.NewTaskService(taskRepo)

	// Step 3: Create a new task handler, injecting the service
	taskHandler := handler.NewTaskHandler(taskService)

	// Simulate task creation
	taskHandler.CreateTask("Learn Go")
	taskHandler.CreateTask("Build a Go project")

	// Simulate listing tasks
	taskHandler.ShowTasks()
}
