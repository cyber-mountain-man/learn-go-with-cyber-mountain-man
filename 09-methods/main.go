package main

import (
	"errors"
	"fmt"
)

// -----------------------------
// 1️⃣ BASE STRUCT
// -----------------------------

// User represents a simple user account with basic fields.
// This will serve as the main data structure for method demonstrations.
type User struct {
	name     string
	age      int
	verified bool
}

// -----------------------------
// 2️⃣ VALUE RECEIVER (READ-ONLY)
// -----------------------------

// greet is a method with a value receiver.
// It uses a copy of the User struct and prints a greeting.
// Since it's read-only, it doesn't affect the actual User.
func (u User) greet() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", u.name, u.age)
}

// isAdult checks if the user is at least 18 years old.
// Uses a value receiver since it doesn't modify the struct.
func (u User) isAdult() bool {
	return u.age >= 18
}

// updateAgeNoEffect demonstrates what happens when you try
// to modify a struct using a value receiver — it has no effect.
// We assign to u.age and then use it to suppress a compiler warning.
func (u User) updateAgeNoEffect(newAge int) {
	u.age = newAge
	_ = u.age // avoids "unused write" warning
}

// -----------------------------
// 3️⃣ POINTER RECEIVER (MODIFIES ORIGINAL)
// -----------------------------

// verifyAccount sets the verified field to true.
// This modifies the actual struct by using a pointer receiver.
func (u *User) verifyAccount() {
	u.verified = true
}

// setName updates the user's name and returns a pointer to allow chaining.
func (u *User) setName(name string) *User {
	u.name = name
	return u
}

// setAge updates the user's age and returns a pointer to allow chaining.
func (u *User) setAge(age int) *User {
	u.age = age
	return u
}

// setAgeSafe validates the age before updating it.
// Returns an error if the value is out of bounds.
func (u *User) setAgeSafe(age int) error {
	if age < 0 || age > 130 {
		return errors.New("invalid age: must be between 0 and 130")
	}
	u.age = age
	return nil
}

// -----------------------------
// 4️⃣ EMBEDDED STRUCT EXAMPLE
// -----------------------------

// Admin embeds User and inherits all User methods.
// This shows how Go uses composition instead of inheritance.
type Admin struct {
	User       // Embedded field
	privileged bool
}

// accessLevel returns the access level based on the Admin's privilege.
func (a Admin) accessLevel() string {
	if a.privileged {
		return "super admin"
	}
	return "limited admin"
}

// -----------------------------
// 5️⃣ MAIN FUNCTION
// -----------------------------

func main() {
	// Creating users using pointers to allow use of pointer receivers
	user1 := &User{name: "Mario", age: 25}
	user2 := &User{name: "Peach", age: 17}

	// Call greet method (value receiver)
	fmt.Println("--- Greeting Users ---")
	user1.greet()
	user2.greet()

	// Check age with isAdult method
	fmt.Println("\n--- Age Check ---")
	fmt.Printf("%s adult? %v\n", user1.name, user1.isAdult())
	fmt.Printf("%s adult? %v\n", user2.name, user2.isAdult())

	// Demonstrate that value receivers do not mutate original data
	fmt.Println("\n--- Value vs Pointer Receiver ---")
	user1.updateAgeNoEffect(99)
	fmt.Println("Age after updateAgeNoEffect:", user1.age) // Should still be 25

	// Use pointer method to change data
	user1.setAge(99)
	fmt.Println("Age after setAge:", user1.age) // Should now be 99

	// Use chained pointer methods
	fmt.Println("\n--- Method Chaining ---")
	user2.setName("Princess").setAge(20).verifyAccount()
	user2.greet()
	fmt.Println("Verified:", user2.verified)

	// Demonstrate safe setters that return errors
	fmt.Println("\n--- Safe Setter with Error Handling ---")
	err := user2.setAgeSafe(-5)
	if err != nil {
		fmt.Println("Failed to set age:", err)
	}

	err = user2.setAgeSafe(30)
	if err == nil {
		fmt.Println("Age successfully updated to:", user2.age)
	}

	// Show that embedded structs inherit methods
	fmt.Println("\n--- Embedded Struct with Methods ---")
	admin := Admin{
		User:       User{name: "Luigi", age: 40, verified: true},
		privileged: true,
	}
	admin.greet() // Calls method from embedded User
	fmt.Println("Admin access level:", admin.accessLevel())
}

/*
🧠 METHODS IN GO

This lesson demonstrates Go's method system, including:

✅ Value receivers (read-only methods)
✅ Pointer receivers (state-modifying methods)
✅ Method chaining with pointer returns
✅ Safe setters with error handling
✅ Composition with embedded structs
✅ Inherited methods through embedding

Go does not use traditional OOP classes, but by attaching methods to structs,
you get the benefits of encapsulation, logical grouping, and clean interfaces
while still writing simple, idiomatic Go code.
*/
