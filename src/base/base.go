package base

// 闭包函数: 作用：缩小变量作用域，减少对全局变量的污染 => 不过以下的sum在编译阶段,会逃逸到堆内存上
func closure() func(int) int {
	var sum = 0
	return func(a int) int {
		sum += a
		return sum
	}
}
