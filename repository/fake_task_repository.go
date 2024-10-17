package repository

import (
	"taskmanager/interfaces"
	"taskmanager/schemas"
)

type FakeTaskRepository struct {
	tasks []schemas.Task
}

// NewFakeTaskRepository creates a new instance of FakeTaskRepository
func NewFakeTaskRepository() interfaces.TaskRepository {
	return &FakeTaskRepository{
		tasks: []schemas.Task{
			{ID: 1, Name: "Task 1"},
			{ID: 2, Name: "Task 2"},
		},
	}
}

func (r *FakeTaskRepository) FetchTasks() []schemas.Task {
	return r.tasks
}

func (r *FakeTaskRepository) SaveTask(task schemas.Task) schemas.Task {
	task.ID = len(r.tasks) + 1 // Generate a new ID
	r.tasks = append(r.tasks, task)
	return task
}
