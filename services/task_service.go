package services

import (
	"taskmanager/interfaces"
	"taskmanager/schemas"
)

type TaskService struct {
	repo interfaces.TaskRepository
}

func NewTaskService(repo interfaces.TaskRepository) interfaces.TaskService {
	return &TaskService{repo: repo}
}

// GetTasks fetches tasks
func (s *TaskService) GetTasks() []schemas.Task {
	return s.repo.FetchTasks()
}

// CreateTask creates a new task
func (s *TaskService) CreateTask(task schemas.Task) schemas.Task {
	return s.repo.SaveTask(task)
}
