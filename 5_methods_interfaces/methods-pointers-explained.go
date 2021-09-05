package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Here we see the Abs and Scale methods rewritten as functions.

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Again, try removing the * and test the output!
// An error will be thrown: cannot use &v (type *Vertex) as type Vertex in argument to Scale
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v)) // 50
}
