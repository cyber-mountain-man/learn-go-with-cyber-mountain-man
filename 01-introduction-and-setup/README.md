# 01 - Introduction and Setup

Welcome to your first step in learning Go! In this section, you'll install Go, set up your workspace, and run your first Go program to make sure everything is working.

---

### 🎯 Objectives

- Install the Go programming language
- Verify your setup
- Run your first Go program

---

### 🛠️ Steps to Get Started

1. **Install Go**  
   Download and install from the official site:  
   👉 [https://go.dev/dl](https://go.dev/dl)

2. **Verify the installation**
   ```bash
   go version
   ```

3. **Run your first Go program**
   ```bash
   go run main.go
   ```

---

### 📄 Example Code

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!!!")
    fmt.Println("Welcome to the introduction of Golang.")
}
```

---

### ✅ Expected Output

```
Hello World!!!
Welcome to the introduction of Golang.
```

---

### 🧠 Key Concepts

- **`package main`**: The entry point for a Go program
- **`import "fmt"`**: Brings in Go's built-in formatting library
- **`func main()`**: The function that runs when the program starts
- **`fmt.Println()`**: Prints output to the terminal

---

🔁 Ready to continue? Move on to the next section:  
👉 `02-your-first-go-file`
```