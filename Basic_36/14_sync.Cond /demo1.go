package main

import (
    "log"
    "sync"
    "time"
)
// --- 条件变量和互斥锁 sync.Cond sync.Locker---
/*
- 条件变量并不是被用来保护临界区和共享资源的，它是用于协调想要访问共享资源的那些线程的。当共享资源的状态发生变化时，它可以被用来通知被互斥锁阻塞的线程
- 条件变量怎样与互斥锁配合使用？
    - 这道题的典型回答是：条件变量的初始化离不开互斥锁，并且它的方法有的也是基于互斥锁的。
- 条件变量提供的方法有三个：等待通知（wait）、单发通知（signal）和广播通知（broadcast）。我们在利用条件变量等待通知的时候，需要在它基于的那个互斥锁
保护下进行。而在进行单发通知或广播通知的时候，却是恰恰相反的，也就是说，需要在对应的 互斥锁解锁之后再做这两种操作
 */

func main() {
    // mailbox 代表信箱。
    // 0代表信箱是空的，1代表信箱是满的。
    var mailbox uint8
    // lock 代表信箱上的锁。
    var lock sync.RWMutex
    // sendCond 代表专用于发信的条件变量。
    /*
       sync.Locker其实是一个接口，在它的声明中只包含了两个方法定义，即：Lock()和
       Unlock()。sync.Mutex类型和sync.RWMutex类型都拥有Lock方法和Unlock方法，只不过
       它们都是指针方法。因此，这两个类型的指针类型才是sync.Locker接口的实现类型。
    */
    sendCond := sync.NewCond(&lock)     // sync.NewCond 需要一个sync.Locker类型的参数值

    /*
    lock变量中用于对[读锁]进行锁定和解锁的方法却是RLock和RUnlock，它们与
    sync.Locker接口中定义的方法并不匹配;
    - lock.RLocker() 这个值所拥有的Lock方法和Unlock方法，在其内部会分别调用lock变量的RLock方法和RUnlock方法。也就
    是说，前两个方法仅仅是后两个方法的代理而已。
     */
    // recvCond 代表专用于收信的条件变量。
    recvCond := sync.NewCond(lock.RLocker())

    // sign 用于传递演示完成的信号。
    sign := make(chan struct{}, 3)
    max := 5
    go func(max int) { // 用于发信。
        defer func() {
            sign <- struct{}{}
        }()
        for i := 1; i <= max; i++ {
            time.Sleep(time.Millisecond * 500)
            lock.Lock()
            for mailbox == 1 {
                sendCond.Wait()
            }
            log.Printf("sender [%d]: the mailbox is empty.", i)
            mailbox = 1
            log.Printf("sender [%d]: the letter has been sent.", i)
            lock.Unlock()
            recvCond.Signal()
        }
    }(max)
    go func(max int) { // 用于收信。
        defer func() {
            sign <- struct{}{}
        }()
        for j := 1; j <= max; j++ {
            time.Sleep(time.Millisecond * 500)
            lock.RLock()
            for mailbox == 0 {
                recvCond.Wait()
            }
            log.Printf("receiver [%d]: the mailbox is full.", j)
            mailbox = 0
            log.Printf("receiver [%d]: the letter has been received.", j)
            lock.RUnlock()
            sendCond.Signal()
        }
    }(max)

    <-sign
    <-sign
}

// --- Wait() 方法在做什么 ---
/*
1. 把调用它的 goroutine（也就是当前的 goroutine）加入到当前条件变量的通知队列中。
2. 解锁当前的条件变量基于的那个互斥锁。
3. 让当前的 goroutine 处于等待状态，等到通知到来时再决定是否唤醒它。此时，这个 goroutine 就会阻塞在调用这个Wait方法的那行代码上。
4. 如果通知到来并且决定唤醒这个 goroutine，那么就在唤醒它之后重新锁定当前条件变量基于的互斥锁。自此之后，当前的 goroutine 就会继续执行后面的代码了。

原则：加锁和解锁成对出现，在一个 goroutine 中
 */


// --- 为什么在 for 循环里使用 Wait() ---
/*
这主要是为了保险起见。如果一个 goroutine 因收到通知而被唤醒，但却发现共享资源的状态，
依然不符合它的要求，那么就应该再次调用条件变量的Wait方法，并继续等待下次通知的到
来。这种情况是很有可能发生的，具体如下面所示:
1. 有多个 goroutine 在等待共享资源的同一种状态。比如，它们都在等mailbox变量的值不为0的时候再把它的值变为0，这就相当于有多个人在等着我向信箱里放置情报。虽然等待的
goroutine 有多个，但每次成功的 goroutine 却只可能有一个。别忘了，条件变量的Wait方法会在当前的 goroutine 醒来后先重新锁定那个互斥锁。在成功的 goroutine 最终解锁互斥
锁之后，其他的 goroutine 会先后进入临界区，但它们会发现共享资源的状态依然不是它们想要的。这个时候，for循环就很有必要了。

2. 共享资源可能有的状态不是两个，而是更多。比如，mailbox变量的可能值不只有0和1，还有2、3、4。这种情况下，由于状态在每次改变后的结果只可能有一个，所以，在设计合
理的前提下，单一的结果一定不可能满足所有 goroutine 的条件。那些未被满足的goroutine 显然还需要继续等待和检查。

3. 有一种可能，共享资源的状态只有两个，并且每种状态都只有一个 goroutine 在关注，就像我们在主问题当中实现的那个例子那样。不过，即使是这样，使用for语句仍然是有必要
的。原因是，在一些多 CPU 核心的计算机系统中，即使没有收到条件变量的通知，调用其Wait方法的 goroutine 也是有可能被唤醒的。这是由计算机硬件层面决定的，即使是操作
系统（比如 Linux）本身提供的条件变量也会如此。
 */