package main

import "fmt"

// A struct literal denotes a newly allocated struct value by listing the values of its fields.
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // has type Vertex
	// You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // X:0 and Y:0
	// The special prefix & returns a pointer to the struct value.
	p = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)      // {1 2} &{1 2} {1 0} {0 0}
	fmt.Println(v1, *p, v2, v3)     // {1 2} {1 2} {1 0} {0 0}
	fmt.Println(&v1, &p, &v2, &v3)  // &{1 2} 0x51f610 &{1 0} &{0 0}
	fmt.Println(&v1, &p, &v2, *&v3) // &{1 2} 0x51f610 &{1 0} {0 0}
}
