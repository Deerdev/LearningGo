package main

import "fmt"

type Vertex struct {
	X, Y float64
}

/// 方法与指针重定向
/// 方法 和 函数 指针接收者 和 指针参数 的区别
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/// 函数传递 指针参数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/// 何时使用指针接收者
// 使用指针接收者的原因有二：
// 首先，方法能够修改其接收者指向的值。
// 其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

func main() {
	v := Vertex{3, 4}
	/// 以指针为接收者的方法被调用时，接收者既能为值又能为指针
	// 该方法不区分 接收者是否是指针，可以直接调用
	v.Scale(2)

	/// 带指针参数的函数必须接受一个指针
	// v是值，所以需要传入指针
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	// 方法不区分 接收者是否是指针，可以直接调用
	p.Scale(3)
	// p是指针，可以直接传入
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
