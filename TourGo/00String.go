package main

import (
    "fmt"
    "strings"
)

func main() {
    // 1. 字符串长度
    // 涉及中文时，编码方式会改变
    // unicode 字符集，存储的时候是 utf8, 但是 utf8是动态编码规则（根据字符类型，确定用几位来保存）
    // utf8：一个字节可以表示英文；中国人：3x3 个字节
    var name string = "123中国人"
    fmt.Println(len(name)) // 计算的是字节数：12
    fmt.Println(name[3])   // 打印第4个字节：184

    // 转换成数组 rune=int32; 相当于全部转成 4 字节存储，空间消耗会增大
    nameArr := []rune(name)
    fmt.Println(len(nameArr)) // 计算的数组长度：6，字节数其实是：4*6=24（比string，占用更多的空间）

    // 2.不使用转义``
    date := `"2021/02/02"`
    fmt.Println(date)

    // 3. 字符串包含
    title := "test"
    if strings.Contains(title, "te") {
        fmt.Println("contains")
    }
    fmt.Println(strings.Index(title, "te"))

    // 4. 字符串统计
    fmt.Println(strings.Count(title, "t"))
    fmt.Println(strings.HasPrefix(title, "te"))
    fmt.Println(strings.HasSuffix(title, "st"))

    // 5. 大小写
    fmt.Println(strings.ToUpper(title))
    fmt.Println(strings.ToLower(title))

    // 6.比较: 0: a==b, 1: a>b, -1: a<b
    // 比较 ascii
    fmt.Println(strings.Compare(title, "test"))

    // 7. trim: 去掉空格或指定字符
    // TrimLeft/TrimRight/TrimFunc
    fmt.Println(strings.TrimSpace(title))
    fmt.Println(strings.Trim(title, "t"))

    // 8. split
    fmt.Println(strings.Split("1 2", " "))

    // 9. join
    arr := strings.Split("1 2 3", " ")
    fmt.Println(strings.Join(arr, ",")) // 1,2,3

    // 10. replace
    // src, old, new, n:替换次数（n<0全部替换）
    fmt.Println(strings.Replace("test123", "123", "456", 1))

    // 11. 格式化输出
    // %T 类型，%v 值，%#v 以符合 Go 语法的方式输出（字符串有引号）
    // 其他和 C 语言一致
    fmt.Printf("title: %v, name: %#v", title, name)

    // 12. 输入
    var x int
    var y float64
    fmt.Println("请输入一个整数，一个浮点类型：")
    // n 扫描的个数：2, 输入：1 1.3
    n, _ := fmt.Scanln(&x, &y) //读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
    fmt.Printf("x的数值：%d，y的数值：%f, n: %v\n", x, y, n)

    var a string
    var b int
    // 输入的内容必须符合指定的格式: sa 1
    fmt.Scanf("%s %d", &a, &b)
    fmt.Println(a, b)
}
