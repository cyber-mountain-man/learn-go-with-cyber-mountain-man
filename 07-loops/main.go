package main

import (
    "fmt"
    "time"
)

func main() {
    // -----------------------------
    // 1️⃣ BASIC FOR LOOP
    // -----------------------------
    fmt.Println("--- Basic for loop ---")

    // This loop starts at 1 and runs until i <= 5.
    // It's the Go equivalent of a classic C-style for loop.
    for i := 1; i <= 5; i++ {
        fmt.Println("Iteration:", i)
    }

    // -----------------------------
    // 2️⃣ WHILE-STYLE FOR LOOP
    // -----------------------------
    fmt.Println("\n--- While-style loop ---")

    // Go doesn't have a separate 'while' keyword.
    // You can achieve the same with a for loop using only a condition.
    counter := 0
    for counter < 3 {
        fmt.Println("Counter:", counter)
        counter++
    }

    // -----------------------------
    // 3️⃣ INFINITE LOOP (with break)
    // -----------------------------
    fmt.Println("\n--- Infinite loop with break ---")

    i := 0
    for {
        fmt.Println("Looping... until i == 2")

        // Exit condition
        if i == 2 {
            break // exits the loop when i equals 2
        }

        i++
    }

    // -----------------------------
    // 4️⃣ LOOP WITH CONTINUE
    // -----------------------------
    fmt.Println("\n--- Loop with continue (skip even numbers) ---")

    // The 'continue' keyword skips the rest of the loop body for the current iteration.
    for i := 1; i <= 5; i++ {
        if i%2 == 0 {
            continue // skip even numbers
        }
        fmt.Println("Odd:", i)
    }

    // -----------------------------
    // 5️⃣ LOOP OVER A SLICE
    // -----------------------------
    fmt.Println("\n--- Loop over slice ---")

    // Create a slice of strings
    fruits := []string{"apple", "banana", "cherry"}

    // 'range' returns both index and value when looping over a slice
    for i, fruit := range fruits {
        fmt.Printf("Index %d: %s\n", i, fruit)
    }

    // -----------------------------
    // 6️⃣ LOOP WITH IGNORED INDEX
    // -----------------------------
    fmt.Println("\n--- Loop with ignored index ---")

    // Use _ (underscore) to ignore the index if it's not needed
    for _, fruit := range fruits {
        fmt.Println("Fruit:", fruit)
    }

    // -----------------------------
    // 7️⃣ LOOP OVER A MAP
    // -----------------------------
    fmt.Println("\n--- Loop over map ---")

    // Create a map of names to scores
    scores := map[string]int{"mario": 90, "luigi": 85}

    // 'range' returns key and value when looping over a map
    for name, score := range scores {
        fmt.Printf("%s scored %d\n", name, score)
    }

    // -----------------------------
    // 8️⃣ FILTERING WITH IF INSIDE LOOP
    // -----------------------------
    fmt.Println("\n--- Filtered values from slice ---")

    // Filter values from a slice during iteration
    numbers := []int{1, 4, 7, 2, 5, 8}
    for _, n := range numbers {
        if n > 4 {
            fmt.Println("High value:", n)
        }
    }

    // -----------------------------
    // 9️⃣ EARLY RETURN IN LOOP FUNCTION
    // -----------------------------
    fmt.Println("\n--- Early return in function loop ---")

    // Define a function that searches for a fruit in the slice
    findFruit := func(fruits []string, target string) {
        for _, fruit := range fruits {
            if fruit == target {
                fmt.Println("Found:", target)
                return // exit the function early when found
            }
        }
        fmt.Println("Not found:", target)
    }

    // Call the function with different targets
    findFruit(fruits, "banana")
    findFruit(fruits, "kiwi")

    // -----------------------------
    // 🔟 INDEX-BASED SLICE MODIFICATION
    // -----------------------------
    fmt.Println("\n--- Double values in slice using index ---")

    // This loop modifies the slice in place using its index
    nums := []int{1, 2, 3}
    for i := 0; i < len(nums); i++ {
        nums[i] *= 2 // multiply each element by 2
    }
    fmt.Println("Doubled:", nums)

    // -----------------------------
    // 1️⃣1️⃣ NESTED LOOPS
    // -----------------------------
    fmt.Println("\n--- Nested loops ---")

    // Example of two loops inside each other
    for i := 1; i <= 2; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }

    // -----------------------------
    // 1️⃣2️⃣ LOOP BENCHMARKING (ADVANCED)
    // -----------------------------
    fmt.Println("\n--- Timing a loop ---")

    // Capture start time using time.Now()
    start := time.Now()

    // Example loop that does a large number of operations
    sum := 0
    for i := 0; i < 1_000_000; i++ {
        sum += i
    }

    // Capture elapsed time using time.Since()
    duration := time.Since(start)
    fmt.Println("Loop completed in:", duration)
}
