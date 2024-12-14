package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://chatgpt.com/c/6758378c-724c-8004-84cc-1ea7cbd6d7cf"

func main()  {
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl) 

	fmt.Println(result.Scheme) // https
	fmt.Println(result.Host) 
	fmt.Println(result.Path) 
	fmt.Println(result.RawQuery) 
	fmt.Println(result.Port()) 

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams) // a map

	for _, val := range qparams {
		fmt.Println("param is:", val)
	}

	// constructing a url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=saksham",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}