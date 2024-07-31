package interview

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	// 1. 同一个调用中，defer是FILO，先进后出的栈
	f1 := func() {
		fmt.Println("f1-A")
	}
	f2 := func() {
		fmt.Println("f2-B")
	}
	f3 := func() {
		fmt.Println("f3-C")
	}
	t.Run("#Defer1", func(t *testing.T) {
		defer f1()
		defer f2()
		defer f3()
		// 打印顺序为：f3-C  f2-B  f1-A
	})

	deferFunc := func() {
		fmt.Println("defer call")
	}
	returnFunc := func() int {
		fmt.Println("return call")
		return 0
	}
	returnAndDefer := func() int {
		// 后执行
		defer deferFunc()
		// 先执行
		return returnFunc()
	}
	// 2.defer和return的关系：return之后的语句先执行，defer后的语句后执行
	t.Run("#Defer2", func(t *testing.T) {
		returnAndDefer()

		// 打印：
		// return call
		// defer call
	})

	// 3.defer与无名返回值
	test := func() int {
		var i int
		defer func() {
			i++
			// 闭包引用，会在defer函数执行时根据上下文确定当前的值 i=2
			fmt.Println("defer1-", i)
		}()

		defer func() {
			i++
			// 闭包引用，会在defer函数执行时根据上下文确定当前的值 i=1
			fmt.Println("defer2-", i)
		}()

		// 先执行return i，把i的值给到一个临时变量，作为函数值返回，即返回的函数值不再会根据defer中的东西变化 i = 0
		return i
	}
	t.Run("#Defer3", func(t *testing.T) {
		fmt.Println("test:", test())
		// 所以打印：执行顺序：return defer2 defer1 => return的时候，因为返回参数匿名，所以用临时变量tmp存i=0，后续return的tmp将不再变化；与下面的命名返回不同！！！
		// defer2-1
		// defer1-2
		// test:0
	})

	// 4. defer返回函数中命名返回值
	t.Run("#Defer4", func(t *testing.T) {
		test := func() (i int) {
			defer func() {
				i++
				fmt.Println("defer1-", i)
			}()
			defer func() {
				i++
				fmt.Println("defer2-", i)
			}()

			return i
		}

		fmt.Println("test-", test())
		// 分析：执行顺序return => defer1 => defer2；但是defer1和defer2会改变i的值，返回值又是i，不需要临时变量存储，最终返回的test值为i=2
		// 打印：
		// defer2-1
		// defer1-2
		// test-2
	})

	// 5.defer meet panic
	t.Run("#Defer5 Meet Panic", func(t *testing.T) {
		test := func() {
			defer func() {
				fmt.Println("defer0 before panic")
			}()
			defer func() {
				fmt.Println("defer1 before panic")
				if err := recover(); err != nil {
					fmt.Println("panic reason:", err)
				}
			}()
			defer func() {
				fmt.Println("defer2 before panic")
			}()

			panic("异常啦")

			defer func() {
				fmt.Println("defer after panic")
			}()
		}

		test()
		fmt.Println("end test case")

		// 分析：panic和return类似，在捕获recover之后，也一样先进后出，但是panic之后的不会再进入
		// 打印：
		// defer2 before panic
		// defer1 before panic
		// panic reason: 异常啦
		// defer0 before panic
		// end test case
	})

	// 6.defer遇到不recover的panic
	t.Run("#Defer Meet Do not recover panic", func(t *testing.T) {
		test := func() {
			defer func() { fmt.Println("defer0") }()
			defer func() { fmt.Println("defer1") }()
			panic("异常啦")
			defer func() { fmt.Println("defer2") }()
		}

		test()
		fmt.Println("end test")
		// 分析：主程序崩溃，在崩溃之前根据先进后出，打印defer的东西
		// defer1
		// defer0
		// panic......
	})

	// 7.接4 plus  考察点：(1).defer + 表达式 (2).defer+匿名函数(有无入参)-(返回值是否匿名)
	// (1).Defer+表达式，变量也会一同压入栈中，后续不会变更；
	// (2.1)Defer+匿名函数有入参：也是调用的一瞬间，入参的值固定，不会再变化；
	// (2.2)Defer+匿名函数无入参：那么内部的变量能够访问到变化后的值；
	// (2.3)Defer+匿名返回值；参考#Defer3 #Defer4的测试用例
	t.Run("#Defer7", func(t *testing.T) {
		test1 := func() (x int) {
			defer fmt.Println("test1-defer x =", x)
			x = 7
			return 8
		}
		test2 := func() (x int) {
			x = 7
			defer fmt.Println("test2-defer x =", x)
			return 9
		}
		test3 := func() (x int) {
			defer func() {
				fmt.Println("test3-defer x =", x)
			}()
			x = 7
			return 10
		}
		test4 := func() (x int) {
			defer func(n int) {
				fmt.Println("test4-defer n =", n)
				fmt.Println("test4-defer x =", x)
			}(x)

			x = 7
			return 11
		}

		fmt.Println("begin")
		fmt.Println("test1 :", test1()) // test1-defer x = 0
		fmt.Println("test2 :", test2()) // test2-defer x = 7
		fmt.Println("test3 :", test3()) // test3-defer x = 10
		fmt.Println("test4 :", test4()) // test4-defer n = 0  test4-defer x = 11
		fmt.Println("end")
	})

	// 8. 挂羊头卖狗肉
	t.Run("#Defer8", func(t *testing.T) {
		f1 := func() (x int) {
			defer func() {
				x++
			}()
			return 0
		}
		f2 := func() (x int) {
			y := 5
			defer func() {
				y += 5
			}()
			return y // 先return，再defer，return的时候把y=5赋值给x了，后续y变化不会影响x的值，返回5
		}
		f3 := func() (x int) {
			defer func(x int) {
				x = x + 5
			}(x)
			return 0 // 返回0，x=0作为参数传给defer的有参函数，后续变化只会修改入参的x
		}
		f4 := func() (x int) {
			y := 5
			defer func() {
				x = x + 5
			}()
			return y // 返回10，y=5作为参数传给defer的有参函数x, x = 5,后续x = x + 5 = 10
		}

		fmt.Println(f1()) // 1
		fmt.Println(f2()) // 5
		fmt.Println(f3()) // 0
		fmt.Println(f4()) // 10
	})
}
