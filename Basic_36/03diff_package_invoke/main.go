package main

import (
    "flag"
	// 源码文件所在的目录相对于 src 目录的相对路径就是它的代码包导入路径，而实际 使用其程序实体时给定的限定符要与它声明所属的代码包名称对应
	"Base/03diff_package/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}
