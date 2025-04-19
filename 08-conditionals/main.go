package main

import (
	"fmt"
)

func main() {
    // -----------------------------
    // 1️⃣ IF / ELSE IF / ELSE
    // -----------------------------
    fmt.Println("--- Basic if / else if / else ---")

    // This block evaluates the numeric value of 'score'
    // and determines the corresponding letter grade.
    score := 85

    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: F")
    }

    // -----------------------------
    // 2️⃣ IF WITH SHORT VARIABLE DECLARATION
    // -----------------------------
    fmt.Println("\n--- if with short variable declaration ---")

    // Here, we declare and use a variable ('length') within the if statement.
    // It's scoped only to this conditional block.
    if length := len("golang"); length > 5 {
        fmt.Println("It's a long word!")
    } else {
        fmt.Println("Short word")
    }

    // -----------------------------
    // 3️⃣ LOGICAL OPERATORS
    // -----------------------------
    fmt.Println("\n--- Logical operators ---")

    // Logical AND (&&) and OR (||) used to combine conditions.
    age := 22
    hasID := true

    if age >= 18 && hasID {
        fmt.Println("Entry allowed")
    }

    if age < 18 || !hasID {
        fmt.Println("Entry denied")
    }

    // -----------------------------
    // 4️⃣ NESTED IF STATEMENTS
    // -----------------------------
    fmt.Println("\n--- Nested if ---")

    // Demonstrates nested conditionals:
    // First verifies authentication, then role-based access.
    user := "admin"
    authenticated := true

    if authenticated {
        if user == "admin" {
            fmt.Println("Access granted to admin panel.")
        }
    }

    // -----------------------------
    // 5️⃣ SWITCH STATEMENT (value-based)
    // -----------------------------
    fmt.Println("\n--- switch statement ---")

    // Traditional switch using a value comparison.
    language := "go"

    switch language {
    case "python":
        fmt.Println("Interpreted language")
    case "go":
        fmt.Println("Compiled language")
    case "javascript":
        fmt.Println("Web language")
    default:
        fmt.Println("Unknown language")
    }

    // -----------------------------
    // 6️⃣ SWITCH WITH MULTIPLE MATCHES
    // -----------------------------
    fmt.Println("\n--- switch with multiple matches per case ---")

    // You can match multiple values using comma-separated cases.
    day := "saturday"

    switch day {
    case "saturday", "sunday":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Weekday.")
    }

    // -----------------------------
    // 7️⃣ SWITCH WITH EXPRESSION
    // -----------------------------
    fmt.Println("\n--- switch with expression ---")

    // Switch using an expression (e.g. x % 2)
    // Useful for logic like even/odd or value category.
    x := 10
    switch x % 2 {
    case 0:
        fmt.Println("Even number")
    case 1:
        fmt.Println("Odd number")
    }

    // -----------------------------
    // 8️⃣ SWITCH WITH FALLTHROUGH
    // -----------------------------
    fmt.Println("\n--- switch with fallthrough ---")

    // 'fallthrough' forces the next case to execute even if matched.
    // This is different from other languages where fallthrough is implicit.
    status := "active"

    switch status {
    case "active":
        fmt.Println("Status: active")
        fallthrough
    case "pending":
        fmt.Println("Also considered pending")
    default:
        fmt.Println("Unknown status")
    }

    // -----------------------------
    // 9️⃣ SWITCH WITHOUT CONDITION
    // -----------------------------
    fmt.Println("\n--- switch without condition ---")

    // A switch without a value acts like a chain of if/else if/else.
    // It evaluates conditions top-down and runs the first true case.
    temp := 35

    switch {
    case temp >= 40:
        fmt.Println("Very hot")
    case temp >= 30:
        fmt.Println("Hot")
    case temp >= 20:
        fmt.Println("Warm")
    default:
        fmt.Println("Cool")
    }

    // -----------------------------
    // 🔟 TYPE SWITCH (ADVANCED)
    // -----------------------------
    fmt.Println("\n--- Type switch ---")

    // A type switch evaluates the underlying type of an interface{}.
    // Useful when handling dynamic or mixed-type data.
    var data interface{} = 42

    switch v := data.(type) {
    case int:
        fmt.Println("Integer value:", v)
    case string:
        fmt.Println("String value:", v)
    default:
        fmt.Println("Unknown type")
    }

    // -----------------------------
    // 1️⃣1️⃣ GUARD CLAUSE PATTERN
    // -----------------------------
    fmt.Println("\n--- Guard clause pattern ---")

    // A guard clause returns early if a condition fails.
    // It helps reduce nesting and keeps logic clean.
    checkAge := func(age int) {
        if age < 18 {
            fmt.Println("Denied: must be 18+")
            return
        }
        fmt.Println("Allowed")
    }

    checkAge(15)
    checkAge(20)
}

/*
🧠 CONDITIONALS IN GO

This program explores conditional logic in Go, covering:

✅ if / else if / else
✅ short variable declarations in conditionals
✅ logical operators (AND, OR, NOT)
✅ nested conditionals for multi-layer logic
✅ switch statements — value-based, expression-based, fallthroughs, and multiple-case handling
✅ type switches — useful for dynamic or interface-based data
✅ guard clauses — a clean pattern for early exits in functions

Together, these constructs form the backbone of **decision-making** in Go,
and are critical for writing clean, safe, and expressive logic.
*/

