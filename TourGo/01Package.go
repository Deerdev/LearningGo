// https://github.com/hyper0x/Golang_Puzzlers


// 每个 Go 程序都是由包构成的。
package main

// 导入包, 用圆括号组合导入
import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	/// 包导出：只有 大写名称 开头的方法变量才能被导出，被包外部使用
	// math.pi 调用失败
	// fmt.Println(math.pi)
	// math.Pi 可以调用
	fmt.Println(math.Pi)
}