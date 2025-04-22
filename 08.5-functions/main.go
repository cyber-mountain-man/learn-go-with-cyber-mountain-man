package main

import (
	"fmt"
)

// greet prints a static message.
// It's a basic function with no parameters and no return value.
func greet() {
	fmt.Println("Hello from the greet function!")
}

// add takes two integers and returns their sum.
// Demonstrates a function with parameters and a single return value.
func add(a int, b int) int {
	return a + b
}

// fullName returns two string values.
// Useful for demonstrating multiple return values in Go.
func fullName() (string, string) {
	return "John", "Smith"
}

// triple multiplies the input by 3 and returns it using a named return value.
// The return variable 'result' is declared in the signature itself.
func triple(x int) (result int) {
	result = x * 3
	return // implicit return of 'result'
}

// variadicSum accepts a variable number of integer arguments (nums...)
// and returns the total sum of those numbers.
func variadicSum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	// -----------------------------
	// 1️⃣ BASIC FUNCTION
	// -----------------------------
	fmt.Println("--- Basic function call ---")
	// Calling a function with no input and no return
	greet()

	// -----------------------------
	// 2️⃣ FUNCTION WITH PARAMETERS
	// -----------------------------
	fmt.Println("\n--- Function with parameters and return value ---")
	// Passing two integers to the add function and printing the result
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// -----------------------------
	// 3️⃣ MULTIPLE RETURN VALUES
	// -----------------------------
	fmt.Println("\n--- Function with multiple return values ---")
	// Unpacking the two return values into 'first' and 'last'
	first, last := fullName()
	fmt.Println("Full name:", first, last)

	// -----------------------------
	// 4️⃣ NAMED RETURN VALUE
	// -----------------------------
	fmt.Println("\n--- Named return value ---")
	// Calling a function that returns a named result without explicit return variable
	tripleResult := triple(7)
	fmt.Println("Triple of 7 is:", tripleResult)

	// -----------------------------
	// 5️⃣ VARIADIC FUNCTION
	// -----------------------------
	fmt.Println("\n--- Variadic function ---")
	// Calling a function with a variable number of integer arguments
	sum := variadicSum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", sum)

	// -----------------------------
	// 6️⃣ ANONYMOUS FUNCTION
	// -----------------------------
	fmt.Println("\n--- Anonymous function ---")
	// Defining and immediately calling a function with no name
	func(msg string) {
		fmt.Println("Message:", msg)
	}("Hello Go Learner!")

	// -----------------------------
	// 7️⃣ FUNCTION AS VALUE
	// -----------------------------
	fmt.Println("\n--- Function as value ---")
	// Assigning a function to a variable and calling it
	double := func(n int) int {
		return n * 2
	}
	fmt.Println("Double of 6 is:", double(6))

	// -----------------------------
	// 8️⃣ CLOSURE (Function with internal state)
	// -----------------------------
	fmt.Println("\n--- Closure function ---")
	// A closure remembers the value of 'count' between calls
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}

/*
🧠 FUNCTIONS IN GO

This lesson introduces Go functions, the foundation of code reusability and modular design.
Here’s what we covered:

✅ Declaring and calling basic functions
✅ Using parameters and return values
✅ Returning multiple values
✅ Named return values for clarity
✅ Variadic functions (flexible argument counts)
✅ Anonymous functions (inline definitions)
✅ Storing functions in variables (first-class functions)
✅ Closures (functions with persistent internal state)

Functions are a key part of idiomatic Go. Once mastered, they unlock more advanced topics
like methods, interfaces, and goroutines — all of which are built around function concepts.
*/
