# 08 - Conditionals in Go

Conditional logic is essential to any program — it helps your code make decisions.  
Go provides clean and expressive ways to handle conditionals using:

- `if / else if / else`
- `switch` statements (with values, expressions, and fallthrough)
- `type switches` for dynamic types
- Guard clauses to keep logic concise and readable

---

## 🎯 Objectives

- Use `if` statements to check conditions and branch logic
- Apply logical operators (`&&`, `||`, `!`)
- Write `switch` statements for cleaner alternatives to `if/else`
- Use multiple conditions, expressions, and `fallthrough` in `switch`
- Leverage `type switch` for dynamic type handling
- Implement guard clause patterns to reduce nesting

---

## 🔍 if / else if / else

```go
score := 85

if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("F")
}
```

---

## ✍️ Short Variable Declaration in Conditionals

```go
if length := len("golang"); length > 5 {
    fmt.Println("Long word")
}
```

---

## 🔗 Logical Operators

```go
if age >= 18 && hasID {
    fmt.Println("Entry allowed")
}

if age < 18 || !hasID {
    fmt.Println("Entry denied")
}
```

---

## 🔄 Nested Conditionals

```go
if authenticated {
    if user == "admin" {
        fmt.Println("Access granted")
    }
}
```

---

## 🔀 switch Statements

### Value-based:

```go
switch language {
case "go":
    fmt.Println("Compiled")
case "python":
    fmt.Println("Interpreted")
}
```

### Multiple matches:

```go
switch day {
case "saturday", "sunday":
    fmt.Println("Weekend!")
}
```

### Expression-based:

```go
switch x % 2 {
case 0:
    fmt.Println("Even")
}
```

### Fallthrough:

```go
switch status {
case "active":
    fmt.Println("Active")
    fallthrough
case "pending":
    fmt.Println("Also pending")
}
```

---

## 🧠 Type Switch

```go
var data interface{} = "hello"

switch v := data.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
}
```

Used to safely inspect the underlying type of dynamic values (`interface{}`).

---

## ✅ Guard Clause Pattern

```go
func checkAge(age int) {
    if age < 18 {
        fmt.Println("Denied")
        return
    }
    fmt.Println("Allowed")
}
```

Avoids deep nesting and keeps logic clean.

---

## 📌 Sample Output

```
Grade: B
It's a long word!
Entry allowed
Access granted to admin panel.
Compiled language
Weekend!
Even number
Status: active
Also considered pending
Hot
Integer value: 42
Denied: must be 18+
Allowed
```

---

## 🧠 Key Takeaways

| Feature | What It Teaches |
|---------|-----------------|
| `if` / `else` | Basic conditional branching |
| `switch` | Cleaner logic for value comparisons |
| `fallthrough` | Go-specific control flow feature |
| Type switch | Enables runtime type inspection |
| Guard clause | Promotes readable, flat logic structure |
| Logical operators | Build complex conditions cleanly |

---