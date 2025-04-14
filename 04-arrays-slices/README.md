# 04 - Arrays and Slices

In this lesson, you'll explore how to store, manipulate, and expand collections of data using **arrays** and **slices** in Go. You'll also learn how to create dynamic slices efficiently using the powerful `make()` function.

---

## 🎯 Objectives

- Declare and manipulate fixed-length arrays
- Understand dynamic slices and their backing arrays
- Use `append()` to grow slices
- Slice ranges using `[:]` syntax
- Understand and apply `len()` and `cap()`
- Use `make()` to preallocate slices with length and capacity
- Learn the difference between `nil` and empty slices
- Use `copy()` to duplicate slices safely

---

## 🧪 Topics Covered

### 🔹 Arrays

```go
var ages = [3]int{20, 25, 30}
names := [4]string{"yoshi", "mario", "peach", "bowser"}
names[1] = "luigi"
```

- Arrays are **fixed in size** and stored in contiguous memory.
- Indexing is zero-based: `names[1] = "luigi"`

---

### 🔹 Slices

```go
scores := []int{100, 50, 60}
scores = append(scores, 85)
```

- Slices are more flexible than arrays.
- Backed by arrays under the hood.
- `append()` dynamically grows the slice.

---

### 🔹 Slicing Syntax

```go
rangeOne := names[1:3]   // includes index 1, excludes index 3
rangeTwo := names[2:]    // from index 2 to end
rangeThree := names[:3]  // from start to index 3 (not included)
```

- Go slices use `[start:end]` syntax (end is non-inclusive).

---

### 🔹 Length vs Capacity

```go
slice := []string{"a", "b"}
fmt.Println(len(slice)) // 2
fmt.Println(cap(slice)) // 2 (or more after append)
```

- `len()` = number of elements
- `cap()` = allocated space in memory (can grow into it)

---

### 🔹 Slice Creation with `make()`

```go
numbers := make([]int, 2, 5)
```

| Value | Meaning |
|-------|---------|
| `2`   | Length — number of accessible elements |
| `5`   | Capacity — how much space is allocated behind the scenes |

```go
numbers[0] = 7
numbers = append(numbers, 9, 10, 11)
```

- Efficiently preallocates space
- Avoids repeated reallocation

✅ Use `make()` when you **know the slice will grow** or want to **tune performance**.

---

### 🔹 Copying Slices

```go
original := []int{1, 2, 3}
copied := make([]int, len(original))
copy(copied, original)
```

- `copy()` duplicates content between slices.
- Prevents unexpected changes due to shared underlying arrays.

---

### 🔹 Nil vs Empty Slices

```go
var nilSlice []int     // nil, uninitialized
emptySlice := []int{}  // initialized but empty
```

- Both have `len() == 0`, but only `nilSlice == nil`
- Important for checking API responses, avoiding panics

---

## ✅ Output Example

```
Ages: [20 25 30] | Length: 3
Names: [yoshi luigi peach bowser] | Length: 4
Scores: [100 50 25 85] | Length: 4
Range One: [luigi peach]
Range Two: [peach bowser]
Range Three: [yoshi luigi peach]
Make slice: [7 8] | Len: 2 | Cap: 5
Make slice (after append): [7 8 9 10 11] | Len: 5 | Cap: 5
Before append: [a b] | Len: 2 | Cap: 2
After append: [a b c d e] | Len: 5 | Cap: 8
Original slice: [1 2 3]
Copied slice:   [1 2 3]
Nil slice: [] Len: 0
Empty slice: [] Len: 0
nilSlice is nil!
emptySlice is not nil!
```

---

## 🧠 Key Takeaways

- Arrays are fixed; slices are dynamic and much more commonly used
- `make()` gives you control over slice allocation and growth
- `append()` returns a **new** slice — always reassign!
- Use `copy()` when you want to avoid shared memory side-effects
- Know how to differentiate between `nil` and empty slices for safer code

---

🔁 When you're ready, move on to:  
📁 `05-structs`
```