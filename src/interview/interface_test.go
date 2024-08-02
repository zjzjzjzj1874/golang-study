package interview

import (
	"testing"
)

type Animal interface { // 申明一个Animal的interface类型
	Speak(string) string
}
type Monkey struct{}

func (s *Monkey) Speak(think string) (talk string) {
	if think == "love" {
		talk = "you are a good monkey"
	} else {
		talk = "Hi"
	}
	return
}

// 测试interface的赋值
func TestInterfaceAssignment(t *testing.T) {
	t.Run("#InterfaceAssignment", func(t *testing.T) {
		// cannot use Monkey{} (value of type Monkey) as Animal value in variable declaration: Monkey does not implement Animal (method Speak has pointer receiver)
		//var an Animal = Monkey{}
		var an Animal = &Monkey{}
		think := "love"
		an.Speak(think)
	})
	//	在golang中对多态的特点体现从语法上并不是很明显。
	//我们知道发生多态的几个要素：
	//1、有interface接口，并且有接口定义的方法。
	//2、有子类去重写interface的接口。
	//3、有父类指针指向子类的具体对象
	//
	//那么，满足上述3个条件，就可以产生多态效果，就是，父类指针可以调用子类的具体方法。
	//所以上述代码报错的地方在var an Animal = Monkey{}这条语句， Monkey{}已经重写了父类Animal{}中的Speak(string) string方法，那么只需要用父类指针指向子类对象即可。
	//所以应该改成var an Animal = &Monkey{} 即可编译通过。（Animal为interface类型，就是指针类型）
}

type S struct{}

func f(x interface{})  {}
func g(x *interface{}) {}

// interface和*interface
func TestInterfaceAndPtrInterface(t *testing.T) {
	t.Run("#PtrInterface", func(t *testing.T) {
		s := S{}
		p := &s

		// ABCD中有哪些有问题
		f(s) // A
		//g(s) // B  //cannot use s (variable of type S) as *interface{} value in argument to g: S does not implement *interface{} (type *interface{} is pointer to interface, not interface)
		f(p) // C
		//g(p) // D // cannot use p (variable of type *S) as *interface{} value in argument to g: *S does not implement *interface{} (type *interface{} is pointer to interface, not interface)

		// 更改
		var sInterface interface{} = s
		g(&sInterface) // B
		var pInterface interface{} = p
		g(&pInterface) // D
	})
	// 总结：
	// interface{}
	//	描述：这是一个空接口类型，可以持有任何类型的值。
	//	用法：用于定义可以接受任意类型值的参数或变量。
	// *interface{}
	//	描述：这是一个指向空接口类型的指针。
	//	用法：用于传递空接口类型的指针，允许修改原始接口值。
	// 区别1-传递方式
	//	interface传递值的副本
	//	*interface传递指针，可以修改原始值
	// 区别2-灵活性
	//	interface非常灵活，很常用，用于函数参数时表示接受任意类型值，万能类型
	//	*interface用的很少，主要在需要修改接口值的情况下使用
}
