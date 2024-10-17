package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"taskmanager/repository"
	"taskmanager/schemas"
	"taskmanager/services"
	"taskmanager/validation"
)

func TaskHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	repo := repository.NewPostgresTaskRepository(db) // Pass the db connection here
	validator := validation.NewTaskValidator()       // New validation service
	service := services.NewTaskService(repo, validator)

	switch r.Method {
	case http.MethodGet:
		tasks := service.GetTasks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var newTask schemas.Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		createdTask, err := service.CreateTask(newTask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdTask)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
