package main

import (
	"fmt"
	"time"
)

// printMessage prints a given message three times with a delay between each print.
// This simulates "work" being done over time.
func printMessage(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("[%s] Message #%d\n", msg, i)
		time.Sleep(500 * time.Millisecond) // Pause for 500 milliseconds to simulate a delay
	}
}

func main() {
	// -----------------------------
	// 1️⃣ NORMAL FUNCTION CALL
	// -----------------------------
	fmt.Println("--- Normal Function Call ---")

	// Synchronous execution: main() waits for printMessage() to finish completely
	printMessage("Synchronous")

	// -----------------------------
	// 2️⃣ GOROUTINE FUNCTION CALL
	// -----------------------------
	fmt.Println("\n--- Goroutine Function Call ---")

	// Launch two goroutines (lightweight concurrent threads)
	go printMessage("Asynchronous 1")
	go printMessage("Asynchronous 2")

	// Main function immediately continues without waiting
	fmt.Println("Main continues while goroutines run...")

	// Temporary solution: Sleep to allow goroutines time to complete
	// (Better practices later involve sync.WaitGroup or channels)
	time.Sleep(2 * time.Second)

	fmt.Println("\nMain function finished!")
}

/*
🧠 GOROUTINES IN GO — WALKTHROUGH

✅ Normal Function Call (Synchronous)
- `printMessage("Synchronous")` runs on the main thread.
- The main program **waits** until it fully finishes before moving on.
- Classic, blocking behavior.

✅ Goroutine Function Calls (Asynchronous)
- `go printMessage("Asynchronous 1")` launches the function **in a separate thread of execution**.
- `go printMessage("Asynchronous 2")` launches another.
- The `main()` function does **NOT** wait — it continues immediately.
- Goroutines run concurrently and overlap their outputs.

✅ Temporary Sleep
- Because the `main()` function exits immediately if nothing is blocking it,
- `time.Sleep()` gives the goroutines time to complete.
- In professional code, you'd use `sync.WaitGroup` or channels to coordinate properly.

✅ Key Concepts
- **Concurrency** made extremely lightweight and simple in Go.
- **Goroutines** are cheaper than traditional threads.
- Mastering goroutines unlocks true parallelism, server backends, worker pools, schedulers, and more.

⚡ Up Next:
- Learn `sync.WaitGroup` to properly wait for goroutines.
- Dive into **Channels** for powerful communication between goroutines.
*/
