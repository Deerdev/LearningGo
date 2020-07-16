package lib

import "fmt"

// 大写，包外引用
func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
