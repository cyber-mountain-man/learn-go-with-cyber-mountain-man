package greetings

import "fmt"

func SayHello(name string) {
	fmt.Printf("Hello, %s! 👋\n", name)
}

func Add(a, b int) int {
	return a + b
}