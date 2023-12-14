package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("handling files")
	content := "lorem impsum dollar emit"
	//!create a file
	file, error := os.Create("./outputFile.txt")
	// fmt.Println("the name of the file is:", file.Name())
	checkNilError(error)
	//! add data to file
	length, error := io.WriteString(file, content)

	checkNilError(error)
	fmt.Println("length of file is: ", length)
	defer file.Close()

	//!read the file
	readFile(file.Name())
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("text inside the file in datbytes is: ", databyte)
	fmt.Println("text inside the file in string is: ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
