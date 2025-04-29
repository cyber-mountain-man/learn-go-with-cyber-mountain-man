# 16 - Channels in Go

Channels are Go's built-in concurrency-safe communication mechanism.  
They allow **goroutines to synchronize and exchange data** without explicit locks or race conditions.

In this lesson, you'll create a channel, send a value from a goroutine, and receive it in the main function.

---

## 🎯 Objectives

- Create and use a channel
- Send values into a channel
- Receive values from a channel (blocking)
- Understand how channels naturally synchronize goroutines

---

## 📦 Project Structure

```
learn-go-with-cyber-mountain-man/
└── 16-channels/
    └── main.go
```

---

## 🔧 Covered Functions and Concepts

| Concept             | Description |
|---------------------|-------------|
| `make(chan type)`    | Create a new unbuffered channel |
| `go function()`      | Launch a function as a goroutine |
| `channel <- value`   | Send data into the channel (blocking if no receiver) |
| `value := <-channel` | Receive data from the channel (blocking until a value is available) |

---

## 🧪 Sample Output

```
Waiting to receive...
Sending message...
Message sent!
Received: Hello from goroutine!
Main function finished!
```

> 🧠 Notice:  
> The main program **waits** ("blocks") at `msg := <-ch` until the goroutine sends its message.  
> Channels **safely synchronize** concurrent work without needing manual locks!

---

## 🧠 Key Concepts

| Concept                 | Explanation |
|--------------------------|-------------|
| Blocking on Send/Receive | Both sending and receiving block until the other side is ready |
| Goroutine Communication  | Channels allow concurrent parts to pass data safely |
| Natural Synchronization  | Channels remove the need for mutexes or shared memory locking |

---

## 📄 Notes

- Channels are **synchronous by default**: no buffering unless specified.
- If you try to receive from an empty channel, the goroutine **blocks** until something is sent.
- If you try to send to a full (unbuffered) channel without a receiver ready, the sender **blocks**.
- Channels and goroutines together form the foundation of Go's elegant concurrency model.

---

🔁 Up Next:  
Lesson `17-web-servers` — build your first simple web server using Go's `net/http` package!
```
