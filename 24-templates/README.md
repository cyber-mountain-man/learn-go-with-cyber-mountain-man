# 24 - Templates in Go

This lesson introduces Go’s `html/template` package to build dynamic, structured, and reusable HTML pages. You’ll learn how to create a layout-based system and render page-specific content while passing dynamic data from your Go backend.

---

## 🎯 Objectives

- Use Go’s built-in `html/template` package
- Create a reusable layout (`layout.html`) and page-specific view (`home.html`)
- Inject shared data (title, user, year) into HTML
- Serve both templated and plain-text routes
- Build a minimal yet professional frontend using only Go

---

## 📁 Project Structure

```

learn-go-with-cyber-mountain-man/
└── 24-templates/
├── main.go                      # Entry point (sets up routes and server)
├── internal/
│   ├── handlers/
│   │   ├── public.go           # Handles public/private routes
│   │   └── routes.go           # Template-rendering logic
│   └── templates/
│       ├── layout.html         # Shared layout with header/footer
│       └── home.html           # Content injected into layout

````

---

## 🚀 Run the App

### 1. Open your terminal and navigate to the lesson:

```bash
cd 24-templates
````

### 2. Start the web server:

```bash
go run main.go
```

### 3. Visit these endpoints in your browser:

* `http://localhost:8080/` → renders the dynamic home page using templates
* `http://localhost:8080/public` → plain public page
* `http://localhost:8080/private` → plain private page (no auth yet)

---

## 🧠 Concepts in Use

| Feature                    | Purpose                                        |
| -------------------------- | ---------------------------------------------- |
| `html/template`            | Safely render HTML and dynamic values          |
| `{{define "layout"}}`      | Declares the shared layout structure           |
| `{{template "content" .}}` | Injects view-specific content                  |
| `map[string]interface{}`   | Used to pass multiple variables into templates |
| `http.HandleFunc()`        | Associates URLs with handler functions         |

---

## 🧠 What You Learn

✅ How to serve HTML pages using Go’s standard library
✅ How to structure templates into a layout/content system
✅ How to pass data from Go to your HTML
✅ How to use dynamic values in your views (`.Title`, `.User`, `.Year`)
✅ How to separate plain routes and templated routes cleanly

---

## ✨ Real-World Use Cases

* Dashboards & Admin Panels
* Server-rendered UIs
* Static site generators
* Internal tools or portals

---

## 🔎 Additional Notes

* Your layout file (`layout.html`) defines the look and feel: header, main content block, footer.
* The content file (`home.html`) is injected where `{{template "content" .}}` appears inside the layout.
* Both templates are parsed together using `template.ParseFiles()` in `routes.go`.

---

## 🔁 What's Next?

Lesson 25: **Forms in Go**
Learn how to handle user input, form submissions, and process data in Go.

---
```

