package main
import "fmt"

func main() {
    // --- 阻塞通道案例 ---
    // 示例1。
    ch1 := make(chan int, 1)
    ch1 <- 1
    //ch1 <- 2 // 通道已满，因此这里会造成阻塞。

    // 示例2。
    ch2 := make(chan int, 1)
    //elem, ok := <-ch2 // 通道已空，因此这里会造成阻塞。
    //_, _ = elem, ok
    ch2 <- 1

    // 示例3。
    var ch3 chan int
    //ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
    //<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
    _ = ch3


    // --- 通过关闭，接收方加保护 ---
    ch4 := make(chan int, 2)
    // 发送方。
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Printf("Sender: sending element %v...\n", i)
            ch4 <- i
        }
        fmt.Println("Sender: close the channel...")
        close(ch4)
    }()

    // 接收方。
    for {
        elem, ok := <-ch4
        if !ok {
            fmt.Println("Receiver: closed channel")
            break
        }
        fmt.Printf("Receiver: received an element: %v\n", elem)
    }

    fmt.Println("End.")
}

/*
channel 底层是环线链表

1. 对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。
2. 发送操作和接收操作中对元素值的处理都是不可分割的。
3. 发送操作在完全完成之前会被阻塞。接收操作也是如此。

发送：发送操作包括了“复制元素值”和“放置副本到通道内 部”这两个步骤。
接收：接收操作通常包含了“复制通道内的元素值”“放置副本到接收方”“删掉原值”三个步
骤。

通道的长度和容量：长度代表通道当前包含的元素个数，容量就是初始化时设置的那个数；
通道的拷贝方式：浅拷贝? go 的结构体都是值类型；而对切片和指针在 channel 中传递时，需要注意，指向的实例不会拷贝；


 */