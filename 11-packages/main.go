package main

import (
	"fmt"
	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/greetings"
)

func main() {
	fmt.Println("--- Using a Custom Package ---")
	greetings.SayHello("Guillermo")
	fmt.Println("5 + 7 =", greetings.Add(5, 7))
}

