package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/subhojyoti99/library-management-system/config"
)

var db *sql.DB

// InitDB initializes the PostgreSQL database connection.
func InitDB() {
	var err error
	db, err = sql.Open("postgres", config.AppConfig.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Check the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
}
