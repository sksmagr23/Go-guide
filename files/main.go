package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "this need to go in my file"

	file, err := os.Create("./mygofile.txt")

	if err != nil {
		panic(err) // it shutdown the program execution and show error
	}
	// checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./mygofile.txt")
}

func readFile(filename string)  {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("text data inside the file is \n", string(databyte))
}

// now we check error various times and everytime writig err if , so we can create a function for same
func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}