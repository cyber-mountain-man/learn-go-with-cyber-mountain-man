# 17 - Web Servers in Go

Go makes it easy to create powerful, production-ready web servers using just the standard library. In this lesson, you’ll build a simple HTTP server, define routes, and return dynamic responses — all with clean and readable code.

---

## 🎯 Objectives

- Understand the structure of an HTTP web server in Go
- Create custom route handlers
- Send responses using `http.ResponseWriter`
- Log requests and errors using the `log` package
- Implement a `/health` check endpoint (production best practice)

---

## 📦 Project Structure

```
learn-go-with-cyber-mountain-man/
└── 17-web-servers/
    └── main.go
```

---

## 🔧 Key Features in This Lesson

| Concept               | Description |
|------------------------|-------------|
| `http.HandleFunc()`    | Maps routes (paths) to handler functions |
| `http.ListenAndServe`  | Starts an HTTP server on the specified port |
| `http.ResponseWriter`  | Used to send data back to the client |
| `*http.Request`        | Provides information about the HTTP request |
| `log.Println()`        | Logs each request to the terminal |
| `log.Fatal()`          | Logs a fatal error and exits the server |
| `/health` endpoint     | Common health check route for uptime monitoring |

---

## 🧪 How to Run

```bash
go run main.go
```

Then visit:

- [http://localhost:8080](http://localhost:8080) → Welcome message  
- [http://localhost:8080/about](http://localhost:8080/about) → Info message  
- [http://localhost:8080/health](http://localhost:8080/health) → `OK` (status 200)

---

## 🔐 Real-World Considerations

- Logging requests helps with monitoring, debugging, and tracing attacks.
- The `/health` route is used by load balancers and orchestrators (e.g., Kubernetes) to determine if the service is up.
- The Go HTTP server handles each request in its **own goroutine** — concurrency is built in!

---

## 🧠 Key Takeaways

- Go makes it simple to get a secure, fast web server running in minutes.
- The `net/http` standard library is sufficient for building REST APIs, serving files, or hosting sites.
- You can expand this server with middleware, templates, routing, and authentication.

---

🔁 Up Next:  
Lesson `18-json` — Learn how to send and receive structured JSON data, build APIs, and parse client requests in Go!
```