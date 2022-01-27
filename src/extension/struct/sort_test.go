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

	// 这些部分可以移动到people中
	heightFunc := func(p1, p2 *People) bool {
		return p1.Height < p2.Height
	}
	ageASCFunc := func(p1, p2 *People) bool {
		return p1.Age < p2.Age
	}
	ageDescFunc := func(p1, p2 *People) bool {
		return p1.Age > p2.Age
	}

	By(heightFunc).Sort(people)
	fmt.Printf("By height ASC:%+v\n", people)
	By(ageDescFunc).Sort(people)
	fmt.Printf("By age DESC:%+v\n", people)
	By(ageASCFunc).Sort(people)
	fmt.Printf("By age ASC:%+v\n", people)
}
