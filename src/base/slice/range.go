package main

import "fmt"

func rangeSlice() {
	a := []int{1, 2, 3}
	fmt.Printf("before adding 1 to elements, a = %v \n", a)
	// before adding 1 to elements, a = [1 2 3]
	for _, n := range a {
		n += 1
	}

	// slice elements haven't changed because n is a copy of slice elements.
	fmt.Printf("after adding 1 to elements, a = %v \n", a)
	// after adding 1 to elements, a = [1 2 3]

	// Fix:
	// to change that address the elements with the index and it should be changed
	for i := range a {
		a[i] += 1
	}

	fmt.Printf("after adding 1 to elements with index, a = %v \n", a)
	// after adding 1 to elements with index, a = [2 3 4]
}
