package deployment

import (
	"database/sql"
	"log"
)

func SetupDB(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}
}
