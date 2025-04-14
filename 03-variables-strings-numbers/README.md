# 03 - Variables, Strings, and Numbers

In this lesson, we explore one of the foundational aspects of Go: **variables** and **basic data types**. You’ll learn how Go handles strings, integers, floating-point numbers, memory-specific types, and how it deals with default values and type inspection.

---

## 🎯 Objectives

- Declare variables using `var` and `:=`
- Understand static typing and type inference
- Work with strings, integers, and floating-point numbers
- Learn about `rune` and `byte` types
- Practice type conversion and memory-aware types (e.g., `int8`, `float64`)
- Inspect variable types at runtime using `reflect` and `fmt`

---

## 🧪 Topics Covered

### 🔹 Variable Declarations

```go
var nameOne string = "mario"
var nameTwo = "luigi"
var nameThree string
nameThree = "bowser"
nameFour := "yoshi"
```

- `var` allows explicit type declaration
- `:=` lets Go infer the type (only usable inside functions)

---

### 🔹 Strings

```go
greeting := "Hello"
message := greeting + ", " + nameOne + "!"
formatted := fmt.Sprintf("%s, %s!", greeting, nameTwo)
```

- Concatenation uses `+`
- `fmt.Sprintf` creates formatted strings (without printing)

---

### 🔹 Integers

```go
var ageOne int = 20
var ageTwo = 30
ageThree := 40
```

- `int` is platform-dependent (usually `int64`)
- Go does **not** allow implicit type casting

---

### 🔹 Type Conversion

```go
var intVal int = 42
var floatVal float64 = float64(intVal)
var uintVal uint = uint(floatVal)
```

- Always use explicit casting: `float64(x)`, `uint(y)`, etc.

---

### 🔹 Floating-Point Numbers

```go
var scoreOne float32 = 25.98
var scoreTwo float64 = 888375661515454156166226.7
scoreThree := 1.5
```

- `float64` is the default type for floats
- Use `float32` for memory efficiency when precision isn’t critical

---

### 🔹 Rune & Byte Types

```go
var letter rune = 'G'
var b byte = 'a'
```

- `rune` (alias for `int32`) holds Unicode characters
- `byte` (alias for `uint8`) is for raw ASCII bytes

---

### 🔹 Default (Zero) Values

```go
var emptyString string      // ""
var zeroInt int             // 0
var zeroFloat float64       // 0.0
var isReady bool            // false
```

- Uninitialized variables in Go are automatically assigned a **zero value**

---

### 🔹 Type Inspection

```go
x := 10
fmt.Println(reflect.TypeOf(x))        // prints: int
fmt.Printf("Type of y: %T\n", "Go")   // prints: string
```

- Use `reflect.TypeOf()` or `%T` to check a variable’s type

---

## ✅ Output Example

```
Names: mario luigi bowser yoshi
Concatenated: Hello, mario!
Formatted: Hello, luigi!
Ages: 20 30 40
Converted types: 42 42 42
Scores: 25.98 8.883756615154542e+23 1.5
Rune: 71 Byte: 97
Zero values ->  0 0 false
Type of x (using reflect): int
Type of y (using %T): string
```

---

## 🧠 Key Takeaways

- Go is statically typed but supports inference
- Variable declarations and memory size awareness are explicit in Go
- Characters are stored as `rune` or `byte`, not `char`
- Type conversion is always manual — no implicit coercion
- Go assigns zero values by default — no garbage data

---

🔁 Ready to continue? Move on to:  
📁 `04-arrays-slices`
```