package main

import "fmt"

// Struct fields can be accessed through a struct pointer.
// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
// However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v) // {1000000000 2}
	fmt.Println(p) // &{1000000000 2}
}
