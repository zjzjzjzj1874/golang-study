package main

import "fmt"

func change(a []int) {
	fmt.Printf("slice's addr in func's header is:%p\n", &a)
	//a[0] = 0
	//a = []int{6, 6, 6}
	a = a[:]
	a[0] = 0
	fmt.Printf("slice's addr in func's tail is:%p\n", &a)
}

func Assign1(s []int) {
	s = []int{6, 6, 6} // 这里会返回一个新的slice
	s[0] = 0           // 所以下面修改新slice的元素,对原slice也没有影响
}

func mainChange() {
	s := []int{1, 2, 3, 4, 5, 6}
	Assign1(s)
	fmt.Println(s) // (1)

	a := []int{1, 2, 3}
	fmt.Printf("slice's addr in main is:%p\n", &a)
	change(a)
	fmt.Println(a)
}
