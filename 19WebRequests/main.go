package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Requests")
	response, err := http.Get(url)
	checkNilError(err)

	fmt.Println("the response is:\n", response) // gives the resposse as whole
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	content := string(databytes)
	fmt.Println(content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
