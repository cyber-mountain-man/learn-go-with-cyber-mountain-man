package db

import (
	"database/sql"        // Go's standard SQL package for DB interaction
	"fmt"                 // For string formatting (used in connection string)
	"os"                  // For reading environment variables

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/models" // Importing the User struct
	_ "github.com/denisenkom/go-mssqldb" // MS SQL Server driver — the underscore means it’s imported for its side-effects (driver registration)
)

var DB *sql.DB // Global variable to hold the database connection pool

// InitDB initializes the database connection using env vars and sets the global DB variable.
func InitDB() error {
	// Create connection string using env vars
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
		os.Getenv("DBUSER"),     // Username for DB
		os.Getenv("DBPASSWORD"), // Password
		os.Getenv("DBHOST"),     // Hostname (e.g., "localhost" or "mssql" in Docker)
		os.Getenv("DBPORT"),     // Port number (usually 1433 for MSSQL)
		os.Getenv("DBNAME"),     // Database name
	)

	// Open a connection pool
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}

	// Ping to confirm connection is working
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	// Assign to the global variable
	DB = db
	return nil
}

// GetAllUsers retrieves all users from the 'users' table.
func GetAllUsers() ([]models.User, error) {
	rows, err := DB.Query(`SELECT id, name, email FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure the rows are closed after we're done

	var users []models.User
	for rows.Next() {
		var u models.User
		// Scan values into struct fields
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// InsertUser adds a new user to the database using parameterized SQL.
func InsertUser(name, email string) error {
	_, err := DB.Exec(`INSERT INTO users (name, email) VALUES (@p1, @p2)`, name, email)
	return err
}

// GetUserByID fetches a single user by ID.
func GetUserByID(id int) (models.User, error) {
	var u models.User
	err := DB.QueryRow(`SELECT id, name, email FROM users WHERE id = @p1`, id).
		Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

// UpdateUser modifies an existing user's name and email based on ID.
func UpdateUser(id int, name, email string) error {
	_, err := DB.Exec(`UPDATE users SET name = @p1, email = @p2 WHERE id = @p3`, name, email, id)
	return err
}

// DeleteUser removes a user by ID.
func DeleteUser(id int) error {
	_, err := DB.Exec(`DELETE FROM users WHERE id = @p1`, id)
	return err
}

/*
🧠 Blurb: Understanding db.go
This Go file provides a full database access layer for managing users in a Microsoft SQL Server database. It defines a global DB connection pool initialized using environment variables, ensuring reusability across handlers.

Each function corresponds to a common operation (CRUD):

GetAllUsers: Reads all users.

InsertUser: Adds a user.

GetUserByID: Fetches a single user.

UpdateUser: Edits a user.

DeleteUser: Removes a user.

It uses parameterized SQL (@p1, @p2) to prevent SQL injection and relies on the standard database/sql package for safe and efficient database interaction.
*/