package main

import "fmt"

func main()  {
	fmt.Println("welcome to arrays in go")

	var fruit [4]string

	fruit[0] = "Apple"
	fruit[1] = "Tomato"
	fruit[3] = "Banana"

	fmt.Println("fruit list is:", fruit) // output:- fruit list is: [Apple Tomato  Banana]

	fmt.Println("list length is: ", len(fruit)) // 4(counts space)

	var veglist = [3]string{"potato", "beans","mushroom"}
	fmt.Println("veglist length is: ", len(veglist)) 
}