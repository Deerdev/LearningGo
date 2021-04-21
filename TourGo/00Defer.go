package main

import (
	"fmt"
)

/// defer
// defer 语句会将函数推迟到外层函数返回之后执行
// 如果有多个defer会出现什么情况 多个defer是按照先入后出的顺序执行
func test() {
	defer fmt.Println("world") // 推迟调用的函数其参数会立即求值, 但直到外层函数返回前该函数都不会被调用。

	fmt.Println("hello")
}

func f1() int {
	x := 10
	defer func() {
		x++
	}()
	tmp := x //x是int类型 值传递
	return tmp
}

func f3() {
	x := 10
	defer func(a int) {
		// 此处打印的是 全局的 x: 11
		fmt.Println(x)
		// 参数 a 在 defer 定义时捕获：10
		fmt.Println(a)
	}(x)
	x++
}

func f2() *int {
	a := 10
	b := &a
	defer func() {
		*b++
	}()
	temp_data := b
	return temp_data
}

func main() {
	// 值传递和引用传递的 捕获
	fmt.Println(f1())  // 10
	fmt.Println(*f2()) // 11

	//defer本质上是注册了一个延迟函数，defer函数的执行顺序已经确定
}
