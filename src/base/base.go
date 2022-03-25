package base

import "fmt"

// 数组相关
func array() {
	arr := [4]int{1}
	arr1 := [...]int{1, 0, 0, 0}

	fmt.Println(arr, arr1, len(arr), len(arr1))
	fmt.Println(arr == arr1)
}
