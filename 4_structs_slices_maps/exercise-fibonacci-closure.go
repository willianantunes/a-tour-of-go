package main

import "fmt"

/*
Exercise: Fibonacci closure

Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers[1] (0, 1, 1, 2, 3, 5, ...).

[1] https://en.wikipedia.org/wiki/Fibonacci_number
*/

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// It works like "classes"
	// You can use internal attributes as states

	// Wikipedia: Each number is the sum of the two preceding ones, starting from 0 and 1
	f0, f1 := 0, 1

	return func() int {
		control := f0
		f0 = f1
		f1 = control + f1
		return control
	}
}

func main() {
	f := fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Println(f())
	}
}
