package main

import (
	"fmt"
	"strings"
)

func main() {
	/// make
	// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
	a := make([]int, 5) // len(a)=5
	printSlice2("a", a)

	// 要指定它的容量，需向 make 传入第三个参数
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice2("b", b)
	// b = b[:cap(b)] // len(b)=5, cap(b)=5
	// b = b[1:]      // len(b)=4, cap(b)=4

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)


	/// 切片的切片
	// 创建一个井字板（经典游戏）, 切片包含切片
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	/// 追加元素到切片 append
	// 当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)


	/// Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份【副本】
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// 忽略不需要的值
	// for i, _ := range pow
	// for _, value := range pow
	// for i := range pow
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}