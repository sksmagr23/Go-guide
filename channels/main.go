package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2) // size 2
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh) // will give deadlock error
	// myCh <- 5

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup){ // going outside the box only
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen) // for 0 check
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup){ // going inside the channel
		// myCh <- 5
		// myCh <- 6
		myCh <- 0 // true
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}