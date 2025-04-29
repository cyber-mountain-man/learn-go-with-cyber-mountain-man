# 22 - Controllers in Go: Struct-Based Route Logic

In this lesson, you’ll learn how to build **modular, scalable route logic** using controller structs and method binding.  
This mirrors how professional APIs and backend services are built in Go — with clarity, testability, and flexibility in mind.

---

## 🎯 Objectives

- Create reusable controller structs for route logic
- Bind HTTP handler methods to controller receivers
- Cleanly separate route wiring from business logic
- Prepare your app for testing, dependency injection, and growth

---

## 📦 Project Structure

```
22-controllers/
├── go.mod
├── cmd/
│   └── app/
│       └── main.go              → Entry point: server setup
└── internal/
    ├── routes/
    │   └── router.go            → Mounts controller routes
    └── controllers/
        └── user.go              → UserController (handlers for /users)
```

---

## 🚀 Getting Started

### 1. Initialize the module (if not done yet)

```bash
cd 22-controllers
go mod init github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/22-controllers
go get github.com/go-chi/chi/v5
```

---

### 2. Run the App

```bash
go run ./cmd/app
```

✅ Output:
```
🚀 Server running at http://localhost:8080
```

---

## 🔍 Try These Routes

| Route            | Description                        |
|------------------|------------------------------------|
| `/users/42`      | Returns mock user profile (validated ID) |
| `/users/abc`     | Returns "Invalid user ID" error     |
| `/users`         | Returns mock user list              |

---

## 🧠 Key Concepts

| Concept                     | What You Learn |
|-----------------------------|----------------|
| Controller struct           | Groups related handlers cleanly |
| Method receiver handlers    | Enables logic reuse and injection |
| chi routing                 | Professional HTTP route handling |
| Separation of concerns      | Keeps `main.go`, routing, and logic cleanly isolated |
| Path parameter validation   | Ensures safe, correct request processing |

---

## 💡 Why This Pattern Matters

- ✅ **Modular** – Each controller handles its own domain
- ✅ **Testable** – Logic can be unit-tested without wiring
- ✅ **Scalable** – Easy to add new services or dependencies
- ✅ **Professional** – Mirrors enterprise Go server design

This is how modern Go APIs are structured — from small tools to large SaaS backends.

---

## 📎 Resources

- [Chi Router Documentation](https://pkg.go.dev/github.com/go-chi/chi/v5)
- [Go Web Application Layouts](https://github.com/golang-standards/project-layout)
- [Effective Go – Structs & Methods](https://go.dev/doc/effective_go#methods)

---

🔁 Up Next:  
Lesson `23-middleware` — apply global and per-route middleware to secure, log, or enhance your server's behavior!
```