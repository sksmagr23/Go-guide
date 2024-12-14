package main

import "fmt"

func main() {
	saksham := User{"Saksham", "saksham@go.org", true, 19}
	fmt.Println(saksham)
	fmt.Printf("my details are: %+v\n", saksham)
	fmt.Printf("Nmae is %v and email is %v.\n", saksham.Name, saksham.Email)

	saksham.GetStatus()
	saksham.NewMail() // it passes a copy

	fmt.Printf("Nmae is %v and email is %v.\n", saksham.Name, saksham.Email) // the old email remains same

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active:", u.Status)
} // a method created

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("new Email of this user is:", u.Email)
}
