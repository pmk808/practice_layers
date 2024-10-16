package services

import (
	"log"
	"taskmanager/interfaces"
	"taskmanager/schemas"
	"taskmanager/validation"
)

type TaskService struct {
	repo interfaces.TaskRepository
}

func NewTaskService(repo interfaces.TaskRepository) interfaces.TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() []schemas.Task {
	return s.repo.FetchTasks()
}

func (s *TaskService) CreateTask(task schemas.Task) schemas.Task {
	err := validation.ValidateTask(task)
	if err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	return s.repo.SaveTask(task)
}
