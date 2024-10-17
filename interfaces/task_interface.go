package interfaces

import "taskmanager/schemas"

// TaskService defines the service layer functions
type TaskService interface {
	GetTasks() []schemas.Task
	CreateTask(task schemas.Task) (schemas.Task, error)
}

// TaskRepository defines the repository layer functions
type TaskRepository interface {
	FetchTasks() []schemas.Task
	SaveTask(task schemas.Task) schemas.Task
}
