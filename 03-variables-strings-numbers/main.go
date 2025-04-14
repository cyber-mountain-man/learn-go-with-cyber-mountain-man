package main

import (
    "fmt"       // Used for formatted I/O operations like printing to the console
    "reflect"   // Used to inspect the type of a variable at runtime
)

func main() {
    // -----------------------------
    // STRING VARIABLES
    // -----------------------------

    // Declaring a string with explicit type and initial value
    var nameOne string = "mario"

    // Type is inferred as string based on assigned value
    var nameTwo = "luigi"

    // Declaring without initialization (defaults to zero value: "")
    var nameThree string

    // Assigning a value to nameThree
    nameThree = "bowser"

    // Short-hand declaration using :=
    nameFour := "yoshi"

    // Print all string variables
    fmt.Println("Names:", nameOne, nameTwo, nameThree, nameFour)

    // Concatenating strings with the + operator
    greeting := "Hello"
    fullMessage := greeting + ", " + nameOne + "!"
    fmt.Println("Concatenated:", fullMessage)

    // Using Sprintf for formatted string output (like printf in C)
    formatted := fmt.Sprintf("%s, %s!", greeting, nameTwo)
    fmt.Println("Formatted:", formatted)

    // -----------------------------
    // INTEGER VARIABLES
    // -----------------------------

    // Explicit type and initialization
    var ageOne int = 20

    // Type inferred
    var ageTwo = 30

    // Short-hand declaration
    ageThree := 40

    fmt.Println("Ages:", ageOne, ageTwo, ageThree)

    // Demonstrating type conversion
    var intVal int = 42

    // Explicit conversion from int to float64
    var floatVal float64 = float64(intVal)

    // Explicit conversion from float64 to uint
    var uintVal uint = uint(floatVal)

    fmt.Println("Converted types:", intVal, floatVal, uintVal)

    // -----------------------------
    // FLOAT VARIABLES
    // -----------------------------

    // float32 allows less precision but uses less memory
    var scoreOne float32 = 25.98

    // float64 is the default type for floating-point numbers and has higher precision
    var scoreTwo float64 = 888375661515454156166226.7

    // Type is inferred as float64
    scoreThree := 1.5

    fmt.Println("Scores:", scoreOne, scoreTwo, scoreThree)

    // -----------------------------
    // RUNE & BYTE TYPES
    // -----------------------------

    // rune represents a Unicode code point (alias for int32)
    var letter rune = 'G'

    // byte represents a single ASCII character (alias for uint8)
    var b byte = 'a'

    fmt.Println("Rune:", letter, "Byte:", b)

    // -----------------------------
    // ZERO VALUES (DEFAULTS)
    // -----------------------------

    // Declaring variables without initializing assigns their type's zero value
    var emptyString string      // ""
    var zeroInt int             // 0
    var zeroFloat float64       // 0.0
    var isReady bool            // false

    // Print the default zero values
    fmt.Println("Zero values ->", emptyString, zeroInt, zeroFloat, isReady)

    // -----------------------------
    // TYPE INSPECTION
    // -----------------------------

    x := 10
    y := "GoLang"

    // reflect.TypeOf shows the dynamic type at runtime
    fmt.Println("Type of x (using reflect):", reflect.TypeOf(x))

    // %T with Printf shows the type inline
    fmt.Printf("Type of y (using %%T): %T\n", y)
}
