# 14 - Reading Files in Go

Reading data from files is a core skill in backend development, data processing, and tool building. In this lesson, you'll learn two ways to read files in Go: all at once, and line-by-line — with proper error handling and resource management.

---

## 🎯 Objectives

- Read the entire content of a file into memory
- Read a file line-by-line efficiently
- Handle file I/O errors properly
- Manage file resources safely with `defer`

---

## 📦 Project Structure

```
learn-go-with-cyber-mountain-man/
└── 14-reading-files/
    └── main.go
```

---

## 🔧 Covered Functions and Concepts

| Function / Concept        | Purpose |
|----------------------------|---------|
| `os.ReadFile()`            | Read the entire file into a byte slice |
| `os.Open()`                | Open a file handle manually |
| `bufio.NewScanner()`       | Read a file line-by-line efficiently |
| `defer file.Close()`       | Ensure file handles are properly closed |
| `scanner.Err()`            | Detect and handle any scan errors |

---

## 🧪 Sample Output

```
--- Reading entire file at once ---
Hello, this is the first line!
Welcome to writing files in Go.
This line is appended at the end.

--- Reading file line by line ---
Line 1: Hello, this is the first line!
Line 2: Welcome to writing files in Go.
Line 3: This line is appended at the end.
```

---

## 🧠 Key Concepts

| Concept                      | Description |
|-------------------------------|-------------|
| Full file read (small files)  | Use `os.ReadFile()` when you need the whole file at once |
| Streaming line-by-line        | Use `bufio.Scanner` for memory efficiency with large files |
| Error Handling                | Always check errors immediately when reading or scanning |
| Resource Safety               | Always `defer` file closing right after opening |

---

## 📄 Notes

- `os.ReadFile()` replaces the older, deprecated `ioutil.ReadFile()`.
- `bufio.Scanner` reads files line-by-line without consuming large amounts of memory.
- `defer file.Close()` should be placed **immediately after a successful `os.Open()`**.

---

🔁 Up Next:  
Lesson `15-go-routines` — learn how to run lightweight threads (goroutines) to perform concurrent tasks efficiently in Go!
```