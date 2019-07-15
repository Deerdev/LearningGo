package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 直接初始化
var mv = map[string]Vertex {
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// 只是一个类型，可以省略
var mv2 = map[string]Vertex {
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

var m1 map[string]Vertex

func main() {
	/// map 的零值为 nil，既没有键，也不能添加键。
	// make 函数会返回给定类型的Map，并将其初始化备用。
	// map[key]value
	m1 = make(map[string]Vertex)		// 必需初始化才能使用
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])

	/// map 的操作
	m := make(map[string]int)

	// 插入或修改
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	// 删除 key-value
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 获取value，如果不存在 ok 为 false， v 为 零值(此处为int的零值 0)
	// elem = m[key] 不存在直接为零值
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}