package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("dice game")

	rand.Seed(time.Now().UnixNano())

	dicenumber := rand.Intn(6) + 1
	fmt.Println("value of dice is", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
		fallthrough // next also comes
	case 2:
		fmt.Println("2 spot")
		fallthrough
	case 3:
		fmt.Println("3 spot")
	case 4:
		fmt.Println("4 spot")
	case 5:
		fmt.Println("5 spot")
	case 6:
		fmt.Println("6 spot and roll again")
	default:
		fmt.Println("what is that!")	
	}
	
}

