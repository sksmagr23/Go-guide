package main

import "fmt"

func main()  {
	logcount := 20
	var result string

	if logcount < 10 {
		result = "Regular user"
	} else if logcount > 20 {
		result = "admin"
	} else {
		result = "main user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	
}