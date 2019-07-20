package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// fmt 包中定义的 Stringer 是最普遍的接口之一
/*
type Stringer interface {
    String() string
}
*/

func (p Person) String() string {
	// Person 类型的自我描述，实现 Stringer 接口
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
