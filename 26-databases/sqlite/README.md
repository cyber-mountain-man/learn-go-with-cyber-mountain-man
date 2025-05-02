# Lesson 26 - SQLite with Go (CRUD API)

This lesson demonstrates how to build a simple RESTful API using:

- 📦 SQLite (file-based database)
- 🧠 Go's `database/sql` standard library
- 🔧 Chi router (`go-chi/chi`)
- 🔁 Full CRUD operations (Create, Read, Update, Delete)

---

## 🗂 Folder Structure

```

sqlite/
├── data/                  # SQLite DB file lives here
├── internal/
│   ├── db/                # Database connection & table init
│   │   └── setup.go
│   ├── handlers/          # HTTP handlers (business logic)
│   │   └── user.go
│   ├── models/            # SQL access layer (queries)
│   │   └── user.go
│   └── routes/            # All route registration
│       └── routes.go
├── go.mod
└── main.go

````

---

## 🚀 Running the App

```bash
cd sqlite
go run main.go
````

Expected output:

```
📦 Connected to SQLite DB: sqlite/data/app.db
✅ Users table ensured.
🚀 Server running at http://localhost:8080
```

---

## 🧪 Testing the API

### ✅ Create User (POST)

#### Windows CMD:

```cmd
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d "{\"name\":\"Mario\",\"email\":\"mario@nintendo.com\"}"
```

#### PowerShell:

```powershell
curl -Method POST -Uri http://localhost:8080/users -Body '{"name":"Mario","email":"mario@nintendo.com"}' -ContentType "application/json"
```

#### Linux/macOS:

```bash
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"Mario","email":"mario@nintendo.com"}'
```

---

### 📋 Get All Users (GET)

```bash
curl http://localhost:8080/users
```

---

### 🔍 Get User by ID (GET)

```bash
curl http://localhost:8080/users/1
```

---

### ✏️ Update User (PUT)

```bash
curl -X PUT http://localhost:8080/users/1 -H "Content-Type: application/json" -d '{"name":"Luigi","email":"luigi@nintendo.com"}'
```

---

### ❌ Delete User (DELETE)

```bash
curl -X DELETE http://localhost:8080/users/1
```

---

## 🧠 What You Learned

* How to connect Go to SQLite
* How to organize models, handlers, and routes
* How to build a clean REST API using Go’s idioms
* How to test APIs with curl from various environments

---

## 🧼 Tips

* You can inspect `sqlite/data/app.db` using DB Browser for SQLite
* Avoid hardcoding paths — consider using `env` in real-world apps
* Use `log.Fatal` to catch issues on startup

---

🔁 Coming Next: PostgreSQL version (`/postgres/`) and then MSSQL for production scenarios.

```
---