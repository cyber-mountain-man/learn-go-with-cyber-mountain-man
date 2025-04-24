# 09 - Methods in Go

Go doesn’t have classes, but it lets you define **methods on types**, which enables powerful and clean object-oriented-like behavior. This lesson builds directly on what you learned in functions and shows how to attach behavior to structs.

---

## 🎯 Objectives

- Define methods using both value and pointer receivers
- Understand when to use each receiver type
- Implement method chaining
- Use methods to safely mutate struct fields
- Inherit behavior via struct embedding

---

## 🔧 Method Syntax in Go

### 🔹 Value Receiver

Does **not** modify the original struct (operates on a copy):

```go
func (u User) greet() {
    fmt.Printf("Hi, I'm %s\n", u.name)
}
```

### 🔹 Pointer Receiver

**Does** modify the original struct (works on the actual memory):

```go
func (u *User) verifyAccount() {
    u.verified = true
}
```

---

## 🔄 Method Chaining

```go
func (u *User) setName(name string) *User {
    u.name = name
    return u
}

user.setName("Peach").setAge(20).verifyAccount()
```

> 🧠 Useful for fluent APIs or config builders

---

## 🛑 Safe Setters with Validation

```go
func (u *User) setAgeSafe(age int) error {
    if age < 0 || age > 130 {
        return errors.New("invalid age")
    }
    u.age = age
    return nil
}
```

---

## 🧱 Struct Embedding for Method Inheritance

```go
type Admin struct {
    User
    privileged bool
}
```

The embedded `User` methods (e.g. `.greet()`, `.verifyAccount()`) are available directly on `Admin`:

```go
admin := Admin{User: User{name: "Luigi"}}
admin.greet()
```

---

## 🧪 Sample Output

```
--- Greeting Users ---
Hi, I'm Mario and I'm 25 years old.
Hi, I'm Peach and I'm 17 years old.

--- Age Check ---
Mario adult? true
Peach adult? false

--- Value vs Pointer Receiver ---
Age after updateAgeNoEffect: 25
Age after setAge: 99

--- Method Chaining ---
Hi, I'm Princess and I'm 20 years old.
Verified: true

--- Safe Setter with Error Handling ---
Failed to set age: invalid age: must be between 0 and 130
Age successfully updated to: 30

--- Embedded Struct with Methods ---
Hi, I'm Luigi and I'm 40 years old.
Admin access level: super admin
```

---

## 🧠 Key Takeaways

| Feature                   | Purpose                                          |
|---------------------------|--------------------------------------------------|
| `func (t T)`              | Value receiver (read-only)                      |
| `func (t *T)`             | Pointer receiver (mutable)                      |
| Method chaining           | Build cleaner fluent APIs                       |
| Safe setters              | Add validation and control to struct mutation   |
| Struct embedding          | Share and reuse behavior via composition        |

---

🔁 Up Next:  
Lesson `10-pointers` — understanding how references and memory access work at a deeper level in Go.

```