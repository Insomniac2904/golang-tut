package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("welcome")
	fmt.Println("enter input as number:")
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Println("thanks for number:",input)

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("added 1 to rating: ",numRating+1)
	}
}