package main

import "fmt"

// A function can take zero or more arguments.
// In this example, add takes two parameters of type int.
// Notice that the type comes after the variable name.
// https://go.dev/blog/declaration-syntax

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
