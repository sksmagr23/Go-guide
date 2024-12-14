package main

import "fmt"

func main() {
	// there is no inheritance in golang; np super or parent

	saksham := User{"Saksham", "saksham@go.org", true, 19}
	fmt.Println(saksham)
	fmt.Printf("my details are: %+v\n", saksham)
	fmt.Printf("Nmae is %v and email is %v.", saksham.Name, saksham.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
