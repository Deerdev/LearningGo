package main

import (
	"fmt"
	"unsafe"
)

/// 指针
// 和C一样，但是 GO 的指针没有运算操作, 如 p++ 等
func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值

	// 数组指针，数组的指针类型，需要明确指明长度(长度也是类型的一部分)：[3]int
	//指针还可以指向数组 指向数组的指针 数组是值类型
	arr := [3]int{1,2,3}
	var ip *[3]int = &arr
	//指针数组
	var ptrs [3]*int //创建能够存放三个指针变量的数组
	// c和c++中指针的功能很强大 指针的转换 指针的偏移 指针的运算
	// go语言没有屏蔽指针，但是做了大量的限制，安全性高，相比 c和c++灵活性降低


	/// ------ make和new函数 ------
	/// new函数用法
	var p0 *int //申明了一个变量p 但是变量没有初始值 没有内存
	*p0 = 10    // 赋值会 crash，因为 *p0 没有内存 (p0 指针本身是有存储，但是没有指向的内存，指向为 nil)

	//默认值 int byte rune float bool string 这些类型都有默认值
	//指针、切片、map， 接口这些默认值是nil
	var a int   // 值类型，默认会有内存
	a = 10
	fmt.Println(a)

	// 使用 new 为 int 申请内存
	var p1 *int = new(int) //go的编译器就知道先申请一个内存空间，这里的内存中的值全部设置为0
	*p1 = 10

	/// make
	//int使用make就不行
	//除了 new 可以申请内存以外 还有一个函数就是make，更加常用的是make函数，make函数一般用于切片、map
	//new函数返回的是这个值的"地址" 指针 make函数返回的是指定类型的"实例"
	var info map[string]string = make(map[string]string)
	info["c"] = "bobby"

	/// nil的一些细节
	var info2 map[string]string
	if info2 == nil {
		fmt.Println("map的默认值是 nil")
	}

	var slice []string
	if slice == nil {
		fmt.Println("slice的默认值是 nil")
	}

	var err error
	if err == nil {
		fmt.Println("error的默认值是 nil")
	}

	//python中的None和go语言中的nil类型不一样，None是全局唯一的
	//go语言中 nil是唯一可以用来表示部分类型的零值的标识符， 它可以代表许多不同内存布局的值
	// unsafe.Sizeof(x): 计算 x 指针所指向的类型占用的空间（假象 type x = new(type)）
	// slice 初始化空间 24，map 初始化空间 8
	/*
	sizeof(slice) = 8 + 8 + 8
	type slice struct {
		array unsafe.Pointer // 元素指针
		len   int // 长度
		cap   int // 容量
	}
	*/
	fmt.Println(unsafe.Sizeof(slice), unsafe.Sizeof(info2))     // 24  8
}