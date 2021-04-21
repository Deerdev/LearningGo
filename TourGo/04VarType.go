package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
Go 的基本类型有

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128

*** int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
*** 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
*/

/// 零值
// 没有明确初始值的变量声明会被赋予它们的 零值。
/*
零值是：

数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）。
*/

// 可以使用组合的形式定义变量
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

/// 数值常量是高精度的 值。
// 一个未指定类型的常量由上下文来决定其类型。
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

/// iota
// iota 只能在常量组中使用，常量组内每次执行后 iota 自动+1（没行）；不同的常量组互不干扰
// x=0, y=1, z=2
const (
	x = iota // iota = 0
	y
	z
)

// 没有表达式的常量定义复用上一行的表达式(b,d 使用上一行的值)
// 从第一行开始，iota 从 0 「逐行」加
const (
	w    = iota       // iota = 0
	a    = 10         // iota = 1
	b                 // b=10, iota =2
	c    = "123"      // iota = 3
	d                 // d = "123", iota = 4
	e, f = iota, iota // e,f = 4,4
	g    = iota       // g = 5
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) // 0 0 false ""

	/// 类型转换
	// 表达式 T(v) 将值 v 转换为类型 T
	var x, y int = 3, 4
	var t float64 = math.Sqrt(float64(x*x + y*y))
	var w uint = uint(t)
	k := float32(3) // 简单写法
	fmt.Println(x, y, t, w, k)

	/// 常量
	// 使用 const 定义
	// 不可以使用 := 语法
	const Truth = true

	fmt.Println(needInt(Small))   // small 是 int
	fmt.Println(needFloat(Small)) // small 是 float64
	fmt.Println(needFloat(Big))   // big 是 float64
}
