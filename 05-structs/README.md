### ✅ `05-structs/README.md`

```markdown
# 05 - Structs in Go

Structs in Go are used to define complex, custom data types. In this lesson, we cover how to create and use structs, define methods, and extend your understanding with advanced topics like struct embedding, anonymous structs, slices of structs, and JSON encoding.

---

## 🎯 Objectives

- Define and instantiate structs
- Access and modify struct fields
- Attach methods to structs using value and pointer receivers
- Understand and apply:
  - Struct embedding (composition)
  - Anonymous structs
  - Slice of structs
  - JSON tags for marshaling

---

## 🧱 Struct Basics

### 🛠 Defining a Struct

```go
type User struct {
    name     string
    age      int
    verified bool
}
```

### 🧪 Creating Instances

```go
user1 := User{name: "Mario", age: 30}
user2 := User{"Luigi", 28, false}
var user3 User
user3.name = "Peach"
```

### 🔍 Accessing Fields

```go
fmt.Println(user1.name)
user2.age = 29
```

---

## 🧠 Struct Methods

```go
func (u User) greet() {
    fmt.Printf("Hi, I'm %s\n", u.name)
}
```

- **Value receiver**: copies the struct, can't modify original
- **Pointer receiver**: can modify the original struct

```go
func (u *User) verifyAccount() {
    u.verified = true
}
```

---

## 🔁 Struct Embedding

```go
type Admin struct {
    User
    accessLevel string
}
```

- Embeds `User` directly into `Admin`
- Inherits all fields and methods

```go
admin := Admin{
    User:        User{name: "Toad", age: 35},
    accessLevel: "super",
}
admin.greet() // from embedded User
```

---

## ⚡ Anonymous Structs

```go
guest := struct {
    name string
    age  int
}{
    name: "Daisy",
    age:  22,
}
fmt.Println(guest.name)
```

- Useful for short-lived, one-off data models (e.g., quick tests, small internal configs)

---

## 🧱 Slice of Structs

```go
users := []User{
    {name: "Yoshi", age: 21, verified: true},
    {name: "Rosalina", age: 29},
}
for _, u := range users {
    u.greet()
}
```

- Very useful for storing and working with collections of structured data

---

## 🌐 JSON Tags (Marshaling)

```go
type Profile struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

p := Profile{Name: "Guillermo", Email: "g@example.com"}
jsonData, _ := json.Marshal(p)
fmt.Println(string(jsonData)) // {"name":"Guillermo","email":"g@example.com"}
```

- JSON tags tell the encoder how to name keys when marshaling
- Required for building APIs or storing data as JSON

---

## ✅ Sample Output

```
{Mario 30 false}
{Luigi 28 false}
{Peach 0 false}
User1 name: Mario
User2 age: 28
Hello, my name is Mario and I am 30 years old.
Hello, my name is Luigi and I am 28 years old.
Hello, my name is Peach and I am 25 years old.
Verified status (before): false
Mario has been verified.
Verified status (after): true

--- Struct Embedding ---
Admin name: Toad
Access level: super
Hello, my name is Toad and I am 35 years old.

--- Anonymous Struct ---
Guest: Daisy (22 years old)

--- Slice of Structs ---
Hello, my name is Yoshi and I am 21 years old.
Hello, my name is Rosalina and I am 29 years old.

--- JSON Output ---
{"name":"Guillermo","email":"g@example.com"}
```

---

## 🧠 Key Takeaways

- Structs are Go’s foundation for modeling custom data types
- Methods let structs encapsulate behavior
- Use pointer receivers when you need to **modify** struct fields
- Structs can be composed using embedding instead of inheritance
- Slices of structs are essential for managing collections of data
- JSON tags prepare your data for API use and marshaling

---

🔁 Next:  
Ready for `06-maps` or `06-loops`? Let’s continue building!
```