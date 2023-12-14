package main

import "fmt"

func main() {
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[2] = "banana"
	fruitlist[3] = "peach"
	fmt.Println(len(fruitlist))
	fmt.Println(fruitlist[1])
	fmt.Println("fruitlist is: ", fruitlist)

	var vegList = [3]string{"potato", "tomato", "mushroom"}
	fmt.Println(vegList)
}
