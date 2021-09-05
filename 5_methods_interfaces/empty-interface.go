package main

import "fmt"

func main() {
	// The interface type that specifies zero methods is known as the empty interface:
	var i interface{}
	describe(i) // (<nil>, <nil>)

	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type.
	// For example, fmt.Print takes any number of arguments of type interface{}.
	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
