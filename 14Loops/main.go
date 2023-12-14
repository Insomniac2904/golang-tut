package main

import "fmt"

func main() {
	days := []string{"sunday", "monday", "tuesday", "friday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	fmt.Println("================================================")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("================================================")

	for index, day := range days {
		fmt.Println(index, day)
	}

	fmt.Println("================================================")
	for _, day := range days {
		fmt.Println(day)
	}

	fmt.Println("================================================")
	rougeVal := 1
	for rougeVal < 10 {
		if rougeVal == 2 {
			goto lco
		}

		if rougeVal == 5 {
			break
		}

		fmt.Println(rougeVal)
		rougeVal++
	}

lco:
	fmt.Println("jump")
}
