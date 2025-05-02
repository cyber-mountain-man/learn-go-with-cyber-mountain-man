# 26 - Databases in Go (PostgreSQL Version)

This lesson teaches how to connect a Go application to a PostgreSQL database using the `database/sql` package. You'll build a full CRUD REST API for managing users, using Chi for routing, and structuring your application into clean layers for DB access, request handling, and server setup.

---

## 🎯 Objectives

- Use `database/sql` to interact with PostgreSQL
- Connect to the database using environment variables and `godotenv`
- Structure database access via a reusable `models` package
- Create handler functions for GET, POST, PUT, DELETE
- Use Chi router to build clean, RESTful endpoints
- Apply best practices for modular design in Go

---

## 📁 Project Structure

```

learn-go-with-cyber-mountain-man/
└── 26-databases/
└── postgres/
├── cmd/
│   └── setup/                 # One-time table creation logic
│       └── main.go
├── internal/
│   ├── db/                    # DB connection logic (Connect)
│   ├── handlers/              # API route handlers (List, Create, Update, Delete users)
│   └── models/                # Data access layer (SQL queries)
├── .env                       # Environment variables (excluded from Git)
├── main.go                    # Server entry point (routes + startup)
├── go.mod / go.sum
└── README.md

````

---

## 🚀 Run the App

### 1. Navigate to the project directory

```bash
cd 26-databases/postgres
````

### 2. Create a `.env` file (or export these values manually)

```env
PGUSER=lesson26_user
PGPASSWORD=lesson26_pass
PGDATABASE=lesson26_postgres
PGHOST=localhost
PGPORT=5432
PGSSLMODE=disable
```

### 3. Run the setup tool to create the `users` table

```bash
go run ./cmd/setup
```

### 4. Start the web server

```bash
go run main.go
```

### 5. Test with curl or a REST client

```bash
curl http://localhost:8080/users
```

---

## 🧠 Concepts in Use

| Feature                   | Purpose                                               |
| ------------------------- | ----------------------------------------------------- |
| `database/sql`            | Core Go package for SQL DBs (PostgreSQL via `lib/pq`) |
| `godotenv`                | Loads `.env` config during local development          |
| `Chi router`              | Lightweight routing with support for path params      |
| `http.HandlerFunc`        | Wraps route handlers that take DB connection          |
| Parameterized SQL queries | Prevent SQL injection                                 |
| Layered architecture      | Separates models, handlers, and setup code            |

---

## 🧠 What You Learn

✅ How to securely connect to a PostgreSQL DB using Go
✅ How to structure a Go app into models, handlers, and entry points
✅ How to build a REST API that supports all CRUD operations
✅ How to read config from a `.env` file and avoid hardcoding credentials
✅ How to return structured JSON data in HTTP responses

---

## ✨ Real-World Use Cases

* Internal REST APIs
* Admin tools for managing users or records
* Building the backend of full-stack web apps
* CLI tools that manipulate PostgreSQL data

---

## 🔎 Additional Notes

* Your handlers are stateless and testable thanks to clean parameter injection.
* `sql.DB` is a connection pool — you don’t need to open/close it on every query.
* The `cmd/setup` folder is a standalone CLI tool to initialize schema (clean pattern).

---

## 🔁 What’s Next?

## 🔁 What’s Next?

Lesson 27: **Sessions in Go**  
Learn how to manage user session state using cookies and server-side storage. You'll set up login/logout placeholders, persist data like user IDs across requests, and lay the foundation for secure authentication in the next lesson.

---
```
