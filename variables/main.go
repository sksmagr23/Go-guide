package main

import "fmt"

// jwtToken := 30000 not allowed
// var jwttoken = 30000 allowed

// a variable starting with capital letter is public 
const Logintoken string = "jefejfie"


func main() {
	var username string = "saksham"
	fmt.Println(username)
	fmt.Printf("variable of type: %T \n", username)

	// same with boolean

	var smallval uint = 250
	fmt.Println(smallval)
	fmt.Printf("variable of type: %T \n", smallval)

	var smallfloat float32 = 255.6756564644
	fmt.Println(smallfloat)
	fmt.Printf("variable of type: %T \n", smallfloat)

	// if variable declared and not given value, by default it shows 0(for int float) not garbage.

	var web = "hi"
	fmt.Println(web) // implicit

	// no var style
	 numberOfUser := 300000
	 fmt.Println(numberOfUser)
}
