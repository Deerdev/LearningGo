// 每个 Go 程序都是由包构成的。
package main

// 导入包, 用圆括号组合导入
import (
	"fmt"
	"math/rand"
)

/// 函数
func add(x int, y int) int {
	return x + y
}

// 类型相同时，可以省略前面的类型
func add2(x, y int) int {
	return x + y
}

// 函数可以返回任意多个值
func swap(x, y string) (string, string) {
	return y, x
}

/// 命名返回值
// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// 没有参数的 return 语句返回已命名的返回值
	return
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)  // 先执行 0-9 的逻辑 入栈，但是
	}

	fmt.Println("done")
}


/// defer 
// defer 语句会将函数推迟到外层函数返回之后执行
func test() {
	defer fmt.Println("world")  // 推迟调用的函数其参数会立即求值, 但直到外层函数返回前该函数都不会被调用。

	fmt.Println("hello")
}