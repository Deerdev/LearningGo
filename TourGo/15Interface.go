package main

import (
	"fmt"
	"math"
)

// 接口类型 是由一组方法签名定义的集合
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/// 接口的隐式实现
// 某类型只要实现了接口的所有方法 即实现了该接口，而不需要显式的去 implement 接口
// 隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备
type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

/// 接口值
// 接口也是值，可以当做参数传递
// 在内部，接口值可以看做包含值和具体类型的元组：(value, type)
// 接口值保存了一个具体底层类型的具体值。接口值调用方法时会执行其底层类型的同名方法。
type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	// 接口类型的变量 可以保存任何实现了这些方法的值。
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v // cannot use v (type Vertex) as type Abser in assignment: Vertex does not implement Abser (Abs method has pointer receiver)
	fmt.Println(a.Abs())

	// 接口值
	var i I

	i = &T{"Hello"}
	describe(i) // (&{Hello}, *main.T)
	i.M()

	i = F(math.Pi)
	describe(i) // (3.141592653589793, main.F)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
