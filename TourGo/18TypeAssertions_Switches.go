package main

import "fmt"

/// 类型断言
// 类型断言 提供了访问接口值底层具体值的方式。

// t := i.(T)  
//断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t
// 若 i 并未保存 T 类型的值，该语句就会触发一个恐慌(panic)

// t, ok := i.(T)
// 若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。
// 否则，ok 将为 false 而 t 将为 T 类型的零值，程序并 不会产生恐慌


/// 
// 类型选择 是一种按顺序从几个类型断言中选择分支的结构。
// 类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）
/*
switch v := i.(type) {
case T:
    // v 的类型为 T
case S:
    // v 的类型为 S
default:
    // 没有匹配，v 与 i 的类型相同
}
*/
// 在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值。
// 在默认 default（即没有匹配）的情况下，变量 v 与 i 的接口类型和值相同
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)



	do(21)
	do("hello")
	do(true)
}