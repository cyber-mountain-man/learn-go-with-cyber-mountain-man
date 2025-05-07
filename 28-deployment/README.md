# HTMX + Go + MSSQL User Management App

A full-stack web application demonstrating dynamic form submission and rendering using HTMX, Go, and Microsoft SQL Server — all containerized with Docker.

---

## 🚀 Features

- Add, edit, and delete users dynamically via HTMX
- Backend powered by Go with Chi router and `database/sql`
- Microsoft SQL Server as persistent data storage
- Dockerized setup for easy development and deployment
- HTML5 frontend using HTMX 2.0+ and minimal CSS
- Safe HTML swapping with wrapper targets to avoid `htmx:targetError`

---

## ⚙️ Quick Start (PowerShell)

```powershell
# Clone the full repository and navigate to the deployment project
git clone https://github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man.git
cd learn-go-with-cyber-mountain-man/28-deployment

# Start containers and build the Go app
docker-compose up --build
````

Then open [http://localhost:8080](http://localhost:8080) in your browser.

---

## 🧪 Inspect the Database (PowerShell)

Use this command to open a `sqlcmd` prompt from a temporary SQL Tools container:

```powershell
docker run -it --rm --network=28-deployment_go-net mcr.microsoft.com/mssql-tools `
  /bin/bash -c "/opt/mssql-tools/bin/sqlcmd -S mssql -U sa -P YourStrong@Passw0rd"
```

Then inside `sqlcmd`, run:

```sql
USE lesson28_mssql;
SELECT * FROM users;
GO
```

---

## 🐳 Containerization Details

This project is fully containerized using Docker:

* **Go App**:

  * Built with a multi-stage Dockerfile
  * First stage compiles the Go binary using `golang:1.24`
  * Second stage uses `gcr.io/distroless/static:nonroot` for a small and secure final image

* **SQL Server**:

  * Runs in a container using `mcr.microsoft.com/mssql/server:2022-latest`
  * Exposes port 1433 for internal DB connections

* **Initialization Container (`mssql-init`)**:

  * Uses `mcr.microsoft.com/mssql-tools` to run `init.sql`
  * Waits for the database to be healthy
  * Creates:

    * Database (`lesson28_mssql`)
    * Login (`lesson28_user`)
    * User table with optional seed data

* **Docker Compose** orchestrates the app, DB, and init scripts on a shared custom network: `28-deployment_go-net`.

---

## 📁 File Structure

```
├── cmd/
│   └── api/
│       └── main.go              # Go app entry point
├── internal/
│   ├── db/
│   │   └── db.go                # DB connection + queries
│   ├── handlers/
│   │   └── user_handler_htmx.go # HTMX-compatible CRUD handlers
│   ├── models/
│   │   └── user.go              # User struct
│   └── router/
│       └── router.go            # Chi router setup
├── static/
│   ├── index.html               # Main HTMX-powered frontend
│   └── templates/
│       └── user-list.html       # Server-side template fragment
├── mssql-init/
│   └── init.sql                 # SQL to create DB, login, schema
├── .env                         # DB connection values for Go app
├── Dockerfile                   # Multi-stage Go + distroless build
├── docker-compose.yml           # Dev environment orchestration
└── README.md                    # This file
```

---

## 🧰 Tech Stack

* **Go** (1.24) with Chi router
* **HTMX** (2.0.4) for frontend interactivity
* **Microsoft SQL Server 2022**
* **Docker Compose**
* **Distroless containers** for secure production builds

---

## 🙌 Credits

* Created by **Guillermo Morrison**
* Designed to demonstrate HTMX + Go + SQL integration in a production-like setup
```
