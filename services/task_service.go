package services

import (
    "taskmanager/interfaces"
    "taskmanager/schemas"
    // "errors"
)

type TaskService struct {
    repo       interfaces.TaskRepository
    validator  interfaces.ValidationService
}

func NewTaskService(repo interfaces.TaskRepository, validator interfaces.ValidationService) interfaces.TaskService {
    return &TaskService{repo: repo, validator: validator}
}

// GetTasks fetches tasks
func (s *TaskService) GetTasks() []schemas.Task {
    return s.repo.FetchTasks()
}

// CreateTask creates a new task after validation
func (s *TaskService) CreateTask(task schemas.Task) (schemas.Task, error) {
    err := s.validator.ValidateTask(task)
    if err != nil {
        return schemas.Task{}, err
    }
    return s.repo.SaveTask(task), nil
}
