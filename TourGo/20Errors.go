package main

import (
	"fmt"
	"time"
)

// Go 程序使用 error 值来表示错误状态。
//go语言认为错误就要自己处理

// 与 fmt.Stringer 类似，error 类型是一个内建接口：
/*
type error interface {
    Error() string
}
*/
// （与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error。）

// 通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。
/*
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
*/
// error 为 nil 时表示成功；非 nil 的 error 表示失败。

type MyError struct {
	When time.Time
	What string
}

// 实现error方法
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// 错误捕获
func div(a, b int) (int, error) {
	if b == 0 {
		panic("can not be zero")
	}
	return a / b, nil
}

// goroute 内部的异常只能自己捕获
func f1() {
	// 此处无法捕获 goroutine 内部的 panic
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获error")
		}
	}()

	go func() {
		// 只有在内部才能捕获自己的 panic
		//defer func() {
		//	err := recover()
		//	if err != nil {
		//		fmt.Println("goroutine 捕获error")
		//	}
		//}()

		panic("err")
	}()
	time.Sleep(10 * time.Second)
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// 使用 recover 捕获异常
	a := 12
	b := 0
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("catch error")
		}
		fmt.Println("defer finish")
	}()

	fmt.Println(div(a, b))

	// goroutine 的 panic 只能自己捕获
	f1()
}
