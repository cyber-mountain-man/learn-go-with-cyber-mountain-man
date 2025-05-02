package main

import (
	"log"

	// Custom database package from internal directory
	"learn-go-with-cyber-mountain-man/26-databases/postgres/internal/db"
)

func main() {
	// Establish a connection to the PostgreSQL database using shared Connect function
	conn := db.Connect()
	// Ensure the database connection is closed when this function completes
	defer conn.Close()

	// SQL statement to create the 'users' table if it doesn't already exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,             -- Auto-incrementing unique ID
		name TEXT NOT NULL,                -- User's name (required)
		email TEXT NOT NULL UNIQUE,        -- Email must be unique (required)
		created_at TIMESTAMPTZ DEFAULT NOW() -- Timestamp of user creation (defaults to current time)
	);`

	// Execute the SQL statement to create the table
	_, err := conn.Exec(createTableQuery)
	if err != nil {
		// Log and terminate if table creation fails
		log.Fatalf("Error creating users table: %v", err)
	}

	// Informational log message if successful
	log.Println("Users table created successfully.")
}

/*
🧠 DATABASE INITIALIZATION — SCHEMA SETUP TOOL (cmd/setup/main.go)

✅ What Happens Here:
- This standalone command-line tool connects to the database and ensures the `users` table exists.
- It uses `CREATE TABLE IF NOT EXISTS` to make the operation idempotent (safe to run multiple times).
- It runs independently of the web server to keep schema management separate.

✅ Why This Matters:
- This is your first step toward writing real **DevOps scripts** or database migration tools in Go.
- Running schema setup as a separate program improves reliability, especially in containerized deployments.
- Keeping schema creation outside the server avoids performance issues and race conditions.

✅ Key Concepts:
| Concept                         | Explanation |
|----------------------------------|-------------|
| `cmd/setup` pattern              | Convention for small, self-contained CLI apps in Go projects |
| `Exec()` with raw SQL            | Executes statements that don't return rows (e.g., CREATE TABLE) |
| `CREATE TABLE IF NOT EXISTS`     | Safe schema creation without overwriting existing data |
| Logging to the console           | Useful for CI/CD, cloud init logs, and troubleshooting |

🔔 Bonus Tip:
- This setup script could later evolve into a full migration engine (e.g., version-controlled schema updates).
- Consider separating setup logic into a reusable package (e.g., `internal/migrations`) for larger apps.

📚 Up Next:
- Add seed data automatically for local dev
- Integrate versioned migrations using tools like `golang-migrate`
*/
