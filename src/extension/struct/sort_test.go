package _struct

import (
	"fmt"
	"testing"
)

func TestBy_Sort(t *testing.T) {

	var people = []People{
		{"Rose", 58, 166},
		{"Daily", 78, 184},
		{"Lamia", 65, 179},
		{"Solla", 68, 177},
	}

	// 结构体不同字段排序函数的实现
	heightFunc := func(p1, p2 *People) bool {
		return p1.Height < p2.Height
	}
	ageASCFunc := func(p1, p2 *People) bool {
		return p1.Age < p2.Age
	}
	ageDescFunc := func(p1, p2 *People) bool {
		return p1.Age > p2.Age
	}

	By(func(p1, p2 *People) bool {
		return p1.Name > p2.Name
	}).Sort(people)
	fmt.Printf("By name DESC:%+v\n", people)
	By(heightFunc).Sort(people)
	fmt.Printf("By height ASC:%+v\n", people)
	By(ageDescFunc).Sort(people)
	fmt.Printf("By age DESC:%+v\n", people)
	By(ageASCFunc).Sort(people)
	fmt.Printf("By age ASC:%+v\n", people)
}
