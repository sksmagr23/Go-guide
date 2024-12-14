package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
	"crypto/rand"
)

func main() {
	fmt.Println("welcome to maths in go")

	// var myNumberone int = 2
	// var myNumberTwo float64 = 4

	// fmt.Println("The sum is:", myNumberone+int(myNumberTwo))

	// random number from math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
