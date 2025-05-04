package main

import (
	"log"

	"learn-go-with-cyber-mountain-man/26-databases/mssql/internal/db"
)

func main() {
	conn := db.Connect()
	defer conn.Close()

	createTableQuery := `
	IF NOT EXISTS (
		SELECT * FROM sysobjects WHERE name='users' AND xtype='U'
	)
	BEGIN
		CREATE TABLE users (
			id INT IDENTITY(1,1) PRIMARY KEY,
			name NVARCHAR(100) NOT NULL,
			email NVARCHAR(100) NOT NULL UNIQUE,
			created_at DATETIME2 DEFAULT SYSDATETIME()
		)
	END
	`

	_, err := conn.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	log.Println("Users table created successfully.")
}
