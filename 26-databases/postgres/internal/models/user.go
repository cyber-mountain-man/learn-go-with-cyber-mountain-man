package models

import (
	"database/sql"
	"time"
)

// User defines the structure of the user entity that maps directly to the 'users' table in PostgreSQL.
// JSON struct tags are included to ensure the fields serialize correctly when returning API responses.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// GetAllUsers retrieves all records from the users table.
// It returns a slice of User structs or an error if the query fails.
func GetAllUsers(db *sql.DB) ([]User, error) {
	users := []User{} // initialize empty slice to avoid returning null

	// Execute SELECT query
	rows, err := db.Query("SELECT id, name, email, created_at FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through result rows and map to User struct
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	// Return all users or any error encountered during row iteration
	return users, rows.Err()
}

// InsertUser inserts a new user into the database using the provided User struct.
// It returns the new user's ID and sets the CreatedAt field in the original struct to match the database value.
func InsertUser(db *sql.DB, u User) (int, error) {
	var id int
	var createdAt time.Time

	// INSERT with RETURNING allows us to get generated fields immediately
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	err := db.QueryRow(query, u.Name, u.Email).Scan(&id, &createdAt)
	u.CreatedAt = createdAt

	return id, err
}

// GetUserByID retrieves a user record from the database by its unique ID.
// Returns a fully populated User struct or an error if the user is not found.
func GetUserByID(db *sql.DB, id int) (User, error) {
	var u User
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	return u, err
}

// UpdateUser modifies the name and email of an existing user based on their ID.
// Returns an error if the update fails.
func UpdateUser(db *sql.DB, u User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := db.Exec(query, u.Name, u.Email, u.ID)
	return err
}

// DeleteUser removes a user from the database based on their ID.
// Returns an error if the deletion fails.
func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

/*
🧠 GO + POSTGRESQL CRUD — DATA MODEL WALKTHROUGH (models/user.go)

✅ What Happens Here:
- We define a `User` struct to model a database row from the `users` table.
- This file handles **all database operations** (Create, Read, Update, Delete) using raw SQL and Go’s `database/sql` package.
- Each exported function connects to the database and performs SQL queries with proper parameterization.
- We cleanly separate the database access logic from HTTP handler logic (defined elsewhere in `handlers/user.go`).

✅ Why This Matters:
- This is the foundation of the **data layer** in real-world web applications.
- Decoupling DB logic from HTTP routing makes your application modular, testable, and maintainable.
- You learn to safely handle query parameters using `?` (SQLite) or `$1, $2...` (PostgreSQL) — protecting against SQL injection.
- Understanding how to map database rows into Go structs is essential for building backends and REST APIs.

✅ Key Concepts:
| Concept                        | Explanation |
|-------------------------------|-------------|
| `sql.DB.QueryRow / Query`     | Executes SQL queries and retrieves rows |
| `Scan()`                      | Transfers column values into Go variables |
| `RETURNING` in PostgreSQL     | Retrieves `id` and `created_at` after inserting a row |
| Parameterized queries         | Prevent SQL injection and improve clarity |
| Separation of concerns        | DB access logic stays here, API routes use these functions |

🔔 Bonus Tip:
- Initializing `[]User{}` avoids returning `null` on empty queries — instead you get an empty list (`[]`), which is API-friendly.
- This layer is portable! You could reuse these functions in CLI apps, services, or jobs outside of HTTP handlers.

📚 Up Next:
- Build test coverage for these model functions
- Add pagination or filtering support (e.g., `GetUsers(limit, offset)`)
- Move toward using interfaces for easier mocking/testing in large projects
*/
