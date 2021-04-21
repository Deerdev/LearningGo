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

// 只是一个类型，可以省略Vertex
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

	/// 遍历
    for k, v := range m {
        fmt.Printf("k:[%v].v:[%v]\n", k, v) // 输出k,v值
    }


}

// 遍历时，k，v 的变量都执行同一个内存地址，每次循环时，copy map 中的 key-value
// 如果存储 k、v 的指针，最终的值都是一样的
func mapLoopQuestion() {
    m := map[string]int{
        "a": 1,
        "b": 2,
    }
    var bs []*int
    for k, v := range m {
        fmt.Printf("k:[%p].v:[%p]\n", &k, &v) // 这里的输出可以看到，k一直使用同一块内存，v也是这个状况
        bs = append(bs, &v) // 对v取了地址
    }
    // k:[0xc000010200].v:[0xc0000140c0]
    // k:[0xc000010200].v:[0xc0000140c0]


    // 输出
    for _, b := range bs {
        fmt.Println(*b) // 输出都是1或者都是2，因为 map loop 的顺序不定，可能最后遍历的是 1 或 2
    }
}


/// map 的 key 要求: 需要支持 == or != 操作
// key的常用类型：int, rune, string, 结构体(每个元素需要支持 == or != 操作), 指针, 基于这些类型自定义的类型
func mapKey() {
    // m0 可以, key类型为string, 支持 == 比较操作
    {
        var m0 map[string]string // 定义map类型变量m0，key的类型为string，value的类型string
        fmt.Println(m0)
    }

    // m1 不可以, []byte是slice，不支持 == != 操作，不可以作为map key的数据类型
    {
        //var m1 map[[]byte]string // 报错： invalid map key type []byte
        //fmt.Println(m1)

        // 准确说slice类型只能与nil比较，其他的都不可以，可以通过如下测试：
        // var b1,b2 []byte
        // fmt.Println(b1==b2) // 报错： invalid operation: b1 == b2 (slice can only be compared to nil)
    }

    // m2 可以, interface{}类型可以作为key，但是需要加入的key的类型是可以比较的
    {
        var m2 map[interface{}]string
        m2 = make(map[interface{}]string)
        //m2[[]byte("k2")]="v2" // panic: runtime error: hash of unhashable type []uint8
        m2[123] = "123"
        m2[12.3] = "123"
        fmt.Println(m2)
    }

    // m3 可以， 数组支持比较
    {
        a3 := [3]int{1, 2, 3}
        var m3 map[[3]int]string
        m3 = make(map[[3]int]string)
        m3[a3] = "m3"
        fmt.Println(m3)
    }

    // m4 可以，book1里面的元素都是支持== !=
    {
        type book1 struct {
            name string
        }
        var m4 map[book1]string
        fmt.Println(m4)
    }

    // m5 不可以, text元素类型为[]byte, 不满足key的要求
    {
        // type book2 struct {
        //  name string
        //  text []byte //没有这个就可以
        // }
        //var m5 map[book2]string //invalid map key type book2
        //fmt.Println(m5)
    }
}
