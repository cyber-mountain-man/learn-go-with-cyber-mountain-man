package utils

import (
	"fmt"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/19-project-setup/internal/user"
)

// PrintUser prints detailed user information in a human-readable format.
func PrintUser(u user.User) {
	fmt.Println("🧑 User Info:")
	fmt.Printf("  Name:  %s\n", u.Name)
	fmt.Printf("  Email: %s\n", u.Email)
	fmt.Printf("  Admin: %v\n", u.Admin)
}

/*
🧠 PKG PACKAGE WALKTHROUGH

✅ Purpose:
- Provide shared, reusable functions that can be used by multiple parts of the app.

✅ Key Concepts:
| Concept           | Description |
|-------------------|-------------|
| `pkg/`             | Public helpers, accessible anywhere inside or outside the app |
| `PrintUser(u User)`| Prints any user nicely without embedding business rules |

🔔 Lesson:
- Keep utility functions pure and independent of application specifics.
- Think "small, clean, reusable" for `pkg/`.
*/
