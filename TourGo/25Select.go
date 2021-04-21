package main

import (
	"fmt"
	"time"
)

// select 语句使一个 goroutine 可以等待多个通信操作。
// 类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收
// 如果没有 case 可运行，它将阻塞，直到有 case 可运行。当多个case都准备好时会随机选择一个执行。

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default: // 如果有default，没有case匹配的时候，select不会阻塞，会执行default语句
			fmt.Println("default")
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// 阻塞直到select的 c 被传入值
			fmt.Println(<-c)
		}
		// 进入quit逻辑时，x被fibonacci写入值，但是没有被取出，所以select会阻塞，然后接受quit信息，select触发执行 case quit逻辑
		quit <- 0
	}()
	fibonacci(c, quit)

	/// 默认Select
	// 当 select 中的其它分支都没有准备好时，default 分支就会执行
	// 为了在尝试发送或者接收时不发生阻塞，可使用 default 分支：
	/*
		select {
		case i := <-c:
				// 使用 i
		default:
				// 从 c 中接收会阻塞时执行
		}
	*/

	// Time
	// tick 和 boom 都是channel，等待一段时间后会触发
	// time.After方法，它返回一个类型为 "<-chan Time" 的单向的channel，在指定的时间发送一个当前时间给返回的channel中。
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			// Tick 和 boom 未触发时，执行default逻辑
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	// Timer 和 Ticker: 关于时间的两个 Channel
	// timer是一个定时器，代表未来的一个单一事件，你可以告诉timer你要等待多长时间，它提供一个Channel，在将来的那个时间那个Channel提供了一个时间值。
	// 下面的例子中第二行会阻塞2秒钟左右的时间，直到时间到了才会继续执行。
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")
	// 如果只是想单纯的等待的话，可以使用time.Sleep来实现。

	// 还可以使用timer.Stop来停止计时器。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// ticker是一个定时触发的计时器，它会以一个间隔(interval)往Channel发送一个事件(当前时间)，
	// 而Channel的接收者可以以固定的时间间隔从Channel中读取事件。下面的例子中ticker每500毫秒触发一次，你可以观察输出的时间。
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	// 类似timer, ticker也可以通过Stop方法来停止。一旦它停止，接收者不再会从channel中接收数据了。
}
