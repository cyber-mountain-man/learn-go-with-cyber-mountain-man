# 19 - Project Setup in Go

In this lesson, you’ll structure a Go project the **professional way**, separating your logic into commands, internal packages, and shared utilities.  
This is the foundation for scalable applications, APIs, and production-level Go projects.

---

## 🎯 Objectives

- Initialize a Go module
- Organize code into `cmd/`, `internal/`, and `pkg/` directories
- Separate core domain models from utility functions
- Build a simple CLI-style app that prints user information

---

## 📦 Project Structure

```
learn-go-with-cyber-mountain-man/
└── 19-project-setup/
    ├── go.mod
    ├── cmd/
    │   └── app/
    │       └── main.go          → App entry point
    ├── internal/
    │   └── user/
    │       └── user.go          → User model and logic
    └── pkg/
        └── utils/
            └── helpers.go       → Utility for printing users
```

---

## 🔧 Step-by-Step Setup

1. Create the folders and files manually, or use scaffolding tools.

2. From inside `19-project-setup/`, initialize the module:

```bash
go mod init github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/19-project-setup
```

3. Your imports in Go files will match your module path automatically.

4. Run the app:

```bash
go run ./cmd/app
```

✅ Output:

```
🚀 Starting the app...
🧑 User Info:
  Name:  Luigi Mario
  Email: Luigi@example.com
  Admin: true
```

---

## 🧠 Key Concepts

| Concept            | Description |
|--------------------|-------------|
| `cmd/`             | Holds the main application entry point (can have multiple apps later) |
| `internal/`        | Private code only accessible inside this module (enforced by Go) |
| `pkg/`             | Publicly reusable code across apps or external modules |
| Clean Separation    | Keeps main logic isolated from startup wiring |
| go.mod              | Declares module path for import resolution |

---

## 🔥 Why This Matters

- **Scalability:** Add new services easily without cluttering your root folder.
- **Maintainability:** Each folder has a clear responsibility.
- **Security:** Internal code can’t be imported accidentally from outside the repo.
- **Professionalism:** Mirrors real Go projects used by companies and open-source tools.

---

## 📚 Best Practices

- Only expose public APIs through `pkg/`, not `internal/`.
- Main files (`cmd/app/main.go`) should be thin — only handle wiring, not business logic.
- Prefer small, focused packages instead of giant monolith files.
- Document key structures and flows as your project grows.

---

🔁 Up Next:  
Lesson `20-handlers` — start building more organized HTTP route handling using real structuring techniques for web servers!
```
