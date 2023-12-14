package main

import "fmt"

func main() {
	// no inheritance in golang and no super or parent or child
	adarsh := User{"adarsh", "asd@asd", true, 23} // mid the { } brackets
	fmt.Println(adarsh)
	fmt.Println(adarsh.Name, adarsh.Age)
	fmt.Printf("details are: %+v\n", adarsh)
}

type User struct { // capitalize first letters
	Name   string
	Email  string
	Status bool
	Age    int
}
