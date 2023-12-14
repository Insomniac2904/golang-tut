package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/insomniac2904/mongoDB/router"
)

func main() {
	fmt.Println("mongoDB api")
	r := router.Router()
	fmt.Println("SERVER GETTING READY...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at 4000")
}
