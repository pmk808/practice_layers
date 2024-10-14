package repository

import "taskmanager/domain"

// TaskRepository is the interface (menu) that defines the contract for the repository.
type TaskRepository interface {
	Create(task domain.Task) error
	GetAll() ([]domain.Task, error)
}
