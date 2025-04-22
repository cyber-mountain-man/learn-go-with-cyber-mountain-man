# 08.5 - Functions in Go

Functions are the building blocks of Go programs. They allow you to break code into reusable, testable, and expressive units of logic. This lesson covers core function concepts every Go developer should know.

---

## 🎯 Objectives

- Declare and call basic functions
- Pass and return values
- Work with multiple and named returns
- Use variadic functions for flexible arguments
- Define and call anonymous functions
- Treat functions as values
- Create closures with internal state

---

## 🔧 Function Types

### 🔹 Basic Function

```go
func greet() {
    fmt.Println("Hello from the greet function!")
}
```

### 🔹 Function with Parameters and Return

```go
func add(a int, b int) int {
    return a + b
}
```

### 🔹 Multiple Return Values

```go
func fullName() (string, string) {
    return "John", "Smith"
}
```

### 🔹 Named Return Values

```go
func triple(x int) (result int) {
    result = x * 3
    return
}
```

### 🔹 Variadic Function

```go
func variadicSum(nums ...int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}
```

### 🔹 Anonymous Function

```go
func(msg string) {
    fmt.Println("Message:", msg)
}("Hello Go Learner!")
```

### 🔹 Function as a Value

```go
double := func(n int) int {
    return n * 2
}
fmt.Println(double(6)) // 12
```

### 🔹 Closure (Function with Memory)

```go
counter := func() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}()
```

---

## 🧪 Sample Output

```
--- Basic function call ---
Hello from the greet function!

--- Function with parameters and return value ---
5 + 3 = 8

--- Function with multiple return values ---
Full name: John Smith

--- Named return value ---
Triple of 7 is: 21

--- Variadic function ---
Sum: 15

--- Anonymous function ---
Message: Hello Go Learner!

--- Function as value ---
Double of 6 is: 12

--- Closure function ---
1
2
3
```

---

## 🧠 Key Takeaways

| Feature              | Description |
|----------------------|-------------|
| Functions            | Encapsulate logic into reusable blocks |
| Parameters/Returns   | Control input/output of logic |
| Multiple Returns     | Enable expressive, idiomatic code |
| Named Returns        | Improve clarity in certain functions |
| Variadic             | Flexible argument handling |
| Anonymous Functions  | Inline logic for single-use cases |
| Functions as Values  | Store and reuse logic like variables |
| Closures             | Maintain private state across calls |

---

🔁 Up Next:  
Lesson `09-methods` — where you'll apply functions as behaviors tied to data (structs).

```