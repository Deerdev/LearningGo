package main

import (
	"fmt"
	"sort"
)

type Course struct {
	Name  string
	Price int
	Url   string
}

type Courses []Course

func (c Courses) Len() int {
	return len(c)
}

func (c Courses) Less(i, j int) bool {
	return c[i].Price < c[j].Price
}

func (c Courses) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	//通过sort来排序
	courses := Courses{
		Course{"django", 300, ""},
		Course{"scrapy", 100, ""},
		Course{"go", 400, ""},
		Course{"torando", 200, ""},
	}
	sort.Sort(courses) //协议 实现 sort 的协议方法
	for _, v := range courses {
		fmt.Println(v)
	}
}
