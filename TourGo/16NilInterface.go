package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	// 便接口内的具体值为 nil，方法仍然会被 nil 接收者调用
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t	// 保存了 nil 具体值的接口其自身并不为 nil。 (<nil>, *main.T)
	describe(i)	// (<nil>, *main.T)  底层值为 nil 的接口值 依然可以调用接口方法
	i.M()	// <nil>

	i = &T{"hello"}
	describe(i)
	i.M()


	/// nil 接口值
	// nil 接口值既不保存值也不保存具体类型
	// 为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型
	var w I
	describe(w)	// (<nil>, <nil>)
	w.M()	// panic: runtime error: invalid memory address or nil pointer dereference
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}