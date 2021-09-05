package main

import "golang.org/x/tour/pic"

/*
Exercise: Slices

Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/

func Pic(dx, dy int) [][]uint8 {
	// fmt.Println("Received values:", dx, dy) // Received values: 256 256
	// It should return a slice of length dy
	// each element of which is a slice of dx 8-bit unsigned integers
	matrixWhichRepresentsAnImage := make([][]uint8, dy)
	for y := 0; y < len(matrixWhichRepresentsAnImage); y++ {
		var line []uint8
		for x := 0; x < dx; x++ {
			// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
			//interestingThing := func() int { return x * y }
			//interestingThing := func() int { return (x + y) / 2 }
			interestingThing := func() int { return x ^ y }
			line = append(line, uint8(interestingThing()))
		}
		matrixWhichRepresentsAnImage[y] = line
	}
	// fmt.Println("What was built:", matrixWhichRepresentsAnImage)
	return matrixWhichRepresentsAnImage
}

func main() {
	// You should execute it here: https://tour.golang.org/moretypes/18
	pic.Show(Pic)
}
