package main

import "fmt"

func main() {
	fmt.Println("hello from main.")
	greeter() // cannot define func inside func
	result := adder(3, 4)
	fmt.Println("Result: ", result)
	proRes, comment := proAdder(2, 5, 6, 7, 2, 8, 2)
	fmt.Println(proRes)
	fmt.Println(comment)
}

func adder(a int, b int) int { // return type after parameter
	return a + b
}

func proAdder(values ...int) (int, string) { // when the no of parameters not pre known
	total := 0 // we can send and get more than one result as return and hence we must declare the types... to send the inputs use comma ok syntax
	for _, val := range values {
		total += val
	}
	return total, "hello i am pro addder"
}

func greeter() {
	fmt.Println("hello from greeter")
}
