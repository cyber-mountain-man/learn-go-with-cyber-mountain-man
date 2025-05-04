package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/denisenkom/go-mssqldb"
)

func Connect() *sql.DB {
	// Load .env if available
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, falling back to environment variables")
	}

	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	server := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	database := os.Getenv("DBNAME")

	if user == "" || password == "" || server == "" || port == "" || database == "" {
		log.Fatal("Missing one or more required MSSQL environment variables")
	}

	// Format MSSQL connection string
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, server, port, database)

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatalf("Failed to open MSSQL connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to MSSQL: %v", err)
	}

	return db
}
