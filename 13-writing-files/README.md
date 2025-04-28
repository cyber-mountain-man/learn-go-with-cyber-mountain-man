# 13 - Writing Files in Go

In Go, working with files is simple, safe, and powerful using the `os` package. In this lesson, you'll learn how to create a new file, write text to it, and append additional content — all with proper error handling.

---

## 🎯 Objectives

- Create a new file (or overwrite an existing one)
- Write text data into a file
- Append new content without deleting existing content
- Handle file errors safely and correctly

---

## 📦 Project Structure

```
learn-go-with-cyber-mountain-man/
└── 13-writing-files/
    └── main.go
```

---

## 🔧 Covered Functions and Concepts

| Function                 | Purpose |
|---------------------------|---------|
| `os.Create()`             | Create a new file (or truncate existing one) |
| `file.WriteString()`      | Write string content into the file |
| `os.OpenFile()`           | Open a file with specific flags (e.g., append, write-only) |
| `defer file.Close()`      | Automatically close the file when done (resource safety) |

---

## 🧪 Sample Output

```
--- Creating a file ---
Wrote 64 bytes to testfile.txt

--- Appending to file ---
Successfully appended to testfile.txt
```

---

## 🧠 Key Concepts

| Concept                          | Description |
|----------------------------------|-------------|
| File Creation                    | `os.Create()` generates a new writable file. |
| Safe Writing                     | Always handle errors and close files using `defer`. |
| Appending Without Overwriting    | `os.OpenFile()` with `O_APPEND | O_WRONLY` lets you add text at the end. |
| Platform Independence            | Go handles file paths and permissions correctly across operating systems. |

---

## 📄 Notes

- **Error Handling:** Every file operation is immediately checked for errors.
- **Safe Closing:** Using `defer` ensures that files are closed even if an error occurs partway through.
- **Permissions:** When appending, we set permission `0644`, which is common for text files (`rw-r--r--`).

---

🔁 Up Next:  
Lesson `14-reading-files` — learn how to read back data from a file and process its contents line-by-line or all at once.
```