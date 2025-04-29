package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User defines the structure for JSON encoding/decoding.
// The struct tags (e.g., `json:"name"`) tell Go how to map struct fields to JSON keys.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Admin bool   `json:"admin"`
}

// sendJSONHandler handles GET requests to "/send" and returns a JSON-encoded User object.
func sendJSONHandler(w http.ResponseWriter, r *http.Request) {
	// Simulated user data to send as JSON
	user := User{
		Name:  "Peach Toadstool",
		Email: "peach@example.com",
		Admin: true,
	}

	// Set the response Content-Type header to indicate JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the user struct as JSON and send it in the response
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		log.Println("Encoding error:", err)
	}
}

// receiveJSONHandler handles POST requests to "/receive" and parses incoming JSON payloads.
func receiveJSONHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow HTTP POST requests to this route
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	// Decode the JSON body into the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Println("Decoding error:", err)
		return
	}

	// Optional: validate that required fields are present
	if user.Name == "" || user.Email == "" {
		http.Error(w, "Missing required fields: name or email", http.StatusBadRequest)
		return
	}

	// Set the response type for clean formatting
	w.Header().Set("Content-Type", "text/plain")

	// Send confirmation back to the client
	fmt.Fprintf(w, "Received user: %+v\n", user)

	// Log the received struct with field names and values
	log.Printf("Parsed user: %#v", user)
}

func main() {
	// Register route handlers for sending and receiving JSON
	http.HandleFunc("/send", sendJSONHandler)
	http.HandleFunc("/receive", receiveJSONHandler)

	// Start server on port 8080
	port := ":8080"
	log.Println("📦 JSON server running at http://localhost" + port)
	log.Println("➡ Try POSTING JSON to http://localhost" + port + "/receive")

	// Start the web server and handle fatal errors
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

/*
🧠 JSON IN GO — WALKTHROUGH

✅ What This Lesson Covers:
- Creating a struct with JSON tags for mapping fields
- Encoding Go structs as JSON responses
- Decoding JSON request bodies into Go structs
- Using HTTP status codes and headers properly
- Validating received data and restricting HTTP methods

✅ Real-World Use Cases:
- `/send` simulates an API endpoint returning JSON (e.g., to a frontend app)
- `/receive` simulates an API receiving JSON (e.g., from a form or client SDK)

✅ Key Concepts:
| Concept                 | Explanation |
|--------------------------|-------------|
| `json:"name"`            | Maps the field to a JSON key |
| `json.NewEncoder(w)`     | Sends JSON in the HTTP response |
| `json.NewDecoder(r.Body)`| Parses JSON from the request body |
| HTTP method restriction  | Ensures endpoint security and clarity |
| Field validation         | Prevents broken or incomplete data |

🔐 Bonus Security Tip:
Always validate and sanitize inputs when accepting JSON. Never assume client data is safe.

📡 Up Next:
You’ll apply these concepts in REST APIs and microservices that receive and respond with structured JSON payloads.
*/
