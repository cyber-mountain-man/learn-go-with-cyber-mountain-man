package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Loads environment variables from .env file
	_ "github.com/lib/pq"      // PostgreSQL driver — blank import to register with database/sql
)

// Connect initializes and returns a PostgreSQL database connection using environment variables.
func Connect() *sql.DB {
	// Attempt to load variables from .env file (optional, useful for local dev)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	// Fetch PostgreSQL connection parameters from environment
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	sslmode := os.Getenv("PGSSLMODE")

	// Validate presence of all required environment variables
	if user == "" || password == "" || dbname == "" || host == "" || port == "" {
		log.Fatal("Missing required PostgreSQL environment variables")
	}

	// Format the DSN (Data Source Name) string required by lib/pq
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	// Open the database connection (driver name must match the imported driver)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	// Ping the database to ensure the connection is valid
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	// Return a pointer to the initialized sql.DB connection pool
	return db
}

/*
🧠 DATABASE CONNECTION — ENVIRONMENT-BASED CONFIG (db/connect.go)

✅ What Happens Here:
- This file defines the `Connect()` function, which initializes a PostgreSQL connection using environment variables.
- It loads variables from a `.env` file using `godotenv`, making it easier to manage secrets and settings locally.
- If `.env` isn't found, it safely falls back to system environment variables — great for production containers.
- It checks for required variables and fails fast if anything is missing.

✅ Why This Matters:
- Environment-based configuration is a **12-Factor App** principle and essential for writing scalable, deployable systems.
- Using `database/sql` + `lib/pq` gives you direct access to PostgreSQL without heavy ORMs — lightweight and performant.
- Centralizing the DB connection setup avoids duplication and supports maintainability.

✅ Key Concepts:
| Concept                   | Explanation |
|---------------------------|-------------|
| `godotenv.Load()`         | Loads key-value pairs from a `.env` file into `os.Getenv()` |
| `sql.Open`                | Initializes a database pool (lazy connection) |
| `db.Ping()`               | Verifies the connection is actually live |
| `log.Fatal()`             | Exits immediately if required config is missing |
| Blank import `_ "lib/pq"` | Registers the PostgreSQL driver with `database/sql` |

🔔 Bonus Tip:
- This setup makes it easy to deploy in Docker, Kubernetes, or CI/CD by swapping `.env` files with environment variable injection.
- Connection pooling is handled automatically by `sql.DB` — no need to manage individual connections manually.

📚 Up Next:
- Add query timeouts using `context.Context`
- Extract DB health checks into a `/health` endpoint
- Add retry logic for connection failures during startup
*/
