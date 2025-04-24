package main

import "fmt"

// doubleValue accepts an int by value (a copy).
// Modifying 'x' inside this function does NOT affect the original.
func doubleValue(x int) {
	x *= 2
	fmt.Println("Inside doubleValue:", x)
}

// doublePointer accepts a pointer to an int.
// Dereferencing it allows direct modification of the original value in memory.
func doublePointer(x *int) {
	*x *= 2
	fmt.Println("Inside doublePointer:", *x)
}

// User is a simple struct to demonstrate how pointer receivers
// can modify struct fields.
type User struct {
	name string
}

// updateName is a method with a pointer receiver (*User).
// It updates the 'name' field of the User struct.
func (u *User) updateName(newName string) {
	u.name = newName
}

func main() {
	// -----------------------------
	// 1️⃣ BASIC POINTERS
	// -----------------------------
	fmt.Println("--- Basic Pointers ---")

	age := 30               // Declare an int variable
	agePtr := &age          // Create a pointer to 'age'

	// Print the original value, the pointer, and the dereferenced value
	fmt.Println("Age:", age)
	fmt.Println("Pointer to age:", agePtr)
	fmt.Println("Value via pointer:", *agePtr) // dereference to get value

	// Modify the value through the pointer
	*agePtr = 35
	fmt.Println("Modified age via pointer:", age)

	// -----------------------------
	// 2️⃣ POINTERS IN FUNCTIONS
	// -----------------------------
	fmt.Println("\n--- Pointers in Functions ---")

	num := 10
	doubleValue(num)               // Passes a copy; original remains unchanged
	fmt.Println("After doubleValue:", num)

	doublePointer(&num)            // Passes memory address; original is modified
	fmt.Println("After doublePointer:", num)

	// -----------------------------
	// 3️⃣ POINTERS WITH STRUCTS
	// -----------------------------
	fmt.Println("\n--- Pointers with Structs ---")

	user := User{name: "Peach"}    // Create a User struct
	fmt.Println("Before update:", user.name)

	user.updateName("Daisy")       // Pointer receiver modifies the struct field
	fmt.Println("After update:", user.name)

	// -----------------------------
	// 4️⃣ NIL POINTERS
	// -----------------------------
	fmt.Println("\n--- Nil Pointers ---")

	var ptr *int // Declared but not initialized; defaults to nil

	// Safe nil check (this will always be true at this point)
	if ptr == nil {
		fmt.Println("Pointer is nil")
	}

	// Allocate memory using 'new', which returns a pointer to a zero-valued int
	ptr = new(int)
	*ptr = 50
	fmt.Println("Value in newly allocated pointer:", *ptr)
}

/*
🧠 POINTERS IN GO

This lesson demonstrates how to use pointers for:

✅ Reading and writing values via memory addresses
✅ Modifying variables and structs from functions
✅ Creating methods that mutate original data
✅ Allocating and checking nil pointers safely

Understanding pointers is key to working efficiently with structs,
slices, functions, and performance-critical code in Go.
*/
