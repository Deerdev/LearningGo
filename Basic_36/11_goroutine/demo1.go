package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// --- 等待执行完成
	num := 10
	sign := make(chan struct{}, num)

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	// 办法1。
	//time.Sleep(time.Millisecond * 500)

	// 办法2：利用空的 struct channel 通信
	for j := 0; j < num; j++ {
		<-sign
	}

	// --- 按顺序执行多个 goroutine ---
	var count uint32 // race condition 需要加锁，原子操作，在sync/atomic包中声明了很多用于原子操作的函数
	trigger := func(i uint32, fn func()) {
		for { // 循环等待与 i 相等的 count 出现，从 0 到 10 一个一个执行
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}

	// 主 goroutine 使用 trigger 等待
	trigger(10, func() {})
}
