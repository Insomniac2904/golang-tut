package main

import "fmt"

func main() {
	loginCount := 100
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if (loginCount > 10) && (loginCount < 15) {
		result = "new user"
	} else {
		result = "hello"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("greater than 10")
	}

	// if err!=nil{
	//
	// }
}
