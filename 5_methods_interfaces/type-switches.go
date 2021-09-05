package main

import "fmt"

func do(i interface{}) {
	// A type switch is a construct that permits several type assertions in series.
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	// In the default case (where there is no match), the variable v is of the same interface type and value as i
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	// I think that justifies the use of generics to have compile time errors, instead of runtime ones
	do(21)      // Twice 21 is 42
	do("hello") // "hello" is 5 bytes long
	do(true)    // I don't know about type bool!
}
