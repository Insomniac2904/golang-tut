package main

import "fmt"

func main() {
	// fmt.Println("hello")
	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 25
	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr + 2
	fmt.Println(*ptr)
}
