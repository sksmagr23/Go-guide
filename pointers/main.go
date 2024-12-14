package main

import "fmt"

func main()  {
	fmt.Println("welcome to pointers")

	var ptr *int
	fmt.Println("value of pointer is ", ptr) // output will be <nil>

	myNumber := 23

	var pitr = &myNumber
	fmt.Println("value of actual pointer is ", pitr) // address
	fmt.Println("value of actual pointer is ", *pitr) // actual value

	*pitr = *pitr + 2
	fmt.Println("new value is: ", myNumber)

}