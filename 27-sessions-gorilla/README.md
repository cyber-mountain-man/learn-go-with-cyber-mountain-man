# 27 - Sessions in Go (Gorilla Edition)

This lesson introduces **secure, encrypted sessions** in Go using the powerful and production-tested [`gorilla/sessions`](https://github.com/gorilla/sessions) package. You’ll build a full session-authenticated app with login, logout, protected routes, and secure cookie handling.

---

## 🎯 Objectives

* Understand the basics of session management using `gorilla/sessions`
* Create login/logout routes that set and clear session cookies
* Protect access to authenticated routes using middleware
* Store session secrets in a secure `.env` file
* Validate sessions and prevent unauthorized access

---

## 📁 Folder Structure

```
27-sessions-gorilla/
├── main.go                           # Starts the web server and sets up routes
├── handlers/
│   └── handlers.go                   # Login, logout, dashboard, home logic
├── middleware/
│   └── session.go                    # Middleware to enforce session authentication
├── internal/
│   └── config/
│       └── secrets.go                # Session store config (auth + encryption keys)
├── .env                              # Session keys for secure cookie signing (ignored by Git)
└── README.md                         # This file
```

---

## 🔐 Session Security Overview

* We use `CookieStore`, which signs (and optionally encrypts) session data.
* Sessions are **stateless** — no DB or Redis required (but can be added later).
* Session cookies use `HttpOnly`, `SameSite`, and are stored in-memory on the server.

---

## 🧪 How to Run and Test

### 1. Prepare the `.env` file

```env
SESSION_AUTH_KEY=your_64_char_hex_key_here
SESSION_ENCRYPT_KEY=your_32_char_hex_key_here
```

You can generate secure keys using:

```sh
openssl rand -hex 64  # For SESSION_AUTH_KEY
openssl rand -hex 32  # For SESSION_ENCRYPT_KEY
```

### 2. Run the Server

```bash
go run main.go
```

Console should show:

```
🍪 Gorilla session server running at http://localhost:8080
```

---

## 🌐 Endpoints

| Endpoint     | Method | Description                        |
| ------------ | ------ | ---------------------------------- |
| `/`          | GET    | Public homepage                    |
| `/login`     | GET    | Creates session and sets cookie    |
| `/logout`    | GET    | Deletes session and cookie         |
| `/dashboard` | GET    | Protected route (requires session) |

---

## ✅ Example Test with PowerShell

```powershell
$response = Invoke-WebRequest -Uri http://localhost:8080/login -SessionVariable session
$response.Content  # ✅ Logged in

Invoke-WebRequest -Uri http://localhost:8080/dashboard -WebSession $session
```

Expected Output:

```
🔐 Welcome to your dashboard, demo_user
```

---

## 🧠 Key Concepts

| Concept                 | Purpose                                      |
| ----------------------- | -------------------------------------------- |
| `gorilla/sessions`      | Secure, signed, optionally encrypted cookies |
| `CookieStore`           | Stores session data inside browser cookies   |
| `session.Save(r, w)`    | Commits session changes to client            |
| `.env` secrets loading  | Ensures safe storage of session keys         |
| `middleware/session.go` | Blocks unauthenticated access to routes      |

---

## ⚠️ Gotchas

* Session key size must be correct:

  * Auth key = 64 bytes (128 hex chars)
  * Encrypt key = 32 bytes (64 hex chars)
* Session won't persist if `.Save()` is not called.
* If PowerShell isn't showing cookies, check `Set-Cookie` headers using:

  ```powershell
  $response.Headers["Set-Cookie"]
  ```

---

## 🔁 What’s Next?

Lesson 28: **Deployment and Production Tips**
Learn how to deploy your Go web applications securely with session support.

---