package main

import (
	"fmt"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/19-project-setup/internal/user"
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/19-project-setup/pkg/utils"
)

func main() {
	// Display a startup message
	fmt.Println("🚀 Starting the app...")

	// Call internal function to get a default User object
	u := user.GetDefaultUser()

	// Use a utility function to nicely print user information
	utils.PrintUser(u)
}

/*
🧠 MAIN FILE WALKTHROUGH

✅ Purpose:
- Acts as the entry point for the application.
- Connects internal logic (domain objects) with reusable utilities.

✅ Key Concepts:
| Concept                | Description |
|-------------------------|-------------|
| `cmd/app/main.go`       | Organized entry for CLI or service |
| Import from `internal/` | Protected domain-specific logic |
| Import from `pkg/`      | Shared helper functions |
| Layer separation        | App startup doesn't need to know inner logic details |

🔔 Lesson:
- Main stays thin (only wiring things together).
- Business logic lives elsewhere (clean separation of concerns).
*/
