package main

import "unsafe"
import "fmt"

type Named interface {
    // Name 用于获取名字。
    Name() string
}

type Dog struct {
    name string
}

func (dog *Dog) SetName(name string) {
    dog.name = name
}

func (dog Dog) Name() string {
    return dog.name
}


func New(name string) Dog {
    // 调用New函数所得到的结果值属于临时结 果，是不可寻址的。
    return Dog{name}
}

func main() {
    // --- 不可寻址类型 ---
    // 示例1。
    const num = 123
    //_ = &num // 常量不可寻址。
    //_ = &(123) // 基本类型值的字面量不可寻址。

    var str = "abc"
    _ = str
    //_ = &(str[0]) // 对字符串变量的索引结果值不可寻址。
    //_ = &(str[0:2]) // 对字符串变量的切片结果值不可寻址。
    str2 := str[0]
    _ = &str2 // 但这样的寻址就是合法的。

    //_ = &(123 + 456) // 算术操作的结果值不可寻址。
    num2 := 456
    _ = num2
    //_ = &(num + num2) // 算术操作的结果值不可寻址。

    //_ = &([3]int{1, 2, 3}[0]) // 对数组字面量的索引结果值不可寻址。
    //_ = &([3]int{1, 2, 3}[0:2]) // 对数组字面量的切片结果值不可寻址。
    _ = &([]int{1, 2, 3}[0]) // 对切片字面量的索引结果值却是可寻址的。
    //_ = &([]int{1, 2, 3}[0:2]) // 对切片字面量的切片结果值不可寻址。
    //_ = &(map[int]string{1: "a"}[0]) // 对字典字面量的索引结果值不可寻址。

    var map1 = map[int]string{1: "a", 2: "b", 3: "c"}
    _ = map1
    //_ = &(map1[2]) // 对字典变量的索引结果值不可寻址。

    //_ = &(func(x, y int) int {
    //	return x + y
    //}) // 字面量代表的函数不可寻址。
    //_ = &(fmt.Sprintf) // 标识符代表的函数不可寻址。
    //_ = &(fmt.Sprintln("abc")) // 对函数的调用结果值不可寻址。

    dog := Dog{"little pig"}
    _ = dog
    //_ = &(dog.Name) // 标识符代表的函数不可寻址。
    //_ = &(dog.Name()) // 对方法的调用结果值不可寻址。

    //_ = &(Dog{"little pig"}.name) // 结构体字面量的字段不可寻址。

    //_ = &(interface{}(dog)) // 类型转换表达式的结果值不可寻址。
    dogI := interface{}(dog)
    _ = dogI
    //_ = &(dogI.(Named)) // 类型断言表达式的结果值不可寻址。
    named := dogI.(Named)
    _ = named
    //_ = &(named.(Dog)) // 类型断言表达式的结果值不可寻址。

    var chan1 = make(chan int, 1)
    chan1 <- 1
    //_ = &(<-chan1) // 接收表达式的结果值不可寻址。


    // ---- 不可寻址的值在使用上有哪些限制 ----
    /*
    对于一个Dog类型的变量dog来说，调用表达式dog.SetName("monster")会被 自动地转译为(&dog).SetName("monster")，
    即:先取dog的指针值，再在该指针值上调用 SetName方法。
    发现问题了吗?由于New函数的调用结果值是不可寻址的，所以无法对它进行取址操作。
    因此， 上边这行链式调用会让编译器报告两个错误，一个是果，即:不能在New("little pig")的 结果值上调用指针方法。一个是因，即:不能取得New("little pig")的地址
     */
    // 示例1。
    //New("little pig").SetName("monster") // 不能调用不可寻址的值的指针方法。

    // 示例2。
    /*
    虽然 Go 语言规范中的语法定义是，只要在++或--的左边添加一个表达式(指针的偏移)，就可以组成一个自 增语句或自减语句，
    但是，它还明确了一个很重要的限制，那就是这个表达式的结果值必须是可 寻址的。这就使得针对值字面量的表达式几乎都无法被用在这里。

    不过这有一个例外，虽然对字典字面量和字典变量索引表达式的结果值都是不可寻址的，但是这
    样的表达式却可以被用在自增语句和自减语句中
     */
    map[string]int{"the": 0, "word": 0, "counter": 0}["word"]++

    map2 := map[string]int{"the": 0, "word": 0, "counter": 0}
    map2["word"]++


    // ---- 通过unsafe.Pointer操纵可寻址的值 ---
    // 示例1。
    /*
    1. 一个指针值(比如*Dog类型的值)可以被转换为一个unsafe.Pointer类型的值，反之亦 然。
    2. 一个uintptr类型的值也可以被转换为一个unsafe.Pointer类型的值，反之亦然。
    3. 一个指针值无法被直接转换成一个uintptr类型的值，反过来也是如此。
     */
    dog2 := Dog{"little pig"}
    dogP := &dog2
    dogPtr := uintptr(unsafe.Pointer(dogP))

    /*
    unsafe.Offsetof函数用于获取 两个值在内存中的起始存储地址之间的偏移量，以字节为单位。
    这两个值一个是某个字段的值，另一个是该字段值所属的那个结构体值。我们在调用这个函数的 时候，需要把针对字段的选择表达式传给它，比如dogP.name。
    有了这个偏移量，又有了结构体值在内存中的起始存储地址(这里由dogPtr变量代表)，把它们相加我们就可以得到dogP的name字段值的起始存储地址了。这个地址由变量namePtr代表。
    此后，我们可以再通过两次类型转换把namePtr的值转换成一个*string类型的值，这样就得 到了指向dogP的name字段值的指针值。

    你可能会问，我直接用取址表达式&(dogP.name)不就能拿到这个指针值了吗?干嘛绕这么大 一圈呢?你可以想象一下，如果我们根本就不知道这个结构体类型是什么，也拿不到dogP这个 变量，那么还能去访问它的name字段吗?
    答案是，只要有namePtr就可以。它就是一个无符号整数，但同时也是一个指向了程序内部数据的内存地址。它可能会给我们带来一些好处，比如可以直接修改埋藏得很深的内部数据。
    但是，一旦我们有意或无意地把这个内存地址泄露出去，那么其他人就能够肆意地改动 dogP.name的值，以及周围的内存地址上存储的任何数据了。
     */
    namePtr := dogPtr + unsafe.Offsetof(dogP.name)
    nameP := (*string)(unsafe.Pointer(namePtr))
    fmt.Printf("nameP == &(dogP.name)? %v\n",
        nameP == &(dogP.name))
    fmt.Printf("The name of dog is %q.\n", *nameP)

    *nameP = "monster"
    fmt.Printf("The name of dog is %q.\n", dogP.name)
    fmt.Println()

    // 示例2。
    // 下面这种不匹配的转换虽然不会引发panic，但是其结果往往不符合预期。
    numP := (*int)(unsafe.Pointer(namePtr))
    num3 := *numP
    fmt.Printf("This is an unexpected number: %d\n", num3)
}
