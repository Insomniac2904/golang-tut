package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=kdvbbkn555vn"

func main() {
	fmt.Println("handling urls in go")
	fmt.Println(myUrl)
	//parsing
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println("rawquery ", result.RawQuery)

	qparams := result.Query()
	fmt.Println("the type of query params are: ", qparams)

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/hello",
		RawPath: "user=adarsh",
	}

	anotherUrl := partsofUrl.String()
	fmt.Println("another Url is: ", anotherUrl)

}
