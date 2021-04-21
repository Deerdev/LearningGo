package main

import (
    "fmt"
    "unsafe"
)

type Vertex struct {
	X int
	Y int
}

// 结构体初始化
var (
    // 不指定属性时，需要全部属性初始化
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)


type ArrayStruct struct {
    value [10]int
}

type SliceStruct struct {
    value []int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// 结构体指针
	p := &v
	p.X = 1e9	// 根据指针的定义，应该使用 (*p).X ，这么写太啰嗦，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	fmt.Println(v)

    //  「零值」结构体 初始化
    var c1 Vertex = Vertex{}
    var c2 Vertex
    var c3 *Vertex = new(Vertex)
    fmt.Println(c1)
    fmt.Println(c2)
    fmt.Println(c3)
    // 各属性的零值: int 0, string 空字符串
    //{0  }
    //{0  }
    //&{0  }

    /// 零值和 nil 结构体
    // 零值是属性值 为零值，但是结构体依然分配了内存
    // nil 结构体，没有实际为结构体分配内存
    var c *Vertex = nil
    fmt.Println(c)

    /// 结构体的拷贝
    // 结构体变量之间的赋值，是浅拷贝，值类型 复制，引用类型 复制指针的值


    /// 数组和切片 占用的空间大小
    var as = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
    var ss = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
    fmt.Println(unsafe.Sizeof(as), unsafe.Sizeof(ss))   // 80 24
}

/// Go 基本数据类型都是 struct
//slice的结构体
//通过观察 Go 语言的底层源码，可以发现所有的 Go 语言内置的高级数据结构都是由结构体来完成的。
//切片头的结构体形式如下，它在 64 位机器上将会占用 24 个字节
//slice的函数传递本质上也是值传递
type slice struct {
    array unsafe.Pointer  // 底层数组的地址
    len int // 长度
    cap int // 容量
}



//字符串头的结构体
//它在 64 位机器上将会占用 16 个字节
type string struct {
    array unsafe.Pointer // 底层数组的地址
    len int
}

//map的结构体
type hmap struct {
    count int
    //...
    buckets unsafe.Pointer  // hash桶地址
    //...
}
