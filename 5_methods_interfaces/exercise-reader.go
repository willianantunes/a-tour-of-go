package main

import (
	//"fmt"
	"strings"
)
import "golang.org/x/tour/reader"

/*
Exercise: Readers

Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (myReader MyReader) Read(buffer []byte) (int, error) {
	var intValueOfASCIICharacterA = 'A'
	for index, _ := range buffer {
		buffer[index] = byte(intValueOfASCIICharacterA)
	}
	return len(buffer), nil
}

func main() {
	reader.Validate(MyReader{})
}
