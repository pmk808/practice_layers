package main

import (
	"log"
	"net/http"
	"os"
	"taskmanager/config"
	"taskmanager/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the PostgreSQL database
	db := config.ConnectDB()
	defer db.Close()

	// Set up routes and pass the db connection to the TaskHandler
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.TaskHandler(db, w, r) // Pass db here
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
