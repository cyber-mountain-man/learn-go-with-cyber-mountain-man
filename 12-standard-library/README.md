# 12 - Standard Library in Go

Go comes with a powerful **standard library** — a rich set of built-in packages that handle strings, math, dates, OS interactions, and much more. No external downloads required!

In this lesson, you'll explore some of the most essential standard packages.

---

## 🎯 Objectives

- Learn to manipulate text with the `strings` package
- Perform mathematical operations with the `math` package
- Work with dates and times using the `time` package
- Interact with the operating system using the `os` package

---

## 📦 Project Structure

```
learn-go-with-cyber-mountain-man/
└── 12-standard-library/
    └── main.go
```

---

## 🔧 Covered Packages

### 🔹 `strings`

- `strings.ToUpper()`
- `strings.ToLower()`
- `strings.HasPrefix()`
- `strings.Count()`

### 🔹 `math`

- `math.Sqrt()`
- `math.Pow()`
- `math.Pi`

### 🔹 `time`

- `time.Now()`
- `time.Add()`
- `time.Format()`

### 🔹 `os`

- `os.Getenv()`
- `os.Getwd()`

---

## 🧪 Sample Output

```
--- strings package ---
Original text: hello, world
Upper: HELLO, WORLD
Lower: hello world
Has prefix 'he': true
Count of 'l' in text: 3

--- math package ---
Square root of 9: 3
Pi constant: 3.141592653589793
Power 2^3: 8

--- time package ---
Current time: 2025-04-28 14:30:00.1234567 -0500 CDT m=+0.000123456
Two hours later: 2025-04-28 16:30:00.1234567 -0500 CDT m=+7200.000123456
Formatted time: Mon, 28 Apr 2025 14:30:00 CDT

--- os package ---
Logged in user: Guillermo Morrison
Current working directory: C:\Users\Guillermo Morrison\Golang-Projects\learn-go-with-cyber-mountain-man
```

---

## 🧠 Key Concepts

| Concept                     | Description |
|-----------------------------|-------------|
| `strings`                   | Provides utilities for working with text |
| `math`                      | Handles basic math functions and constants |
| `time`                      | Provides date/time manipulation and formatting |
| `os`                        | Allows access to environment variables and file system info |

---

🔁 Up Next:  
Lesson `13-writing-files` — learn how to create and write data into files programmatically using Go!
```