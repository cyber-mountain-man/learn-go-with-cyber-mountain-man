# 25 - HTML Forms in Go

In this lesson, you'll learn how to **create**, **render**, and **process** HTML forms using Go's `net/http` and `html/template` packages.  
You'll build a simple feedback form that uses **GET** for rendering and **POST** for submission, with validation and a confirmation screen.

---

## 🎯 Objectives

- Serve an HTML form using Go's standard library
- Handle form submissions with POST routes
- Perform basic input validation
- Render confirmation messages and redirect on error
- Organize form logic cleanly using templates and handlers

---

## 🗂️ Project Structure

```

25-forms/
├── main.go                     # Entry point
├── internal/
│   ├── handlers/
│   │   └── form.go             # Form rendering and processing logic
│   └── templates/
│       ├── layout.html         # Shared base layout
│       └── form.html           # Page-specific content

````

---

## 🚀 Run the Server

Make sure you’re inside the `25-forms/` directory, then:

```bash
go run main.go
````

Open your browser and visit:
[http://localhost:8080](http://localhost:8080) → Redirects to `/form`
[http://localhost:8080/form](http://localhost:8080/form) → Renders the HTML form

---

## 🧪 Test It Out

Submit the form using valid input and see a confirmation message.
Leave a field blank and you'll receive a validation error.

---

## 🔑 Key Features

| Feature                 | Explanation                                                 |
| ----------------------- | ----------------------------------------------------------- |
| `http.HandleFunc()`     | Defines both GET and POST routes (`/form` and `/submit`)    |
| `template.ParseFiles()` | Combines layout + form views for rendering                  |
| `r.FormValue()`         | Extracts values from submitted POST data                    |
| `http.Redirect()`       | Redirects `/` to `/form` cleanly                            |
| Input validation        | Ensures all fields are filled before submission is accepted |

---

## 🧠 Learning Outcomes

* Understand how to wire GET and POST handlers in Go
* Learn how to pass data into templates for clean rendering
* Practice redirecting between routes using `http.Redirect`
* Gain experience in secure, idiomatic form processing in Go

---

## ✅ Example Output

**Form submission success:**

```
✅ Form submitted successfully!
Name: Mario
Email: mario@nintendo.com
Message: Let’s-a go!
```

**Validation error (missing fields):**

```
400 - All fields are required.
```

---

## 💡 Pro Tips

* Always sanitize and validate input server-side, even if your frontend does it too
* You can expand this by saving the input to a file or database
* Using `POST` for handling sensitive or writable operations is a web standard

---

## 🔁 What’s Next?

> Up next: **Lesson 26 - Databases**
> Where you'll learn to persist this submitted data using SQL and Go!
```