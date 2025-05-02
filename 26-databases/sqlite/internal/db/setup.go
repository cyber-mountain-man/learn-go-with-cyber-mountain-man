package db

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver anonymously
)

var DB *sql.DB // Expose a global database connection (used by other packages)

// InitDB initializes the SQLite connection and ensures the schema is set up.
// It stores the connection in the package-level `DB` variable.
func InitDB(relativePath string) error {
	// Resolve full path (optional safety)
	fullPath := filepath.Clean(relativePath)

	// Open a SQLite database file (creates it if not present)
	db, err := sql.Open("sqlite3", fullPath)
	if err != nil {
		return err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return err
	}

	log.Println("📦 Connected to SQLite DB:", fullPath)

	// Save the connection globally for reuse
	DB = db

	// Create the users table if it doesn't exist
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	);
	`

	if _, err := DB.Exec(schema); err != nil {
		return err
	}

	log.Println("✅ Users table ensured.")
	return nil
}
