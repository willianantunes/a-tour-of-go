package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Exercise: rot13Reader
A common pattern is an io.Reader[1] that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader[2] function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader
that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the
rot13[3] substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

[1] https://pkg.go.dev/io#Reader
[2] https://pkg.go.dev/compress/gzip#NewReader
[3] https://en.wikipedia.org/wiki/ROT13
*/

var rotateBy13Places = make(map[int32]int32)

// Very naive implementation, but it solves the problem! Haha!
func buildRotateBy13Places() {
	fillIt := func(start int32, end int32) {
		for rawValue := start; rawValue < end; rawValue++ {
			cypheredValue := rawValue + 13
			if cypheredValue > end {
				cypheredValue = (cypheredValue - end - 1)
				cypheredValue = start + cypheredValue
			}
			rotateBy13Places[cypheredValue] = rawValue // RAW 78 = 65
		}
	}
	// UPPER CASE
	upperCaseStart, upperCaseEnd := 'A', 'Z' // 65, 90
	fillIt(upperCaseStart, upperCaseEnd)
	// LOWER CASE
	lowerCaseStart, lowerCaseEnd := 'a', 'z' // 97, 122
	fillIt(lowerCaseStart, lowerCaseEnd)
}

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(buffer []byte) (n int, err error) {
	content := reader.r

	// Grabbing all values
	for {
		// Buffer receives the data, according to buffer size
		numberOfBytes, err := content.Read(buffer)
		// TSHOOT purpose
		fmt.Println(buffer[:numberOfBytes])
		if err == io.EOF {
			break
		}
	}
	// Now translating them
	for index, value := range buffer {
		valueFromBufferAsByte := int32(value)
		trueValue, ok := rotateBy13Places[valueFromBufferAsByte]
		if ok {
			buffer[index] = byte(trueValue)
		}
	}
	// Done!
	return len(buffer), io.EOF
}

func main() {
	buildRotateBy13Places()
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // You cracked the code!
}
