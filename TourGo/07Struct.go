package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// 结构体初始化
var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// 结构体指针
	p := &v
	p.X = 1e9	// 根据指针的定义，应该使用 (*p).X ，这么写太啰嗦，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	fmt.Println(v)
}