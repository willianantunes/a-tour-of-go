package main

import "fmt"

type I interface {
	M()
}

type T struct {
	// 1 - A type implements an interface by implementing its methods.
	// There is no explicit declaration of intent, no "implements" keyword.
	// 2 - Implicit interfaces decouple the definition of an interface from its implementation,
	// which could then appear in any package without prearrangement.
	S string
}

// This method means type T implements the interface I, but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var x I = T{"hello"}
	var p T = T{"hello"}
	var t = T{"hello"}
	x.M() // hello
	p.M() // hello
	t.M() // hello
}
