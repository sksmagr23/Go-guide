package main

import "fmt"

func main()  {
	fmt.Println("functions in golang")
	greet()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proRes, myMessage := proadder(2,5,8,3) // can use _ if necessary for other argument
	fmt.Println("pro result is:", proRes)
	fmt.Println("pro message is:", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
} // int at the end specifies function signatures


func proadder(values ...int) (int, string) {
	total := 0
	for _, val := range values{
		total += val
	}

	return total, "hi pro result function"
}

func greet()  {
	fmt.Println("hello from greet")
}