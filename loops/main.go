package main

import "fmt"

func main()  {
	fmt.Println("loops in golang")

	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	// fmt.Println(days)

	for d:=0; d< len(days); d++{
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days{
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	for _, day := range days{
		fmt.Printf("value is %v\n", day)
	}

	rougevalue := 1

	for rougevalue < 10 {

		if rougevalue == 2 {
			goto jlo
		}

		// if rougevalue == 5 {
		// 	break
		// }
		if rougevalue == 7 {
			rougevalue++
			continue
		}

		fmt.Println("value is:", rougevalue)
		rougevalue++    // wrong ++rougevalue
	} 

	jlo:
	    fmt.Println("i jumped from loop")




}