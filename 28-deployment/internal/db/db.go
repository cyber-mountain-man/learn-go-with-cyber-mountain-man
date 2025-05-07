package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/models"
	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDB() error {
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	)

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	DB = db
	return nil
}


func GetAllUsers() ([]models.User, error) {
	rows, err := DB.Query(`SELECT id, name, email FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func InsertUser(name, email string) error {
	_, err := DB.Exec(`INSERT INTO users (name, email) VALUES (@p1, @p2)`, name, email)
	return err
}

func GetUserByID(id int) (models.User, error) {
	var u models.User
	err := DB.QueryRow(`SELECT id, name, email FROM users WHERE id = @p1`, id).
		Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

func UpdateUser(id int, name, email string) error {
	_, err := DB.Exec(`UPDATE users SET name = @p1, email = @p2 WHERE id = @p3`, name, email, id)
	return err
}

func DeleteUser(id int) error {
	_, err := DB.Exec(`DELETE FROM users WHERE id = @p1`, id)
	return err
}
