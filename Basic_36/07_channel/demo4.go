package main

import "fmt"

var channels = [3]chan int{
    nil,
    make(chan int),
    nil,
}

var numbers = []int{1, 2, 3}

func main() {
    /*
    当select语句被执行时，它会根据一套分支选择规则选中某一个分支并执行其中的代码。
    如果 所有的候选分支都没有被选中，那么默认分支(如果有的话)就会被执行。
    注意，发送和接收操 作的阻塞是分支选择规则的一个很重要的依据
     */

    // 先执行 case 语句，再选择分支
    select {
    case getChan(0) <- getNumber(0):
        fmt.Println("The first candidate case is selected.")
    case getChan(1) <- getNumber(1):
        fmt.Println("The second candidate case is selected.")
    case getChan(2) <- getNumber(2):
        fmt.Println("The third candidate case is selected")
    default:
        fmt.Println("No candidate case is selected!")
    }
}

func getNumber(i int) int {
    fmt.Printf("numbers[%d]\n", i)
    return numbers[i]
}

func getChan(i int) chan int {
    fmt.Printf("channels[%d]\n", i)
    return channels[i]
}

/*
问题 2:select语句的分支选择规则都有哪些? 规则如下面所示。
1. 对于每一个case表达式，都至少会包含一个代表发送操作的发送表达式或者一个代表接收 操作的接收表达式，同时也可能会包含其他的表达式。比如，如果case表达式是包含了接
收表达式的短变量声明时，那么在赋值符号左边的就可以是一个或两个表达式，不过此处的 表达式的结果必须是可以被赋值的。当这样的case表达式被求值时，它包含的多个表达式 总会以从左到右的顺序被求值。
2. select语句包含的候选分支中的case表达式都会在该语句执行开始时先被求值，并且求值 的顺序是依从代码编写的顺序从上到下的。结合上一条规则，在select语句开始执行时， 排在最上边的候选分支中最左边的表达式会最先被求值，然后是它右边的表达式。仅当最上 边的候选分支中的所有表达式都被求值完毕后，从上边数第二个候选分支中的表达式才会被 求值，顺序同样是从左到右，然后是第三个候选分支、第四个候选分支，以此类推。
3. 对于每一个case表达式，如果其中的发送表达式或者接收表达式在被求值时，相应的操作 正处于阻塞状态，那么对该case表达式的求值就是不成功的。在这种情况下，我们可以 说，这个case表达式所在的候选分支是不满足选择条件的。
4. 仅当select语句中的所有case表达式都被求值完毕后，它才会开始选择候选分支。这时 候，它只会挑选满足选择条件的候选分支执行。如果所有的候选分支都不满足选择条件，那 么默认分支就会被执行。如果这时没有默认分支，那么select语句就会立即进入阻塞状 态，直到至少有一个候选分支满足选择条件为止。一旦有一个候选分支满足选择条件， select语句(或者说它所在的 goroutine)就会被唤醒，这个候选分支就会被执行。
5. 如果select语句发现同时有多个候选分支满足选择条件，那么它就会用一种伪随机的算法 在这些分支中选择一个并执行。注意，即使select语句是在被唤醒时发现的这种情况，也会这样做。
6. 一条select语句中只能够有一个默认分支。并且，默认分支只在无候选分支可选时才会被 执行，这与它的编写位置无关。
7. select语句的每次执行，包括case表达式求值和分支选择，都是独立的。不过，至于它的 执行是否是并发安全的，就要看其中的case表达式以及分支中，是否包含并发不安全的代 码了。
 */