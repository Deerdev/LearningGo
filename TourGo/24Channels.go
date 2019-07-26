package main

import "fmt"

/// 信道
// 信道是带有类型的管道，通过信道操作符 <- 来发送或者接收值。

// ch <- v    // 将 v 发送至信道 ch。
// v := <-ch  // 从 ch 接收值并赋予 v。
// （“箭头”就是数据流的方向。）

// 和映射与切片一样，信道在使用前必须创建：
// ch := make(chan int)

// *** 默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

/// Buffered Channels 带缓存的Channels
// 信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：
// ch := make(chan int, 100)
// *** 仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

/// Close and range
// 发送者可通过 close 关闭一个信道来表示没有需要发送的值了
// 接收者可以通过为接收表达式 分配第二个参数 来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完
// v, ok := <-ch
// 之后 ok 会被设置为 false。

// 循环 for i := range c 会不断从信道接收值，直到它被关闭

// *注意：* 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。
// *还要注意：* 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// 此处阻塞，等待 sum 的返回，哪个先返回，哪个先赋值？
	x, y := <-c, <-c // 从 c 中接收
	fmt.Println(x, y, x+y)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // 没有可接收的数据 fatal error: all goroutines are asleep - deadlock!
	// 只能发送多少个，接收多少个；一旦发送超出缓存大小，或者接收超出缓存大小 就会无限期等待，deadlock

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		// 循环直到 C 被关闭
		fmt.Println(i)
	}
}
