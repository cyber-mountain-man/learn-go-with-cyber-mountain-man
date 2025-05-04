# 26 - Databases in Go (MSSQL Version)

This lesson demonstrates how to build a REST API in Go backed by **Microsoft SQL Server**. You'll learn how to structure a web application that connects to MSSQL using environment variables, organizes database access logic, and exposes user CRUD endpoints over HTTP.

---

## 🎯 Objectives

- Connect Go to MSSQL using `github.com/denisenkom/go-mssqldb`
- Organize database, model, and handler logic
- Use environment variables to store credentials securely
- Perform full CRUD operations on a `users` table
- Understand how to set up MSSQL for local Go development

---

## 🧱 Folder Structure

```

26-databases/
└── mssql/
├── cmd/
│   └── setup/
│       └── main.go           # Schema creation tool for 'users' table
├── internal/
│   ├── db/                   # MSSQL connection logic (.env-based)
│   ├── handlers/             # HTTP route logic (CRUD)
│   └── models/               # SQL functions for interacting with MSSQL
├── .env                      # Local database credentials (excluded in Git)
├── main.go                   # Entry point for the API server
└── README.md                 # This file

````

---

## 🛠️ MSSQL Setup Instructions

### 1. Start SQL Server

- Open **SQL Server Configuration Manager**
- Make sure `SQL Server (MSSQLSERVER)` is **running**
- Optional: start `SQL Server Browser`

### 2. Enable TCP/IP Access

- In Configuration Manager:  
  → `SQL Server Network Configuration` → `Protocols for MSSQLSERVER`  
  ✅ Enable **TCP/IP**
- Under `TCP/IP` → `Properties` → `IPAll`:
  - Set **TCP Port** = `1433`
  - Clear **TCP Dynamic Ports**
- Restart `SQL Server (MSSQLSERVER)`

### 3. Create Login, Database, and Grant Permissions

Open `sqlcmd` or SSMS with an admin account:

```sql
CREATE LOGIN lesson26_user WITH PASSWORD = 'lesson26_pass';
GO

CREATE DATABASE lesson26_mssql;
GO

USE lesson26_mssql;
GO

CREATE USER lesson26_user FOR LOGIN lesson26_user;
GO

ALTER ROLE db_owner ADD MEMBER lesson26_user;
GO
````

---

## 🔐 .env File

Create a `.env` file in the `mssql/` folder:

```
DBUSER=lesson26_user
DBPASSWORD=lesson26_pass
DBHOST=localhost
DBPORT=1433
DBNAME=lesson26_mssql
```

---

## 🚀 Running the App

### 1. Run schema setup

```bash
go run ./cmd/setup
```

Creates the `users` table (safe to run multiple times).

### 2. Start the API Server

```bash
go run main.go
```

Visit: [http://localhost:8080/users](http://localhost:8080/users)

---

## 📬 Test with curl

### 🟦 CMD (Windows):

```cmd
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d "{\"name\":\"Sam\",\"email\":\"sam@example.com\"}"
```

### ⚙️ PowerShell:

```powershell
curl -Method POST http://localhost:8080/users `
  -Headers @{"Content-Type"="application/json"} `
  -Body '{"name":"Sam","email":"sam@example.com"}'
```

### 🐧 macOS/Linux:

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Sam","email":"sam@example.com"}'
```

---

## 📦 Endpoints

| Method | Route         | Description       |
| ------ | ------------- | ----------------- |
| GET    | `/users`      | List all users    |
| POST   | `/users`      | Create a new user |
| GET    | `/users/{id}` | Get user by ID    |
| PUT    | `/users/{id}` | Update user by ID |
| DELETE | `/users/{id}` | Delete user by ID |

---

## 🔁 What’s Next?

Lesson 27: **Sessions in Go**
Learn how to manage sessions, secure your API routes, and implement login/logout features.

---
```
