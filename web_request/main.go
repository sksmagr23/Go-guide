package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://chatgpt.com/c/6758378c-724c-8004-84cc-1ea7cbd6d7cf"

func main()  {
	fmt.Println("A sample web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("response is of type: %T\n", response)

	defer response.Body.Close() // it is caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}