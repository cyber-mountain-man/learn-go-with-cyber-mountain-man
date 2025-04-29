# 18 - Working with JSON in Go

JSON (JavaScript Object Notation) is the universal standard for data interchange between servers and clients.  
In this lesson, you'll learn how to **send** and **receive** JSON using Go's powerful `encoding/json` package.

---

## 🎯 Objectives

- Create Go structs that map to JSON data
- Encode Go structs into JSON responses
- Decode JSON request bodies into Go structs
- Set appropriate Content-Type headers
- Implement basic field validation and method restrictions

---

## 📦 Project Structure

```
learn-go-with-cyber-mountain-man/
└── 18-json/
    └── main.go
```

---

## 🔧 Covered Concepts and Functions

| Concept                         | Description |
|----------------------------------|-------------|
| `json:"fieldname"` struct tags   | Maps Go struct fields to JSON keys |
| `json.NewEncoder(w).Encode(v)`   | Encode a Go struct and send it as JSON |
| `json.NewDecoder(r.Body).Decode(v)` | Parse incoming JSON into a Go struct |
| HTTP Status Codes                | Return `400 Bad Request`, `405 Method Not Allowed`, `500 Internal Server Error` |
| Validation Checks                | Ensure required fields are present |

---

## 🧪 How to Run

Start the server:

```bash
go run main.go
```

---

## 🔥 How to Test Endpoints

### Test `/send` (GET Request)

Open your browser or run:

```bash
curl http://localhost:8080/send
```

✅ Should return a JSON object like:

```json
{
  "name": "Peach Toadstool",
  "email": "peach@example.com",
  "admin": true
}
```

---

### Test `/receive` (POST JSON)

Depending on your environment:

| Environment | Working curl Command |
|:------------|:----------------------|
| **Linux / macOS / Git Bash** | ```curl -X POST http://localhost:8080/receive -H "Content-Type: application/json" -d '{"name":"Mario", "email":"mario@nintendo.com", "admin":false}'``` |
| **Windows CMD (Command Prompt)** | ```curl -X POST http://localhost:8080/receive -H "Content-Type: application/json" -d "{\"name\":\"Mario\",\"email\":\"mario@nintendo.com\",\"admin\":false}"``` |
| **Windows PowerShell** | ```curl -Method POST -Uri http://localhost:8080/receive -Body '{"name":"Mario","email":"mario@nintendo.com","admin":false}' -ContentType "application/json"``` |

✅ Should respond with:

```
Received user: {Name:Mario Email:mario@nintendo.com Admin:false}
```

✅ Logs will show:

```
Parsed user: main.User{Name:"Mario", Email:"mario@nintendo.com", Admin:false}
```

---

## 🧠 Key Takeaways

| Feature                  | Benefit |
|---------------------------|---------|
| Struct Tags (`json:""`)    | Clean mapping between Go fields and JSON |
| Method Restriction (`POST`) | Ensures endpoint security and clarity |
| Field Validation           | Protects against incomplete/bad data |
| Proper Headers             | Ensures clients interpret responses correctly |

---

## 📄 Notes

- Go's JSON encoder/decoder automatically handles common types (`string`, `bool`, `int`, etc.).
- Always validate and sanitize incoming data in real-world applications.
- More complex APIs will involve nested structs, arrays, and custom unmarshalling.

---

🔁 Up Next:  
Lesson `19-project-setup` — start planning and scaffolding your own Go projects!

```