package main

import (
	"flag" // https://golang.google.cn/pkg/flag/
	"fmt"
	"os"
)

// 方式1：常规
// 方式2：使用底层方法 flag.CommandLine 定义flag的一些设置
// 方式3：自定义全局私有的命令参数容器 cmdLine, 任何设置都和 flag.CommandLine 分开

// 方式1。
var name string
// flag.String 和 flag.StringVar 的区别
// var name = flag.String("name", "everyone", "The greeting object.")

// 方式3。
//var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// 方式2。
	// 在调用flag包中的一些函数(比如StringVar、Parse等等)的时候， 实际上是在调用flag.CommandLine变量的对应方法。
	// flag.CommandLine相当于默认情况下的命令参数容器。
	// 所以，通过对flag.CommandLine 重新赋值，我们可以更深层次地定制当前命令源码文件的参数使用说明

	// flag.PanicOnError和flag.ExitOnError都是预定义在flag包中的常量。
		// flag.ExitOnError的含义是，告诉命令参数容器，当命令后跟--help或者参数设置的不正 确的时候，在打印命令参数使用说明后以状态码2结束当前程序。
		// 状态码2代表用户错误地使用了命令，而flag.PanicOnError与之的区别是在最后抛出“运行 时恐慌(panic)”
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	// 方式3。
	//cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")

	// 方式1/2
	// 第 1 个参数是用于存储该命令参数的值的地址，具体到 这里就是在前面声明的变量name的地址了，由表达式&name表示。
	// 第 2 个参数是为了指定该命令参数的名称，这里是name。
	// 第 3 个参数是为了指定在未追加该命 令参数时的默认值，这里是everyone。
	// 第 4 个函数参数，即是该命令参数的简短说明了，这在打印命令说明时会用到。
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// 方式1。
	// Usage 打印 --help信息
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}

	// 方式1 / 2
	// 真正解 析命令参数，并把它们的值赋给相应的变量
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)

	// 方式3。
	//cmdLine.Parse(os.Args[1:])	// os.Args[1:]指的就是我们给定的那些命令参数


	
}

// go run Basic/flag_cmd/01.go -name="lining"
// go run Basic/flag_cmd/01.go --help