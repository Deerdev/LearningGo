package main

import (
	"fmt"
	"math"
)

// 函数也是值。它们可以像其它值一样传递。
// 函数值可以用作函数的参数或返回值。
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

/// 闭包
// 闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起
func adder() func(int) int {
	sum := 0
	// 函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))



	// pos neg 都绑定了一个sum（分开的）
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}





