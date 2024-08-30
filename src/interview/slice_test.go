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

// Golang中传递是值传递，切片时引用类型，如果修改切片内的元素，也会导致原始切片受影响；但是如果是append扩容后，对扩容后的切片进行操作，则不会修改原来的值
func TestSlicePassValue(t *testing.T) {
	t.Run("#TestSlice", func(t *testing.T) {
		nums := make([]int, 0, 4)
		for i := 0; i < 4; i++ {
			nums = append(nums, i*100)
		}
		fmt.Println("nums === ", nums)
		appendSlice(nums)
		fmt.Println("after append nums === ", nums)

		modifySliceValue(nums)
		fmt.Println("after modify nums === ", nums)

	})
}

// 修改切片的元素,倒个序
func modifySliceValue(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func appendSlice(nums []int) {
	nums = append(nums, []int{1, 2, 3, 4}...)
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
	fmt.Println("in append: nums = ", nums)
}
