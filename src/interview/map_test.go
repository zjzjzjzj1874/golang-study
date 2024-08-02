package interview

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//var sMap map[string]Student

// 测试map的赋值过程
func TestMapAssignment(t *testing.T) {
	t.Run("#ErrorAssignment", func(t *testing.T) {
		list := make(map[string]Student) // 初始化
		student := Student{Name: "Alice"}

		list["student"] = student // 值拷贝的过程
		// list["student"]是一个值引用，值引用的特点是`只读`，所以`list["student"].Name = "bob"`不允许修改
		//list["student"].Name = "bob"
		fmt.Println(list["student"])
	})

	//关于golang中map的这种古怪的特性有这样几个观点：
	//1）map作为一个封装好的数据结构，由于它底层可能会由于数据扩张而进行迁移，所以拒绝直接寻址，避免产生野指针；
	//2）map中的key在不存在的时候，赋值语句其实会进行新的k-v值的插入，所以拒绝直接寻址结构体内的字段，以防结构体不存在的时候可能造成的错误；
	//3）这可能和map的并发不安全性相关

	//x = y 这种赋值的方式，你必须知道 x的地址，然后才能把值 y 赋给 x。
	//但 go 中的 map 的 value 本身是不可寻址的，因为 map 的扩容的时候，可能要做 key/val pair迁移
	//value 本身地址是会改变的
	//不支持寻址的话又怎么能赋值呢
	t.Run("#Assignment1", func(t *testing.T) {
		list := make(map[string]Student) // 初始化
		student := Student{Name: "Alice"}
		list["student"] = student // 值拷贝的过程
		//list["student"].Name = "bob"
		// list["student"]是一个值引用，值引用的特点是`只读`，所以`list["student"].Name = "bob"`不允许修改

		// 方法一
		tmpStu := list["student"]
		tmpStu.Name = "bob"
		list["student"] = tmpStu
		// 解释：先做一次值拷贝，做出一个tmpStu的副本，然后修改副本，最后再通过值拷贝复制回去。
		// 性能很差，整个过程有两次结构体的值拷贝
		fmt.Println(list["student"])
	})
	t.Run("#Assignment2", func(t *testing.T) {
		list := make(map[string]*Student) // 初始化
		student := Student{Name: "Alice"}
		list["student"] = &student // 引用指针
		// 我们把list的value从值类型改成指针类型。这样我们实际上每次修改的都是指针指向的student空间；
		// 指针本身是常指针，不能修改，只读，但指向的Student是可以随意修改，这里也不会发生值拷贝，只是一个指针的赋值。
		list["student"].Name = "bob"
		fmt.Println(list["student"])
	})
}

// map的遍历
func TestIterate(t *testing.T) {
	t.Run("#ErrorIterate", func(t *testing.T) {
		m := make(map[string]*Student)
		sts := []Student{
			{Name: "Alice", Age: 18},
			{Name: "Bob", Age: 20},
			{Name: "Chris", Age: 19},
		}

		for _, stu := range sts { // 因为range在迭代的时候，本身是一个值拷贝, &stu最后指向的是同一个地址，导致复制出错。
			m[stu.Name] = &stu
		}
		for name, stu := range m {
			fmt.Println(name, "==>", stu)
		}
		//	Chris ==> &{Chris 19}
		//	Alice ==> &{Chris 19}
		//	Bob ==> &{Chris 19}
		// 分析：最后stu指向了同一个副本；因为range在迭代的时候，本身是一个值拷贝&stu最后指向的是同一个地址，导致复制出错。
	})
	t.Run("#Iterate", func(t *testing.T) {
		m := make(map[string]*Student)
		sts := []Student{
			{Name: "Alice", Age: 18},
			{Name: "Bob", Age: 19},
			{Name: "Chris", Age: 20},
		}

		for i := 0; i < len(sts); i++ {
			m[sts[i].Name] = &sts[i]
		}

		for name, stu := range m {
			fmt.Println(name, "==>", stu)
		}
		// Alice ==> &{Alice 18}
		// Bob ==> &{Bob 19}
		// Chris ==> &{Chris 20}
	})
}
