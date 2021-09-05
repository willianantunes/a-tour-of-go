package main

import (
	"fmt"
	"math"
)

/*
- Interfaces

Note: There is an error in the example code on line 30.
Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
*/

// An interface type is defined as a set of method signatures.
type Abser interface {
	// A value of interface type can hold any value that implements those methods.
	Abs() float64
}

// https://tour.golang.org/methods/9
func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex) and does NOT implement Abser.
	a = v // Vertex does not implement Abser (Abs method has pointer receiver)

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
