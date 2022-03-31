package base

import "fmt"

type FmtStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (f *FmtStruct) String() string {
	return fmt.Sprintf("Person Name:%s,age:%d\n", f.Name, f.Age)
}
