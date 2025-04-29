# 15 - Goroutines in Go

Goroutines are Go's lightweight concurrency primitives that allow you to run functions independently, without blocking the main program. They are incredibly cheap compared to traditional threads — Go can easily handle thousands of goroutines at once.

In this lesson, you'll learn how to launch basic goroutines and observe concurrent behavior.

---

## 🎯 Objectives

- Understand what a goroutine is
- Launch functions asynchronously using the `go` keyword
- Observe how concurrent execution differs from synchronous execution
- Use `time.Sleep()` temporarily to allow goroutines to finish

---

## 📦 Project Structure

```
learn-go-with-cyber-mountain-man/
└── 15-go-routines/
    └── main.go
```

---

## 🔧 Covered Concepts and Functions

| Concept                         | Description |
|----------------------------------|-------------|
| `go funcName(args)`              | Launch a function as a goroutine (concurrent thread) |
| `time.Sleep(duration)`           | Pause the main function temporarily |
| Synchronous vs. Asynchronous     | Main thread waits vs. launches and continues immediately |

---

## 🧪 Sample Output

```
--- Normal Function Call ---
[Synchronous] Message #1
[Synchronous] Message #2
[Synchronous] Message #3

--- Goroutine Function Call ---
Main continues while goroutines run...
[Asynchronous 1] Message #1
[Asynchronous 2] Message #1
[Asynchronous 1] Message #2
[Asynchronous 2] Message #2
[Asynchronous 1] Message #3
[Asynchronous 2] Message #3

Main function finished!
```

> 🚀 Notice how the asynchronous messages from both goroutines are **interleaved**, proving they are running concurrently!

---

## 🧠 Key Concepts

| Concept                     | Explanation |
|------------------------------|-------------|
| Goroutine                    | A lightweight managed thread launched with `go` |
| Non-blocking execution       | `main()` continues while goroutines are running |
| `time.Sleep()` (temporary)   | Keeps `main()` alive long enough for goroutines to complete |
| Later Improvements           | Use `sync.WaitGroup` or channels instead of sleeping |

---

## 📄 Notes

- Goroutines are extremely lightweight — you can run thousands easily.
- If the main program exits, unfinished goroutines are immediately killed.
- `time.Sleep()` is a hacky way to allow time; we'll learn better coordination soon (WaitGroups, Channels).

---

🔁 Up Next:  
Lesson `16-channels` — learn how to **communicate safely** between goroutines using Go's most powerful concurrency tool: **Channels**!
```