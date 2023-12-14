package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("hello from 21")
	// PerformGetRequest()
	// PerformPostRequest()
	performFormPost()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content) // got raw data
	fmt.Println("using string builder size if ", byteCount)
	fmt.Println("using string builder string is ", responseString.String())

	fmt.Println("using string", string(content))
}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"name": "Adarsh",
            "age": 21
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func performFormPost() {
	const myUrl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("name", "adarsh")
	data.Add("age", "21")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
