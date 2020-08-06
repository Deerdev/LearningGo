package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/// if
func sqrt(x float64) string {
	// 无括号
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if 语句可以在条件表达式前执行一个简单的语句
	// 变量 v 的作用域仅在 if 内
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func main() {
	/// for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// 前置和后置语句是可选的
	for ; sum < 1000; {
		sum += sum
	}

	// go 的 while 也是 for
	for sum < 1000 {
		sum += sum
	}

	// 无效循环
	// for {}

	/// switch
	// Go 只运行选定的 case，而非之后所有的 case (不用写break)。 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 
	// 除非以 fallthrough 语句结束，否则分支会自动终止。 
	// Go 的另一点重要的不同在于 switch 的 case 无需为【常量】，且取值不必为整数。
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	
	// case 可以执行一个函数 或 一个语句
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// 没有条件的switch, 同 if-then-else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}