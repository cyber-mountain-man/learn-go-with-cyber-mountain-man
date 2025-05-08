package models

// User represents the structure of a user in the system.
// It is used for both database records and JSON serialization.
type User struct {
	ID    int    `json:"id"`    // Unique identifier for the user
	Name  string `json:"name"`  // Name of the user
	Email string `json:"email"` // Email address of the user
}

/*
 Blurb: Purpose of the User Model
This User struct defines a shared data format for users across your application. It plays three important roles:

Database mapping – Represents rows in your users table for scanning (rows.Scan(&u.ID, ...)).

JSON serialization – Enables JSON encoding/decoding via json.Marshal and json.Unmarshal, thanks to the struct tags (json:"...").

Data exchange – Used by both API routes and template handlers to carry consistent user data between layers.

Having a centralized User model promotes type safety, code clarity, and reduces duplication across your handlers and db logic.
*/