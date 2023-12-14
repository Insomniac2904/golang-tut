package main

import (
	"fmt"
	"time"
)

// GOOS="windows" go build --> to build exe for windows
func main() {
	println("hello from time")
	presentTime := time.Now()
	createdAt := time.Date(2023, time.September, 29, 23, 0, 0, 0, time.UTC)
	fmt.Println(createdAt.Format("01-02-2006 Monday"))
	fmt.Println(presentTime)
}
