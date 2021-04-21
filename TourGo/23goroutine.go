

package main

import (
	"fmt"
	"time"
)

// goroutine 是由 Go 运行时管理的轻量级线程

// go f(x, y, z) 
// 会启动一个新的 Go 程并执行 f(x, y, z)
// f 的调用发生在当前 goroutine，而 f 的执行（x,y,z的求值）发生在新的 goroutine
// goroutine 在相同的地址空间中运行，因此在访问共享的内存时必须进行同步(加锁)。sync 包提供了这种能力，不过在 Go 中并不经常用到

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

