package main

import (
	"fmt"
)

/*
- Exercise: Errors
Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type
	> type ErrNegativeSqrt float64
and make it an error by giving it a
	> func (e ErrNegativeSqrt) Error() string
method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
*/

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//return fmt.Sprintf("cannot Sqrt negative number: %g", e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	invalidInput := x < 0
	if invalidInput {
		// https://golang.org/ref/spec#Handling_panics
		return 0, ErrNegativeSqrt(x)
	}
	guess := 1.0 // I could have done also: guess := float64(1)
	currentGuess := guess
	for {
		currentGuess -= (guess*guess - x) / (2 * guess)

		areGuessesEqual := currentGuess == guess
		areGuessesPrettyClose := (currentGuess - guess) < -0.00000000000001

		if areGuessesEqual || areGuessesPrettyClose {
			break
		}
		guess = currentGuess
		//fmt.Println("Current guess:", guess)
	}
	return guess, nil
	//return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
