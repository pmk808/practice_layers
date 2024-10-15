package main

import (
	"net/http"
	"taskmanager/handlers"
)

func main() {
	// Routes for GET and POST
	http.HandleFunc("/tasks", handlers.TaskHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
