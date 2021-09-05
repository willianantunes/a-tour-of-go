package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful[1].

[1] https://pkg.go.dev/strings#Fields
*/

func WordCount(s string) map[string]int {
	// I am learning Go!
	// The quick brown fox jumped over the lazy dog.
	// I ate a donut. Then I ate another donut.
	// A man a plan a canal panama.
	fmt.Println("Received string:", s)
	splitWords := strings.Fields(s)
	fmt.Printf("Fields are: %q\n", splitWords) // ["I" "am" "learning" "Go!"]
	wordCountMap := make(map[string]int)
	for _, word := range splitWords {
		counter := wordCountMap[word]
		wordCountMap[word] = counter + 1
	}
	return wordCountMap
}

func main() {
	// You should execute it here: https://tour.golang.org/moretypes/23
	wc.Test(WordCount)
}
