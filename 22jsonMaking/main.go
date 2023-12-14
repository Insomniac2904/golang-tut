package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"Call me"`
	Price    int
	Password string   `json:"-"`              // does not show the fiels that have "-"  as json value
	Tags     []string `json:"tags,omitempty"` //omitempty omits the empty array when value is null . must not put space between tags & omitempty
}

func main() {
	fmt.Println("hello from json making")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	newCourses := []course{
		{"React", 100, "hello is Password", []string{"tag1", "tag2", "tag3"}},
		{"Angular", 100, "hello is Password", []string{"tag1", "tag2", "tag3"}},
		{"Vue", 100, "hello is Password", nil},
	}

	// finalJson, err := json.Marshal(newCourses)
	finalJson, err := json.MarshalIndent(newCourses, "", "\t") // indenting the final json usign marshalIndent
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "Call me": "React",
                "Price": 100,
                "tags": [
                        "tag1",
                        "tag2",
                        "tag3"
                ]
        }`)

	var details course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON WAS VALID")
		json.Unmarshal(jsonDataFromWeb, &details)
		fmt.Printf("%#v\n", details)
	} else {
		fmt.Println("INVALID JSON")
	}

	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	for key, val := range onlineData {
		fmt.Printf("key: %v , value: %v \n", key, val)
	}
}
