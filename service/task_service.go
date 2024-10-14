package service

import (
    "taskmanager/domain"
    "taskmanager/repository"
)

// TaskService defines the services related to tasks.
type TaskService struct {
    repo repository.TaskRepository
}

// NewTaskService creates a new instance of TaskService, with dependency injection.
func NewTaskService(repo repository.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

// AddTask adds a new task by calling the repository's Create function.
func (s *TaskService) AddTask(name string) error {
    tasks, err := s.repo.GetAll() // Get both tasks and error
    if err != nil {
        return err // Return the error if something goes wrong
    }

    task := domain.Task{
        ID:     len(tasks) + 1, // Use the length of the tasks slice
        Name:   name,
        Status: "Pending",
    }
    return s.repo.Create(task) // Call the repository to create the task
}

// ListTasks lists all tasks by calling the repository's GetAll function.
func (s *TaskService) ListTasks() ([]domain.Task, error) {
    return s.repo.GetAll()
}
