package main

import (
	"fmt"
	"time"
)

// sendMessage simulates a worker goroutine that sends a message to a channel
func sendMessage(ch chan string) {
	fmt.Println("Sending message...")
	time.Sleep(1 * time.Second) // Simulate work with a 1-second delay

	// Send a string value into the channel
	ch <- "Hello from goroutine!"
	fmt.Println("Message sent!")
}

func main() {
	// -----------------------------
	// 1️⃣ CREATE A CHANNEL
	// -----------------------------
	// Create a channel that can transport string values
	ch := make(chan string)

	// -----------------------------
	// 2️⃣ LAUNCH GOROUTINE TO SEND MESSAGE
	// -----------------------------
	// Start a new goroutine that will send a message after some work
	go sendMessage(ch)

	// -----------------------------
	// 3️⃣ RECEIVE MESSAGE
	// -----------------------------
	// <- blocks execution until a value is sent into the channel
	fmt.Println("Waiting to receive...")
	msg := <-ch // Receive the message from the channel (blocking)
	fmt.Println("Received:", msg)

	fmt.Println("Main function finished!")
}

/*
🧠 CHANNELS IN GO — WALKTHROUGH

✅ What is a Channel?
- A **channel** is a pipeline used by goroutines to safely communicate and share data.
- It ensures thread-safety and eliminates the need for manual locks.

✅ How This Program Works:
1. A **channel** is created using `make(chan string)`.
2. A **goroutine** runs `sendMessage()`, simulating some work with a delay.
3. Once ready, the goroutine sends a message into the channel: `ch <- "Hello from goroutine!"`.
4. Meanwhile, the **main goroutine** waits for a message with `msg := <-ch`, naturally **blocking** until data arrives.
5. Once received, the message is printed, and the program gracefully exits.

✅ Key Properties:
- **Send (`ch <- value`)** blocks until a receiver is ready.
- **Receive (`value := <-ch`)** blocks until a sender sends data.
- Channels allow **safe communication** between concurrent parts of your program.
- No race conditions, no manual mutexes needed!

✅ Real-World Use Cases:
- Worker pools
- Streaming results
- Managing concurrent tasks
- Building efficient and safe backend systems

🚀 Mastering channels is the key to mastering Go's elegant concurrency model!
*/
