
# 27 - Sessions in Go (Standard Library)

This lesson demonstrates how to implement basic **session management** using Go’s standard library — no external frameworks. You’ll learn how to issue secure cookies, store in-memory sessions, protect routes using middleware, and pass user data through request context.

---

## 🎯 Objectives

- Simulate login/logout flows without external authentication providers
- Create secure `HttpOnly` session cookies
- Use an in-memory session store (`map[string]string`)
- Enforce protected routes using custom middleware
- Pass user context using `context.WithValue()`

---

## 📁 Project Structure

```

27-sessions-standard/
├── main.go                         # Entry point with server setup
├── handlers/
│   └── handlers.go                 # Public and protected HTTP handlers (login, logout, dashboard)
├── middleware/
│   └── session.go                  # Middleware for validating session and injecting user context
├── internal/
│   └── sessionstore/
│       └── session.go              # In-memory session management (create, get, delete)
├── go.mod
└── README.md                       # You are here

````

---

## 🚀 Run the App

### 1. Open your terminal:

```bash
cd 27-sessions-standard
go run main.go
````

### 2. Try these endpoints:

| Endpoint     | Description                              |
| ------------ | ---------------------------------------- |
| `/`          | Public homepage                          |
| `/login`     | Logs in (creates session + cookie)       |
| `/dashboard` | Protected route (requires valid session) |
| `/logout`    | Logs out and clears session              |

---

## 🧠 Concepts in Use

| Feature               | Purpose                                         |
| --------------------- | ----------------------------------------------- |
| `http.Cookie`         | Stores the session token securely on the client |
| `crypto/rand`         | Generates unpredictable session IDs             |
| `map[string]string`   | Simple server-side session management           |
| Middleware            | Verifies sessions and injects context           |
| `context.WithValue()` | Adds user identity to request context           |
| `SameSiteStrictMode`  | Mitigates CSRF attacks by limiting cookie scope |

---

## ✅ What You Learn

* How to build a secure session workflow using Go’s standard library
* How to separate auth logic using middleware and context
* How to implement login/logout without any frontend
* How to avoid global variables by passing context across layers

---

## 🔐 Security Considerations

* Session cookies are marked `HttpOnly` and `SameSite=Strict`
* Context key uses a custom type to avoid key collisions (`type contextKey string`)
* Session store is in-memory only — ideal for prototypes, not production

---

## ✨ Real-World Use Cases

* Admin dashboards with session-based login
* Internal tools that don't require OAuth or JWTs
* Starter APIs with protected routes

---

## 🧪 Bonus Ideas

* Add session expiration logic (TTL or idle timeout)
* Move session store to Redis or PostgreSQL
* Add roles or permissions to the session context

---

## 🔁 What’s Next?

Lesson 28: **Building Forms with CSRF Protection**
We’ll look at how to protect form submissions and enhance your Go app with secure form-based workflows.

---

```