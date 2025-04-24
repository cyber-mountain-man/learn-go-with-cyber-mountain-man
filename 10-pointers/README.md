# 10 - Pointers in Go

Pointers in Go allow you to directly access and modify memory. This concept is essential for writing efficient and idiomatic Go code — especially when dealing with large structs, modifying data across function calls, or working with performance-critical logic.

---

## 🎯 Objectives

- Understand what pointers are and how they work
- Use pointers to read and write memory
- Modify values in functions using pointers
- Apply pointer receivers to struct methods
- Handle `nil` pointers safely

---

## 🧠 What is a Pointer?

A pointer holds the **memory address** of a variable. You use:

- `&` to get the address of a variable
- `*` to dereference a pointer and access the value it points to

---

## 🔧 Syntax and Examples

### 🔹 Creating and Using Pointers

```go
age := 30
agePtr := &age       // pointer to age
*agePtr = 35         // modifies original value
```

### 🔹 Pointers in Functions

```go
func doubleValue(x int) { x *= 2 }           // No effect on original
func doublePointer(x *int) { *x *= 2 }       // Modifies original
```

### 🔹 Pointer Receivers in Struct Methods

```go
type User struct { name string }

func (u *User) updateName(newName string) {
    u.name = newName
}
```

> Methods with pointer receivers can modify the struct’s fields.

---

## 🛑 Nil Pointers and `new`

```go
var ptr *int  // nil by default

if ptr == nil {
    fmt.Println("Pointer is nil")
}

ptr = new(int)   // allocates zero-value memory
*ptr = 50
```

---

## 🧪 Sample Output

```
--- Basic Pointers ---
Age: 30
Pointer to age: 0xc0000140b8
Value via pointer: 30
Modified age via pointer: 35

--- Pointers in Functions ---
Inside doubleValue: 20
After doubleValue: 10
Inside doublePointer: 20
After doublePointer: 20

--- Pointers with Structs ---
Before update: Peach
After update: Daisy

--- Nil Pointers ---
Pointer is nil
Value in newly allocated pointer: 50
```

---

## 🧠 Key Takeaways

| Concept                | Description |
|------------------------|-------------|
| `&x`                   | Get memory address of `x` |
| `*ptr`                 | Dereference pointer to access or modify value |
| Pointer receiver       | Lets methods mutate struct fields |
| `new(T)`               | Allocates zero-value memory and returns `*T` |
| `nil` pointers         | Pointers default to `nil` and must be checked |

---

🔁 Up Next:  
Lesson `11-packages` — learn how to organize your Go code into reusable modules.
```
