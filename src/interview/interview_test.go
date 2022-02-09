package interview

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("#结构体比较", func(t *testing.T) {
		sn1 := struct {
			age  int
			name string
		}{age: 11, name: "qq"}
		sn2 := struct {
			age  int
			name string
		}{age: 11, name: "11"}

		if sn1 == sn2 {
			fmt.Println("sn1 == sn2")
		}
		psn1 := struct {
			age  int
			name string
		}{age: 11, name: "qq"}
		psn2 := struct {
			age  int
			name string
			// 如果age和name调换一个顺序,则不能再进行比较
		}{name: "11", age: 11}

		if psn1 == psn2 {
			fmt.Println("sn1 == sn2")
		}

		//sm1 := struct {
		//	age int
		//	m   map[string]string
		//}{age: 11, m: map[string]string{"a": "1"}}
		//sm2 := struct {
		//	age int
		//	m   map[string]string
		//}{age: 11, m: map[string]string{"a": "1"}}

		//if sm1 == sm2 { // 编译不能通过
		//	fmt.Println("sm1 == sm2")
		//}

		// 结构体的几个注意事项:
		// 1.结构体只能比较是否相等，但是不能比较大小；
		// 2.想同类型的结构体才能进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关；
		// 3.如果struct的所有成员都可以比较，则该struct就可以通过==或!=进行比较是否相同，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
	})
}
