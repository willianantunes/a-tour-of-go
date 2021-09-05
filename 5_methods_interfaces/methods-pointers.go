package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
// SIMPLY PUT: Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	// Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.
	fmt.Println(v.Abs()) // 50
	// After removing * from the method: fmt.Println(v.Abs()) // 5
}
