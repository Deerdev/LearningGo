package main

import "fmt"

type Teacher struct {
    name string
    age int
}

type Student struct {
    name string
    grade int
}

type Course struct {
    // 内嵌结构体
    teacher Teacher
    // 匿名内嵌结构体, 会展开所有字段
    Student
    price int
    name string
    url string
}

func getInfo(c Course){
    // c.grade, 匿名嵌套，直接展开属性，所以可以直接调用
    // 重名时，需要显式指定 c.Student.name
    fmt.Println(c.teacher.name, c.teacher.age, c.grade, c.Student.name)
}

func main() {
    var c Course = Course {
        teacher: Teacher{
            name: "bob",
            age:18,
        },
        // 匿名结构体的初始化
        Student: Student{
            name: "joy",
            grade: 6,
        },
        price: 100,
        name: "english",
        url: "test",
        // 注意这里的逗号不能少
    }
    getInfo(c)
}
