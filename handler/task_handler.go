package handler

import (
	"fmt"
	"taskmanager/service"
)

// TaskHandler will be responsible for handling user interaction.
type TaskHandler struct {
	taskService *service.TaskService
}

// NewTaskHandler initializes a TaskHandler with a TaskService.
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: service}
}

// CreateTask creates a new task and prints a confirmation message.
func (h *TaskHandler) CreateTask(name string) {
	err := h.taskService.AddTask(name)
	if err != nil {
		fmt.Println("Error creating task:", err)
	} else {
		fmt.Println("Task created successfully!")
	}
}

// ShowTasks lists all tasks.
func (h *TaskHandler) ShowTasks() {
	tasks, err := h.taskService.ListTasks()
	if err != nil {
		fmt.Println("Error listing tasks:", err)
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("- %d: %s [%s]\n", task.ID, task.Name, task.Status)
	}
}
