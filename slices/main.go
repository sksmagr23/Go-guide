package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("welcome to slices")

	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("type of fruit list is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)


	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 321) // apend will reallocate the memory

	fmt.Println(highScores)
    
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores) // sorts in increasing order
	fmt.Println(highScores)

	// how to remove a value from slices based on index

	var courses = []string{"react", "js", "python", "rust", "swift"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) // value at index 2 is removed

}