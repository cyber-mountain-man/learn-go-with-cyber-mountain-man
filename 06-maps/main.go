package main

import "fmt"

func main() {
    // -----------------------------
    // BASIC MAP CREATION
    // -----------------------------

    // Declare and initialize a map using a map literal.
    // The key type is string and the value type is int.
    ages := map[string]int{
        "mario": 32,
        "luigi": 30,
        "peach": 28,
    }

    // Create an empty map using the built-in make() function.
    // This allocates memory for the map so it can be written to.
    scores := make(map[string]int)

    // Add key-value pairs to the map.
    scores["yoshi"] = 100
    scores["bowser"] = 80

    // Print entire map contents.
    fmt.Println("Ages map:", ages)
    fmt.Println("Scores map:", scores)

    // -----------------------------
    // ACCESSING VALUES
    // -----------------------------

    // Access values by key using square bracket syntax.
    fmt.Println("Mario's age:", ages["mario"])
    fmt.Println("Yoshi's score:", scores["yoshi"])

    // Accessing a key that doesn't exist returns the zero value
    // of the value type — for int, it's 0.
    fmt.Println("Daisy's score (not in map):", scores["daisy"])

    // -----------------------------
    // UPDATING VALUES
    // -----------------------------

    // Assigning a new value to an existing key will update it.
    scores["yoshi"] = 105
    fmt.Println("Updated Yoshi score:", scores["yoshi"])

    // -----------------------------
    // DELETING KEYS
    // -----------------------------

    // The built-in delete() function removes a key-value pair from a map.
    delete(scores, "bowser")
    fmt.Println("After deleting bowser:", scores)

    // -----------------------------
    // CHECKING IF A KEY EXISTS
    // -----------------------------

    // When accessing a map, you can also check if the key exists.
    // The second returned value (exists) is a boolean.
    score, exists := scores["luigi"]
    if exists {
        fmt.Println("Luigi's score:", score)
    } else {
        fmt.Println("Luigi not found in scores map.")
    }

    // -----------------------------
    // ITERATING OVER A MAP
    // -----------------------------

    // Use a for loop with range to iterate over all key-value pairs in a map.
    fmt.Println("\n--- Loop through ages ---")
    for key, value := range ages {
        fmt.Printf("%s is %d years old\n", key, value)
    }

    fmt.Println("\n--- Loop through scores ---")
    for name, score := range scores {
        fmt.Printf("%s has a score of %d\n", name, score)
    }

    // -----------------------------
    // NIL MAP EXAMPLE
    // -----------------------------

    // Declaring a map without initializing it results in a nil map.
    var nilMap map[string]string
    fmt.Println("\nNil map:", nilMap)

    // Writing to a nil map causes a runtime panic:
    // nilMap["hello"] = "world" // ❌ unsafe

    // Always check if a map is nil before using it
    if nilMap == nil {
        fmt.Println("nilMap is nil and cannot be written to.")
    }

    // -----------------------------
    // MAP WITH SLICE VALUES
    // -----------------------------

    // You can also store more complex types as map values, such as slices.
    favoriteGames := map[string][]string{
        "mario": {"kart", "party", "odyssey"},
        "luigi": {"mansion", "kart"},
    }

    fmt.Println("\n--- Map with Slice Values ---")
    // Loop through the map and access each slice
    for character, games := range favoriteGames {
        fmt.Println(character, "likes:", games)
    }
}
