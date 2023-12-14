package main

import "fmt"

func main() {
	// defers print in reverse order of their call
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("World") // defer picks the command and pashtes it just above the closing bracket of the function.
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
