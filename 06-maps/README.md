Absolutely, Guillermo! 🗺️  
Here’s your polished and professional `README.md` for:

📁 `06-maps`

It covers:
- 🔑 Map basics (creation, access, update, delete)
- ✅ Existence checking
- 🔁 Iteration
- ⚠️ Nil maps
- 🧩 Advanced: map with slice values

---

### ✅ `06-maps/README.md`

```markdown
# 06 - Maps in Go

Maps are one of the most powerful and flexible data structures in Go. They allow you to store and retrieve values using keys — similar to dictionaries or hash tables in other languages.

---

## 🎯 Objectives

- Create maps using literals and `make()`
- Access, update, and delete key-value pairs
- Check for key existence safely
- Iterate over a map
- Handle nil maps correctly
- Store complex values like slices in a map

---

## 🔑 Map Basics

### 🧱 Creating Maps

```go
// Using a map literal
ages := map[string]int{
    "mario": 32,
    "luigi": 30,
}

// Using make()
scores := make(map[string]int)
scores["yoshi"] = 100
```

---

## 🔍 Accessing & Updating Keys

```go
fmt.Println(ages["mario"]) // 32

// Update an existing key
scores["yoshi"] = 105
```

---

## 🧹 Deleting Keys

```go
delete(scores, "bowser")
```

- `delete()` safely removes a key-value pair
- Does nothing if the key does not exist

---

## ✅ Existence Check

```go
score, exists := scores["luigi"]
if exists {
    fmt.Println("Score:", score)
} else {
    fmt.Println("Luigi not found")
}
```

- Prevents false assumptions (e.g. mistaking `0` as a real score)

---

## 🔁 Iterating Over a Map

```go
for key, value := range ages {
    fmt.Printf("%s is %d years old\n", key, value)
}
```

- Order is **not guaranteed** — maps are unordered in Go

---

## ⚠️ Nil Maps

```go
var nilMap map[string]string
```

- Declaring a map without initializing it results in a `nil` map
- Reading is fine, but **writing will cause a panic**
- Use `make()` to safely initialize:

```go
safeMap := make(map[string]string)
```

---

## 🧩 Map with Slice Values

```go
favoriteGames := map[string][]string{
    "mario": {"kart", "party", "odyssey"},
    "luigi": {"mansion", "kart"},
}
```

- You can store more complex types as values (like slices, structs, etc.)
- Great for representing one-to-many relationships

---

## ✅ Sample Output

```
Ages map: map[mario:32 luigi:30 peach:28]
Scores map: map[yoshi:100 bowser:80]
Mario's age: 32
Yoshi's score: 100
Daisy's score (not in map): 0
Updated Yoshi score: 105
After deleting bowser: map[yoshi:105]
Luigi not found in scores map.

--- Loop through ages ---
mario is 32 years old
luigi is 30 years old
peach is 28 years old

--- Loop through scores ---
yoshi has a score of 105

Nil map: map[]
nilMap is nil and cannot be written to.

--- Map with Slice Values ---
mario likes: [kart party odyssey]
luigi likes: [mansion kart]
```

---

## 🧠 Key Takeaways

- Maps in Go are fast, flexible, and built-in
- Always check for key existence when needed
- Avoid writing to uninitialized (nil) maps
- Use maps with complex values like slices to model real-world data

---

🔁 Next:  
Ready for `07-loops`? Let’s continue building!
```