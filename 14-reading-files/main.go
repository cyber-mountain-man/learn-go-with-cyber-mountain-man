package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// -----------------------------
	// 1️⃣ READ ENTIRE FILE AT ONCE
	// -----------------------------
	fmt.Println("--- Reading entire file at once ---")

	// os.ReadFile reads the entire contents of a file and returns it as a byte slice
	content, err := os.ReadFile("testfile.txt")
	if err != nil {
		// If the file doesn't exist or can't be read, print the error and exit
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert byte slice to string for printing
	fmt.Println(string(content))

	// -----------------------------
	// 2️⃣ READ FILE LINE BY LINE
	// -----------------------------
	fmt.Println("\n--- Reading file line by line ---")

	// Open the file in read-only mode
	file, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when function exits

	// Create a scanner to read the file line-by-line
	scanner := bufio.NewScanner(file)

	lineNumber := 1
	for scanner.Scan() {
		// Print each line with its line number
		fmt.Printf("Line %d: %s\n", lineNumber, scanner.Text())
		lineNumber++
	}

	// Check for scanner errors (EOF is not considered an error)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
	}
}

/*
🧠 FILE READING IN GO — WALKTHROUGH

✅ Full File Read (os.ReadFile)
- Best for small/medium-sized files
- Reads the entire file into memory
- Returns []byte, which you can convert to a string
- Simple, fast, great for quick reads

✅ Line-by-Line Reading (bufio.Scanner)
- Best for large files or when you need to process line-by-line
- Efficient memory usage (no need to load entire file)
- `scanner.Text()` gives the current line as a string

✅ Safe Resource Management
- `defer file.Close()` ensures that open file handles are always closed safely
- Helps avoid resource leaks and accidental file locks

✅ Error Handling
- All file operations are wrapped with proper error checking
- `scanner.Err()` must be checked after `Scan()` loop finishes

📦 Key Packages:
- `os` → general file handling and reading entire files
- `bufio` → efficient buffered reading (especially for lines)

This lesson prepares you to build tools like log parsers, config readers, file searchers, and data pipelines.
*/
