# 23 - Middleware in Go (with chi)

Middleware is a powerful pattern in Go that allows you to intercept, modify, or reject HTTP requests **before** they reach your route handlers. In this lesson, you'll build and compose your own middleware for logging and access control.

---

## 🎯 Objectives

- Understand what middleware is and how it works
- Build reusable middleware for logging and authentication
- Apply global and route-specific middleware with chi
- Create clean, composable HTTP server layers

---

## 📁 Project Structure

```
23-middleware/
├── go.mod
├── cmd/
│   └── app/
│       └── main.go              → Server entry point
└── internal/
    ├── middleware/
    │   ├── logger.go            → Logs request method/path and duration
    │   └── auth.go              → Mock header-based authentication
    └── handlers/
        └── routes.go            → Public and private route logic
```

---

## 🚀 How to Run

```bash
cd 23-middleware
go run ./cmd/app
```

You’ll see output like:

```
🚀 Server running at http://localhost:8080
```

---

## 🔐 Test It with curl

### ✅ Public Route (No Auth)

```bash
curl http://localhost:8080/public
```

Expected Output:

```
🌐 Welcome to the public endpoint!
```

---

### ❌ Private Route (No Header)

```bash
curl http://localhost:8080/private
```

Expected Output:

```
Unauthorized - missing or invalid X-Auth header
```

---

### ✅ Private Route (With Header)

```bash
curl -H "X-Auth: secret123" http://localhost:8080/private
```

Expected Output:

```
🔐 Welcome to the private endpoint! You are authenticated.
```

---

## 🧠 Concepts Covered

| Concept                  | Description |
|--------------------------|-------------|
| `r.Use(middleware)`      | Globally applies middleware to all routes |
| `r.Group()`              | Wraps middleware around specific route groups |
| `next.ServeHTTP(w, r)`   | Passes control to the next handler in the chain |
| Header-based filtering   | Common strategy in API key or token validation |
| Timing requests          | Useful for performance logging and tracing |

---

## 📌 Why Middleware Matters

✅ Separates logic cleanly — handlers stay focused on business rules  
✅ Adds flexibility — reuse logic like logging, CORS, auth across services  
✅ Enhances observability — log, trace, and debug requests efficiently  
✅ Enables scalable security — gatekeep critical routes without duplication  

Middleware is the glue of robust Go web servers.  
It enforces the right rules in the right places — **before** requests reach your logic.

---

🔁 Up Next:  
**Lesson 24: Templates** — render dynamic HTML content with Go’s `html/template` engine and build web-ready frontends.
```