package main

import "fmt"

func main() {
    /// ---- 数组 ------
    // 需要指明数组的大小和存储的数据类型。
	var a [2]string // 默认初始化为 ["", ""]
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 直接初始化
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// 自动计算个数；不指定"..."是切片的初始化
	b := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    fmt.Println(b)

    // 指定某些索引的数值
    f := [...] int{0: 1, 4: 1, 9: 1}
    fmt.Println(f) // [1 0 0 0 1 0 0 0 0 1]

    // 数组长度
    fmt.Println(len(f))

    // range 遍历
    for i, v := range b {
        // index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
    }

    // 数组是值类型，并且数组的大小是类型的一部分，长度不同数组，类型也不一样：[3]int 和 [5]int 是不同的类型
    var al []int     // 创建slice
    sl:=[3]int{1,2,3} //创建有初始化元素的slice
    s2:=[4]int{1,2,3,4} //创建有初始化元素的slice
    fmt.Printf("%T, %T, %T", al, sl, s2)    // []int, [3]int, [4]int





    /// ------ 切片 ------
	/// 切片 a[low : high]
	var ws []int = primes[1:4]  // 提取 1-2-3， [1:4)
	fmt.Println(ws)	// [3 5 7]
	/// 切片就像数组的引用, 和原始数组公用一个值
	// 切片并不存储任何数据，它只是描述了底层数组中的一段。
	// 更改切片的元素会修改其底层数组中对应的元素。
	// 与它共享底层数组的切片都会观测到这些修改
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	aa := names[0:2]
	bb := names[1:3]
	fmt.Println(aa, bb) // [John Paul] [Paul George]

	bb[0] = "XXX"
	fmt.Println(aa, bb) // [John XXX] [XXX George]
	fmt.Println(names)	// [John XXX George Ringo]

	// 切片定义：切片定义类似于没有长度的数组定义
	// 先定义个数组，然后构建一个引用这个数组的切片
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	w := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(w)

	// 切片的默认定义: low 默认为 0, high 默认为数组长度
	/*
	以下等价：
	a[0:10]
	a[:10]
	a[0:]
	a[:]
	*/

	/// 切片的长度 和 容量
	// len(s) cap (s)
	// 长度：切片的个数
	// 容量: 从切片的第一个元素开始数，到其底层数组元素末尾的个数。
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // 6，6

	// 截取切片使其长度为 0，6
	s = s[:0]
	printSlice(s)

	// 拓展其长度 4，6
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值 2，4
	s = s[2:]
	printSlice(s)


	/// 切片零值 nil
	var sn []int
	fmt.Println(sn, len(sn), cap(sn)) // [], 0, 0
	if sn == nil {
		fmt.Println("nil!")  // 打印 nil!
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
