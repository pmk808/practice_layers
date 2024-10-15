package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"taskmanager/repository"
	"taskmanager/schemas"
	"taskmanager/services"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewFakeTaskRepository()
	service := services.NewTaskService(repo)

	switch r.Method {
	case http.MethodGet:
		tasks := service.GetTasks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var newTask schemas.Task

		// Use io.ReadAll instead of ioutil.ReadAll
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Unmarshal request body into newTask
		err = json.Unmarshal(body, &newTask)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
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
