package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 值接收者
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/// 使用指针作为接收者
// 指针接收者的方法可以修改接收者指向的值
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/// 指针接收者 和 值接收者 区别
// 值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作, 对于函数的值参数也是如此。
// 指针接收者 可以更改 main 函数中声明的 Vertex 的值。
func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
