package repository

import (
	"database/sql"
	"log"
	"taskmanager/interfaces"
	"taskmanager/schemas"

	_ "github.com/lib/pq"
)

type PostgresTaskRepository struct {
	db *sql.DB
}

func NewPostgresTaskRepository(db *sql.DB) interfaces.TaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) FetchTasks() []schemas.Task {
	rows, err := r.db.Query("SELECT id, name FROM tasks")
	if err != nil {
		log.Fatalf("Error fetching tasks: %v", err)
	}
	defer rows.Close()

	var tasks []schemas.Task
	for rows.Next() {
		var task schemas.Task
		if err := rows.Scan(&task.ID, &task.Name); err != nil {
			log.Fatalf("Error scanning task: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks
}

// SaveTask saves a new task to the PostgreSQL database
func (r *PostgresTaskRepository) SaveTask(task schemas.Task) schemas.Task {
	err := r.db.QueryRow(
		"INSERT INTO tasks (name) VALUES ($1) RETURNING id",
		task.Name,
	).Scan(&task.ID)
	if err != nil {
		log.Fatalf("Error saving task: %v", err)
	}

	return task
}
