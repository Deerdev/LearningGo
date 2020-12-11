package main

import (
    "fmt"
    "math/rand"
)

// 单向通道

func main() {
    // 示例1。
    // 只能发不能收的通道。
    var uselessChan = make(chan<- int, 1)
    // 只能收不能发的通道。
    var anotherUselessChan = make(<-chan int, 1)
    // 这里打印的是可以分别代表两个通道的指针的16进制表示。
    fmt.Printf("The useless channels: %v, %v\n",
        uselessChan, anotherUselessChan)

    // 示例2。
    intChan1 := make(chan int, 3)
    SendInt(intChan1)

    // 示例4。
    /*
    一、这样一条for语句会不断地尝试从intChan2种取出元素值，即使intChan2被关闭，它 也会在取出所有剩余的元素值之后再结束执行。
    二、当intChan2中没有元素值时，它会被阻塞在有for关键字的那一行，直到有新的元素值 可取。
    三、假设intChan2的值为nil，那么它会被永远阻塞在有for关键字的那一行。
     */
    intChan2 := getIntChan()
    for elem := range intChan2 {
        fmt.Printf("The element in intChan2: %v\n", elem)
    }

    // 示例5。
    _ = GetIntChan(getIntChan)
}

// 限制函数的参数类型，在函数体内只能使用接收或发送
// 示例2。
func SendInt(ch chan<- int) {
    ch <- rand.Intn(1000)
}

// 示例3。
type Notifier interface {
    SendInt(ch chan<- int)
}

// 示例4。
func getIntChan() <-chan int {
    num := 5
    ch := make(chan int, num)
    for i := 0; i < num; i++ {
        ch <- i
    }
    close(ch)
    return ch
}

// 示例5。
type GetIntChan func() <-chan int
