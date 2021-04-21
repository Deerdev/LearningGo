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

//省略号, 参数数组
func add3(params ...int) (sum int) {
	//不能解决一个问题，我可能有不定个int值传递进来
	for _, v := range params {
		sum += v
	}
	params[0] = 9
	return
}

func main() {
	//通过省略号去动态设置多个参数值
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(add3(slice...)) //将slice打散成单个传入；因为 slice 持有的是数组引用，如果函数内部修改了 slice 的元素，会对外部的 slice 产生 side effect
	fmt.Println(slice)
	//这种效果slice
	//区别，slice是一种类型， 还是引用传递， 我们要慎重

	fmt.Println("My favorite number is", rand.Intn(10))
	// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // 先执行 0-9 的逻辑 入栈，但是
	}

	fmt.Println("done")
}
