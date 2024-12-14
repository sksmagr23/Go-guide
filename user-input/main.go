package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "welcome to user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating of pizza:")

	// comma ok // err err
	input, _ := reader.ReadString('\n')       // _ can have err
	fmt.Println("thanks for rating,", input)
	fmt.Printf("type of rating is %T", input)
	fmt.Printf("\n")
}