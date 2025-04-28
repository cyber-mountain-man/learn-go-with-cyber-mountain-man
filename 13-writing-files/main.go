package main

import (
	"fmt"
	"os"
)

func main() {
	// -----------------------------
	// 1️⃣ CREATE OR OPEN A FILE
	// -----------------------------
	fmt.Println("--- Creating a file ---")

	// os.Create creates a new file or truncates an existing one
	file, err := os.Create("testfile.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // defer ensures the file is closed at the end of main()

	// -----------------------------
	// 2️⃣ WRITE TO FILE
	// -----------------------------
	content := "Hello, this is the first line!\nWelcome to writing files in Go.\n"

	// WriteString writes the string content into the file
	bytesWritten, err := file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Wrote %d bytes to testfile.txt\n", bytesWritten)

	// -----------------------------
	// 3️⃣ APPEND TO EXISTING FILE
	// -----------------------------
	fmt.Println("\n--- Appending to file ---")

	// OpenFile opens the file in append and write-only mode (O_APPEND | O_WRONLY)
	appendFile, err := os.OpenFile("testfile.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for append:", err)
		return
	}
	defer appendFile.Close() // defer closing the append file as well

	appendContent := "This line is appended at the end.\n"

	// WriteString appends new content without erasing existing data
	_, err = appendFile.WriteString(appendContent)
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("Successfully appended to testfile.txt")
}

/*
🧠 WRITING FILES IN GO — WALKTHROUGH

✅ Step 1: Create a new file (or overwrite if it already exists)
- We use `os.Create()` to make a fresh file ready for writing.
- Always use `defer file.Close()` to ensure you don't leak file descriptors.

✅ Step 2: Write content to the file
- `WriteString()` allows you to easily dump simple text into the file.
- It returns the number of bytes written — useful for validation.

✅ Step 3: Append additional content
- `os.OpenFile()` with `O_APPEND` and `O_WRONLY` flags lets us *add* new text without overwriting existing content.
- Again, use `WriteString()` to insert text at the end.

✅ Error Handling
- Every file operation is checked for errors immediately.
- Always `return` early if an error occurs.

✅ Key Concepts:
- Safe file handling (`defer Close()`)
- Different file modes (create, write, append)
- Managing multiple file handles safely

🔔 Real-world uses: creating logs, saving data, writing configuration files, generating reports, etc.

*/
