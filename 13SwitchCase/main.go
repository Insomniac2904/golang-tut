package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceVal := rand.Intn(6) + 1
	fmt.Println("value of dice is: ", diceVal)

	switch diceVal {
	case 1:
		fmt.Println("val= 1")
	case 2:
		fmt.Println("val= 2")
		fallthrough // execute all the below cases also
	case 3:
		fmt.Println("val= 3")
		fallthrough // execute all the below cases also
	case 4:
		fmt.Println("val= 4")
	case 5:
		fmt.Println("val= 5")
	case 6:
		fmt.Println("val= 6")
	default:
		fmt.Println("default")
	}
}
