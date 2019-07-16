package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// go 没有类，但是有方法
// 方法 主要给 结构体 额外添加的方法，结构体可以直接调用

// 方法就是一类带特殊的 接收者 参数的函数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 记住：方法只是个带接收者参数的函数
// 函数形式：
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/// 也可以为非结构体类型声明方法。
// 只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
// 就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法

// 把基础类型alias一下成为包内的类型
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
