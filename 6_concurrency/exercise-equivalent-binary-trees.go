package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

/*
------- Exercise: Equivalent Binary Trees
There can be many different binary trees with the same sequence of values stored in it.
For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

A function to check whether two binary trees store the same sequence is quite complex in most languages.
We'll use Go's concurrency and channels to write a simple solution.

This example uses the tree package, which defines the type:

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:
	> go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

The documentation for Tree can be found here[1].

[1] https://pkg.go.dev/golang.org/x/tour/tree?utm_source=godoc#Tree
*/

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, channel chan int) {
	if t != nil {
		Walk(t.Left, channel)
		channel <- t.Value
		Walk(t.Right, channel)
	}
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1Channel, t2Channel := make(chan int), make(chan int)

	go Walk(t1, t1Channel)
	go Walk(t2, t2Channel)

	for i := 0; i < 10; i++ {
		valueFromT1Channel := <-t1Channel
		valueFromT2Channel := <-t2Channel
		if valueFromT1Channel != valueFromT2Channel {
			return false
		}
	}

	return true
}

// https://tour.golang.org/concurrency/8
func main() {
	sampleTree := tree.New(1)
	fmt.Println(sampleTree) //((((1 (2)) 3 (4)) 5 ((6) 7 ((8) 9))) 10)

	honestChannel := make(chan int)
	go Walk(tree.New(1), honestChannel)

	for i := 0; i < 10; i++ {
		whatWhatSentThroughWalkFunction := <-honestChannel
		fmt.Println(whatWhatSentThroughWalkFunction)
	}

	shouldBeTrue := Same(tree.New(1), tree.New(1))
	shouldBeFalse := Same(tree.New(1), tree.New(2))
	if shouldBeTrue == false || shouldBeFalse == true {
		panic("Wrong results!")
	}
}
