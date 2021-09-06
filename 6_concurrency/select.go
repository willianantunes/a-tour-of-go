package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// The select statement lets a goroutine wait on multiple communication operations.
		// A select blocks until one of its cases can run (THAT'S REALLY NICE), then it executes that case.
		// It chooses one at random if multiple are ready.
		select {
		// The absence of a default case means "If I can't send to c and I can't read from quit, block until I can."
		case c <- x: // If I can send to c
			x, y = y, x+y
		case <-quit: // If I can receive from quit then I'm supposed to exit
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// Delightful explanation: https://stackoverflow.com/a/34931225
	// Using go to spin off the anonymous goroutine function lets you consume from c so fibonacci will continue!
	go func() {
		for i := 0; i < 10; i++ {
			// time.Sleep(10000 * time.Millisecond) // // You can notice how the `c <- x` works through this sleep more clearly!
			whatWhatSendThroughFibonacciFunction := <-c
			fmt.Println(whatWhatSendThroughFibonacciFunction)
			// fmt.Println(<-c) // Send what is in the channel to `Println`
		}
		// time.Sleep(10000 * time.Millisecond) // You can notice how the `case <- quit` works through this sleep more clearly!
		quit <- 0 // Send 0 to channel quit, it means the case above will be satisfied, then killing the main process
	}()
	fibonacci(c, quit)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
