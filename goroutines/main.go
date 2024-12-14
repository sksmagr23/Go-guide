package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

// Declare a wait group to synchronize goroutines
var wg sync.WaitGroup // usually pointers
var mut sync.Mutex // pointer

func main() {
	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://lco.dev",
	}

	for _, web := range websitelist {
		// Start a new goroutine to get the status code of the website
		go getStatusCode(web)
		// Increment the wait group counter
		wg.Add(1)
	}

	// Wait for all goroutines to finish
	wg.Wait() // always at end of main function
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	// Decrement the wait group counter when the function completes
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

//  now the result is much faster