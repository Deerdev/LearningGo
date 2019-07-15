package main

import (
	"fmt"
	"math/rand"
	"math"
)

var c, python, java bool
// 赋初值
var i, j int = 1, 2

func main() {
	// var 声明变量，类型在后面
	var i int
	fmt.Println(i, c, python, java)

	/// 赋初值后，可以不加类型，通过值类型推断
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	/// 短变量声明，***只能在函数内使用***
	k := 3	// 同 var k = 3 、var k int = 3
}