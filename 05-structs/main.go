package main

import (
    "encoding/json" // For encoding structs into JSON
    "fmt"           // For printing to console
)

// -----------------------------
// STRUCT DEFINITION
// -----------------------------

// User defines a custom struct type with three fields.
// Each field has a specific type and starts with lowercase (unexported outside this package).
type User struct {
    name     string
    age      int
    verified bool
}

// greet() is a method with a value receiver.
// It can access struct data but cannot modify the original struct.
func (u User) greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", u.name, u.age)
}

// verifyAccount() is a method with a pointer receiver.
// It modifies the original struct by updating the verified field.
func (u *User) verifyAccount() {
    u.verified = true
    fmt.Println(u.name, "has been verified.")
}

func main() {
    // -------------------------------------
    // BASIC STRUCT CREATION
    // -------------------------------------

    // Declare and initialize a struct using named fields
    user1 := User{name: "Mario", age: 30, verified: false}

    // Use positional values (must be in exact order)
    user2 := User{"Luigi", 28, false}

    // Declare an empty struct, assign fields later
    var user3 User
    user3.name = "Peach"
    user3.age = 25 // verified will default to false

    // Print struct contents directly
    fmt.Println(user1)
    fmt.Println(user2)
    fmt.Println(user3)

    // -------------------------------------
    // ACCESSING STRUCT FIELDS & METHODS
    // -------------------------------------

    // Access individual fields using dot notation
    fmt.Println("User1 name:", user1.name)
    fmt.Println("User2 age:", user2.age)

    // Call struct methods (value receivers)
    user1.greet()
    user2.greet()
    user3.greet()

    // Call pointer receiver method to modify the struct
    fmt.Println("Verified status (before):", user1.verified)
    user1.verifyAccount()
    fmt.Println("Verified status (after):", user1.verified)

    // ====================================================
    // 1️⃣ STRUCT EMBEDDING
    // ====================================================

    // Admin embeds the User struct, gaining access to all its fields and methods.
    type Admin struct {
        User           // Embedded struct (acts like inheritance)
        accessLevel string
    }

    // Initialize an Admin with embedded User
    admin := Admin{
        User:        User{name: "Toad", age: 35},
        accessLevel: "super",
    }

    fmt.Println("\n--- Struct Embedding ---")
    fmt.Println("Admin name:", admin.name)        // Access field from embedded User
    fmt.Println("Access level:", admin.accessLevel)
    admin.greet()                                 // Call method from embedded User

    // ====================================================
    // 2️⃣ ANONYMOUS STRUCT
    // ====================================================

    // Declare a one-off struct inline without defining a new type
    guest := struct {
        name string
        age  int
    }{
        name: "Daisy",
        age:  22,
    }

    fmt.Println("\n--- Anonymous Struct ---")
    fmt.Printf("Guest: %s (%d years old)\n", guest.name, guest.age)

    // ====================================================
    // 3️⃣ SLICE OF STRUCTS
    // ====================================================

    // Create a slice of User structs
    users := []User{
        {name: "Yoshi", age: 21, verified: true},
        {name: "Rosalina", age: 29, verified: false},
    }

    fmt.Println("\n--- Slice of Structs ---")
    // Loop through the slice and call a method on each user
    for _, u := range users {
        u.greet()
    }

    // ====================================================
    // 4️⃣ JSON TAGS (FOR ENCODING)
    // ====================================================

    // Define a Profile struct with JSON tags
    // Tags tell the JSON encoder to use lowercase keys
    type Profile struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    // Create a Profile instance
    profile := Profile{Name: "Toad", Email: "t@example.com"}

    // Convert struct to JSON
    jsonData, err := json.Marshal(profile)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
    } else {
        fmt.Println("\n--- JSON Output ---")
        fmt.Println(string(jsonData)) // Output: {"name":"Toad","email":"t@example.com"}
    }
}
