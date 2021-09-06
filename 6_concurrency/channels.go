package main

import (
	"fmt"
)

/*
- Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
	> ch <- v    // Send v to channel ch.
	> v := <-ch  // Receive from ch, and
	>            // assign value to v.
(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:
	> ch := make(chan int)
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// time.Sleep(5000 * time.Millisecond)
	c <- sum // send sum to c
}

// https://tour.golang.org/concurrency/2
func main() {
	// The example code sums the numbers in a slice, distributing the work between two goroutines.
	// Once both goroutines have completed their computation, it calculates the final result.
	sliceOfInts := []int{7, 2, 8, -9, 4, 0}
	partition := len(sliceOfInts) / 2 // With {7, 2, 8, -9, 4, 0} equals 3

	honestChannel := make(chan int)
	go sum(sliceOfInts[:partition], honestChannel) // it should send (7+2+8=17) to channel
	go sum(sliceOfInts[partition:], honestChannel) // it should send (-9+4+0=-5) to channel
	// By default, sends and receives block until the other side is ready.
	// This allows goroutines to synchronize without explicit locks or condition variables.

	//time.Sleep(5000 * time.Millisecond) // Try to use this one here and in the sum function!
	x, y := <-honestChannel, <-honestChannel // receive from honestChannel

	fmt.Println(x, y, x+y) // -5 17 12
}
