package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string 和 int 互转
	{
		a := 16
		s := strconv.Itoa(a)
		fmt.Printf("%s\n", s)

		s1 := "15"
		a1, err := strconv.Atoi(s1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d\n", a1)
	}
	// string 和其他类型转换
	/*
		strconv.ParseBool
		strconv.ParseFload
		strconv.ParseInt
		strconv.ParseUint
	 */
	{
		// base: 进制；bitSize: 存储大小 32、64
		a, _ := strconv.ParseInt("100", 10, 64)
		fmt.Println(a)
		b, _ := strconv.ParseBool("true")
		fmt.Println(b)
	}
}
