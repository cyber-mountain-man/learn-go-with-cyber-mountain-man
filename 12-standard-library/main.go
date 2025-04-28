package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	// -----------------------------
	// 1️⃣ STRINGS PACKAGE
	// -----------------------------
	fmt.Println("--- strings package ---")

	text := "hello, world"
	upperText := strings.ToUpper(text)
	lowerText := strings.ToLower("HELLO WORLD")
	hasPrefix := strings.HasPrefix(text, "he")
	count := strings.Count(text, "l")

	fmt.Println("Original text:", text)
	fmt.Println("Upper:", upperText)
	fmt.Println("Lower:", lowerText)
	fmt.Println("Has prefix 'he':", hasPrefix)
	fmt.Println("Count of 'l' in text:", count)

	// -----------------------------
	// 2️⃣ MATH PACKAGE
	// -----------------------------
	fmt.Println("\n--- math package ---")

	number := 9.0
	fmt.Println("Square root of 9:", math.Sqrt(number))
	fmt.Println("Pi constant:", math.Pi)
	fmt.Println("Power 2^3:", math.Pow(2, 3))

	// -----------------------------
	// 3️⃣ TIME PACKAGE
	// -----------------------------
	fmt.Println("\n--- time package ---")

	now := time.Now()
	fmt.Println("Current time:", now)

	later := now.Add(2 * time.Hour)
	fmt.Println("Two hours later:", later)

	fmt.Println("Formatted time:", now.Format(time.RFC1123))

	// -----------------------------
	// 4️⃣ OS PACKAGE
	// -----------------------------
	fmt.Println("\n--- os package ---")

	user := os.Getenv("USERNAME")
	if user == "" {
		user = os.Getenv("USER") // Fallback for Linux/macOS
	}
	fmt.Println("Logged in user:", user)

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		// Optional: os.Exit(1) to terminate on critical error
	} else {
		fmt.Println("Current working directory:", workingDir)
	}
}

/*
🧠 STANDARD LIBRARY IN GO

This lesson demonstrates powerful, built-in Go packages:

✅ strings — manipulating text easily
✅ math — common math operations
✅ time — managing timestamps and formatting
✅ os — interacting with the operating system (env variables, file paths)
*/
