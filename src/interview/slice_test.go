package interview

import (
	"fmt"
	"testing"
)

// 测试Slice的初始化
func TestSliceInit(t *testing.T) {
	t.Run("#SliceInit", func(t *testing.T) {
		list := new([]int)
		//list = append(list, 1) // 编译出错：first argument to append must be a slice; have list (variable of type *[]int)
		*list = append(*list, 1)
		fmt.Println(list)
	})
	// make只用于slice，map和channel的初始化，make返回类型本身，不过这三个也是引用类型，所以其返回也可以看成是指针类型
	// new用于所有类型的初始化，返回的是指向类型的指针
	// make和new分配的内存都是在堆上
}
