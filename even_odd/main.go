package main

import (
	"fmt"
)

func main() {

	// create a slice of ints 0 - 10
	myNumbers := []int{0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10}

	// iterate through the slice, check to see if a number is even/odd
	for _, number := range myNumbers {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}
