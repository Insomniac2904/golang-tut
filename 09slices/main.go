package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruit = []string{"apple", "tomato", "peach"} // need to declare as empty ot some value
	fruit = append(fruit, "mango", "banana")
	fmt.Println(fruit)
	fruit = append(fruit[1:]) // number exclusive
	fmt.Println(fruit)
	fruit = append(fruit[1:3]) // number exclusive
	fmt.Println(fruit)

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 555
	highscores[2] = 645
	highscores[3] = 789

	highscores = append(highscores, 555, 666, 777) // appended
	fmt.Println(highscores)
	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	/// removing elements from the slice
	var course = []string{"reactjs", "javascript", "swift", "python", "rust"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)

}
