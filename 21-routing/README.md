# 21 - Routing in Go with Chi

This lesson introduces **dynamic routing, route grouping, and clean server structure** using the popular lightweight Go router — `chi`.  
You'll move from basic HTTP servers to flexible, professional route handling.

---

## 🎯 Objectives

- Set up a scalable router using [chi](https://github.com/go-chi/chi)
- Implement dynamic URL parameters (e.g., `/users/{id}`)
- Organize route groups (e.g., `/users`, `/admin`)
- Handle unknown routes with a custom 404 fallback
- Validate and respond to dynamic input cleanly

---

## 📦 Project Structure

```
21-routing/
├── go.mod
├── cmd/
│   └── app/
│       └── main.go             → Server setup and router wiring
└── internal/
    └── routes/
        ├── user.go             → /users/{userID} routes
        ├── admin.go            → /admin routes
        └── fallback.go         → Custom fallback 404 handler
```

---

## 🔧 Setup Instructions

1. Initialize module if you haven't:

```bash
cd 21-routing
go mod init github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/21-routing
go get github.com/go-chi/chi/v5
```

2. Run the app:

```bash
go run ./cmd/app
```

✅ Output:

```
🚀 Server running at http://localhost:8080
```

---

## 📚 Example Routes to Test

| Route                 | Behavior                          |
|------------------------|-----------------------------------|
| `/users/42`            | Shows user profile for ID 42      |
| `/users/abc`           | Responds with "Invalid user ID"   |
| `/admin/dashboard`     | Loads the Admin Dashboard         |
| `/notarealpage`        | Triggers a custom 404 fallback    |

---

## 🧠 Key Concepts

| Concept                     | Description |
|------------------------------|-------------|
| `chi.NewRouter()`            | Clean router initialization |
| `Route("/group", fn)`        | Group related routes |
| `chi.URLParam(r, "param")`    | Extract dynamic path parameters |
| Input validation             | Parse and check path parameters safely |
| `r.NotFound(handler)`        | Custom fallback for unknown routes |
| Separation into `internal/routes` | Professional organization for scaling apps |

---

## 🔥 Why This Matters

- **Dynamic URLs** like `/users/{id}` are fundamental to any real-world REST API.
- **Route groups** prevent clutter and make growing servers manageable.
- **Custom 404 handlers** show polish and better control over user experience.
- **Input validation** early in handlers protects your application from bad inputs and errors.
- **Modular structure** mirrors production-level Go backend designs.

---

## 🧩 Extra Credit / Stretch Goals

- Add `/products/{sku}` routes for a mock e-commerce API.
- Add middleware to log request duration and paths.
- Try using chi's built-in `Middleware` features for even more control.

---

## 📎 Resources

- [Go Chi Router Documentation (Official)](https://pkg.go.dev/github.com/go-chi/chi/v5)
- [Chi Router Examples and Tutorials (GitHub Wiki)](https://github.com/go-chi/chi/wiki)
- [Building Go APIs with chi (YouTube Walkthrough)](https://www.youtube.com/watch?v=d_GVNgXZQxE)

---

🔁 Up Next:  
Lesson `22-controllers` — organizing route logic even further into controller packages for large, real-world applications!
```