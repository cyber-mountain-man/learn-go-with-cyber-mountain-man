# 20 - Handlers in Go

This lesson introduces how to organize HTTP route logic using clean architectural separation.  
Instead of cramming all handlers into `main.go`, you'll delegate route behavior to a dedicated `handlers` package —  
a structure used by real-world Go web servers.

---

## 🎯 Objectives

- Create custom HTTP handlers for routes like `/`, `/about`, and `/health`
- Separate route logic from server setup using an internal package
- Use `http.NewServeMux()` to explicitly manage routing
- Implement a custom 404 fallback
- Apply basic logging for observability

---

## 📦 Project Structure

```
20-handlers/
├── go.mod
├── cmd/
│   └── app/
│       └── main.go         → App entry: routes + server startup
└── internal/
    └── handlers/
        └── routes.go       → Route logic for pages + /health
```

---

## 🚀 Running the Server

From the root of `20-handlers/`:

```bash
go run ./cmd/app
```

✅ Output in terminal:

```
🚀 Server running at http://localhost:8080
```

Then try:

- [http://localhost:8080](http://localhost:8080)
- [http://localhost:8080/about](http://localhost:8080/about)
- [http://localhost:8080/health](http://localhost:8080/health)
- [http://localhost:8080/not-a-page](http://localhost:8080/not-a-page)

---

## 🔁 Route Behavior

| Route           | Behavior                        |
|------------------|---------------------------------|
| `/`              | Shows home page                 |
| `/about`         | Shows info about the server     |
| `/health`        | Returns "server is healthy" ✅  |
| any unknown path | Redirected to custom `/404`     |

---

## 📜 Key Concepts

| Concept                  | Description |
|---------------------------|-------------|
| `http.NewServeMux()`      | Explicit router for modularity |
| `handlers` package        | Keeps business logic out of main |
| Logging with `log.Println`| Helps with observability |
| `w.WriteHeader(...)`      | Shows how to return HTTP status codes |
| `/health` endpoint        | Useful for load balancers and uptime checks |

---

## 💡 Educational Insights

- This is a professional web server layout, scaled down for learning.
- Splitting handlers into a separate package prepares you for:
  - Middleware
  - Testing
  - Modular APIs
- `ServeMux` avoids global state and supports testing/multiple servers.

---

## 🛡️ Next Steps

- 🧱 Add more routes and group them by domain (e.g., `auth`, `products`)
- 🧪 Add unit tests using `httptest.NewRecorder()`
- 🧰 Integrate graceful shutdown with `context`
- 🌐 Introduce environment-based configuration (`os.Getenv("PORT")`)

---

🔁 Up Next:  
Lesson `21-routing` — dive deeper into routing systems and dynamic URL parameters using third-party libraries like `chi` or `gorilla/mux`.

```