package user

// User defines the core domain model for a system user.
type User struct {
	Name  string // User's full name
	Email string // User's email address
	Admin bool   // Whether the user has admin privileges
}

// GetDefaultUser returns a predefined sample user.
// In real apps, this could load from a database, environment, or config file.
func GetDefaultUser() User {
	return User{
		Name:  "Luigi Mario",
		Email: "luigi@example.com",
		Admin: true,
	}
}

/*
🧠 INTERNAL PACKAGE WALKTHROUGH

✅ Purpose:
- Encapsulate application-specific entities (`User`) and logic (`GetDefaultUser`).

✅ Key Concepts:
| Concept            | Description |
|--------------------|-------------|
| `internal/`         | Restricted access: only code inside the module can import it |
| `User` struct       | Domain object modeling a real-world user |
| `GetDefaultUser()`  | Simulates fetching/configuring a user |

🔔 Lesson:
- Use `internal/` for protected application internals.
- Avoid leaking internal details into external/public code.
*/
