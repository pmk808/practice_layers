package interfaces

import "taskmanager/schemas"

type TaskService interface {
	GetTasks() []schemas.Task
	CreateTask(task schemas.Task) (schemas.Task, error)
}

type TaskRepository interface {
	FetchTasks() []schemas.Task
	SaveTask(task schemas.Task) schemas.Task
}
