package repository

import "taskmanager/domain"

// InMemoryTaskRepository is an implementation of TaskRepository using in-memory storage.
type InMemoryTaskRepository struct {
	Tasks []domain.Task // Make this field exported by capitalizing it
}

// Create adds a new task to the in-memory slice.
func (repo *InMemoryTaskRepository) Create(task domain.Task) error {
	repo.Tasks = append(repo.Tasks, task) // Use the exported field here
	return nil
}

// GetAll returns all the tasks in memory.
func (repo *InMemoryTaskRepository) GetAll() ([]domain.Task, error) {
	return repo.Tasks, nil // Use the exported field here
}
