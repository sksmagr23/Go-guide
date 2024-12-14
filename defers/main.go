package main

import "fmt"

func main() {
	defer fmt.Println("World") // this line having defer runs and prints at very last of func
	defer fmt.Println("one")
	defer fmt.Println("two")

	// defer follows last in first out

	fmt.Println("hello")
	mydefer()

}

// output:-
// hello
// two
// one
// world

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Output:-
// hello
// 4
// 3
// 2
// 1
// 0
// two
// one
// World
