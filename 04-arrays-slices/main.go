package main

import "fmt"

func main() {
    // --- ARRAYS ---

    // Declare a fixed-size array with 3 integers
    var ages = [3]int{20, 25, 30}

    // Declare and initialize a fixed-size array of 4 strings
    names := [4]string{"yoshi", "mario", "peach", "bowser"}
    names[1] = "luigi" // Reassign one of the values

    fmt.Println("Ages:", ages, "| Length:", len(ages))
    fmt.Println("Names:", names, "| Length:", len(names))

    // --- SLICES ---

    // Create a slice from literal values (uses an underlying array)
    var scores = []int{100, 50, 60}
    scores[2] = 25 // Modify element
    scores = append(scores, 85) // Append new element

    fmt.Println("Scores:", scores, "| Length:", len(scores))

    // --- SLICE RANGES ---

    // Slicing arrays (start:stop) - stop index is excluded
    rangeOne := names[1:3]   // ["luigi", "peach"]
    rangeTwo := names[2:]    // ["peach", "bowser"]
    rangeThree := names[:3]  // ["yoshi", "luigi", "peach"]

    fmt.Println("Range One:", rangeOne)
    fmt.Println("Range Two:", rangeTwo)
    fmt.Println("Range Three:", rangeThree)

    // Append to a slice range (creates a new underlying array if needed)
    rangeOne = append(rangeOne, "koopa")
    fmt.Println("Range One (after append):", rangeOne)

    // --- SLICE WITH MAKE ---

    // Create a slice using make() with length 2 and capacity 5
    numbers := make([]int, 2, 5)
    numbers[0] = 7
    numbers[1] = 8
    fmt.Println("Make slice:", numbers, "| Len:", len(numbers), "| Cap:", cap(numbers))

    // Append beyond initial length
    numbers = append(numbers, 9, 10, 11)
    fmt.Println("Make slice (after append):", numbers, "| Len:", len(numbers), "| Cap:", cap(numbers))

    // --- CAPACITY VS LENGTH ---

    short := []string{"a", "b"}
    fmt.Println("Before append:", short, "Len:", len(short), "Cap:", cap(short))

    short = append(short, "c", "d", "e")
    fmt.Println("After append:", short, "Len:", len(short), "Cap:", cap(short))

    // --- COPYING SLICES ---

    original := []int{1, 2, 3}
    copied := make([]int, len(original)) // Preallocate a slice of same length
    copy(copied, original)

    fmt.Println("Original slice:", original)
    fmt.Println("Copied slice:  ", copied)

    // --- NIL VS EMPTY SLICE ---

    var nilSlice []int           // Declared but uninitialized (nil)
    emptySlice := []int{}        // Initialized as an empty slice

    fmt.Println("Nil slice:", nilSlice, "Len:", len(nilSlice))
    fmt.Println("Empty slice:", emptySlice, "Len:", len(emptySlice))

    // Check if slice is nil
    if nilSlice == nil {
        fmt.Println("nilSlice is nil!")
    }

    if emptySlice == nil {
        fmt.Println("emptySlice is nil!")
    } else {
        fmt.Println("emptySlice is not nil!")
    }
}
