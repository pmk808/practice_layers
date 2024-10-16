package handlers

import (
	"encoding/json"
	"net/http"
	"taskmanager/config"
	"taskmanager/repository"
	"taskmanager/schemas"
	"taskmanager/services"

	_ "github.com/lib/pq"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	// Establish DB connection
	db := config.ConnectDB()
	defer db.Close()

	repo := repository.NewPostgresTaskRepository(db)
	service := services.NewTaskService(repo)

	switch r.Method {
	case http.MethodGet:
		tasks := service.GetTasks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var newTask schemas.Task

		// Read request body
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Create the new task
		createdTask := service.CreateTask(newTask)

		// Respond with the created task
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdTask)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
