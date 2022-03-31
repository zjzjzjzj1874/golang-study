package base

import (
	"fmt"
	"testing"
)

func TestFmtStruct_String(t *testing.T) {
	t.Run("format print with String()", func(t *testing.T) {
		fs := &FmtStruct{
			Name: "Chou Jay",
			Age:  18,
		}
		// 注意:这里打印输出,需要结构体实现一个String()方法,然后用指针来打印
		fmt.Printf("%v\n", fs)
		fmt.Printf("%+v\n", fs)
		fmt.Printf("%#v\n", fs)
	})

	t.Run("format print", func(t *testing.T) {
		fs := FmtStruct{
			Name: "Chou Jay",
			Age:  18,
		}
		fmt.Printf("%v\n", &fs) // 可以调用String()方法
		fmt.Printf("%+v\n", fs) // 不能调用String()方法
		fmt.Printf("%#v\n", fs)
	})
}
