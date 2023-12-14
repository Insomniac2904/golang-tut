package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["C++"] = "CPP"
	languages["C"] = "C"
	languages["NJ"] = "Node JS"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "C++")
	fmt.Println(languages)

	//loops on maps

	for key, value := range languages {
		fmt.Printf("key %v, value %v\n", key, value)
	}
}
