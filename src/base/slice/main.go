package main

import "fmt"

func main() {
	mainAppend()

	//mainChange()

	//rangeSlice()

	//var h = 2
	//// region 1.切片的申明
	// 1.1 var变量申明(不会初始化)
	var s1 []int
	if s1 == nil {
		fmt.Println("slice1 is empty")
	}
	var s11 *[]int
	if s11 == nil {
		fmt.Println("slice11 is empty")
	}
	////s1[0] = 0 // panic:这个时候是nil的
	//s1 = append(s1, 1) // ok
	//s1 = append(s1, 2) // ok
	//fmt.Println(s1)

	// 1.2 make变量申明
	// The len and cap functions will both return 0 for a nil slice.
	// 翻译:nil的切片长度和容量都为0.  反过来这个说法不成立,下面是证明
	s2 := make([]int, 0)
	if s2 == nil {
		fmt.Println("slice2 is empty")
	} else {
		fmt.Println("slice2 is not empty")
	}
	s12 := make([]int, 0, 0)
	if s12 == nil {
		fmt.Println("slice12 is empty")
	} else {
		fmt.Println("slice12 is not empty")
	}

	// 1.3 冒号截取原有数组申明
	//s3 := []int{1, 2, 3, 4, 5, 6}
	//s4 := s3[:3]
	//fmt.Println(s3, s4)
	//s4[0] = 100 // s3,s4共享一段内存,所以修改对两个都有效
	//fmt.Println(s3, s4)

	// 1.4 new出来的slice
	s := new([]int)
	fmt.Println("s == nil:", s == nil)
	fmt.Println("*s == nil:", *s == nil)
	fmt.Printf("s type:%T,*s type:%T", s, *s)
	*s = append(*s, 1)
	fmt.Println("new int,s == ", s)

	// 	2.切片的追加
	//p := []byte{2, 3, 5}
	//p = AppendByte(p, 7, 11, 13)
	//fmt.Println(p) // []byte{2, 3, 5, 7, 11, 13}

	// 过滤器可以写匿名函数
	//s1 = Filter(s1, func(i int) bool {
	//	if i%h == 0 {
	//		return true
	//	}
	//	return false
	//})
	//fmt.Println(s1)

	b := []int{1, 2, 3, 4, 5, 6}
	b = Filter(b, func(b int) bool {
		if b%2 == 0 {
			return true
		}
		return false
	})

	res := make([]int, len(b))
	res = append(res, b...)
	fmt.Println("before b == ", b)
	fmt.Println("before res == ", res)

	res[0] = 100
	fmt.Println("after b == ", b)
	fmt.Println("after res == ", res)
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func FilterByte(s []byte, fn func(byte) bool) []byte {
	var p []byte // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
