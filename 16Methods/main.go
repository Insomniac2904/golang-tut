package main

import "fmt"

func main() {
	// no inheritance in golang and no super or parent or child
	adarsh := User{"adarsh", "asd@asd", true, 23} // mid the { } brackets
	fmt.Println(adarsh)
	fmt.Println(adarsh.Name, adarsh.Age, adarsh.Email)
	fmt.Printf("details are: %+v\n", adarsh)
	adarsh.getStatus()
	adarsh.setEmail()
	fmt.Println(adarsh.Email) // return the old value
	adarsh.setEmailByReference()
	fmt.Println(adarsh.Email) // return the old value
}

type User struct { // capitalize first letters
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() { // !method
	fmt.Println("Is user active: ", u.Status)
}

func (u User) setEmail() { // !as the value is pass by value the value will not change inplace
	u.Email = "adarsh@ak.com"
	fmt.Println("Email is: ", u.Email)
}
func (u *User) setEmailByReference() { // !as the value is pass by reference the value will change inplace
	u.Email = "adarsh@ak.com"
	fmt.Println("Email is: ", u.Email)
}
