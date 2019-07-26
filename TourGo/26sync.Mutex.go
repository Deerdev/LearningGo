package main

import (
	"fmt"
	"sync"
	"time"
)

/// 互斥锁
// 当不需要通信的时候，channel 就不适合使用
// 若保证每次只有一个 goroutine 能够访问一个共享的变量，需要使用 互斥锁（Mutex）
// Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法：
// Lock
// Unlock
// 可以用 defer 语句来保证互斥锁一定会被解锁

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
