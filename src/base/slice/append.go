package main

import "fmt"

func mainAppend() {
	a := []int{1, 2, 3}
	fmt.Printf("before append, a = %v\n", a)
	myAppend(a, 4)
	fmt.Printf("after append, a = %v\n", a) // after append, a = [1 2 3]

	myAppend2(&a, 4)
	fmt.Printf("after append with pointer, a = %v\n", a)

	a = myAppend3(a, 5)
	fmt.Printf("after append with return, a = %v\n", a)
}

// 值传递,append只会对新的传入的值修改,原来的slice也没有影响
func myAppend(a []int, i int) {
	a = append(a, i) //
}

func myAppend2(a *[]int, i int) {
	*a = append(*a, i)
}

func myAppend3(a []int, i int) []int {
	a = append(a, i)

	return a
}
