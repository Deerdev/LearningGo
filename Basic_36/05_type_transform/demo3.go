package main

import "fmt"

func main() {
    // 示例1。
    {
        type MyString = string
        str := "BCD"
        myStr1 := MyString(str)
        myStr2 := MyString("A" + str)
        fmt.Printf("%T(%q) == %T(%q): %v\n",
            str, str, myStr1, myStr1, str == myStr1)
        fmt.Printf("%T(%q) > %T(%q): %v\n",
            str, str, myStr2, myStr2, str > myStr2)
        fmt.Printf("Type %T is the same as type %T.\n", myStr1, str)

        strs := []string{"E", "F", "G"}
        myStrs := []MyString(strs)
        fmt.Printf("A value of type []MyString: %T(%q)\n",
            myStrs, myStrs)
        fmt.Printf("Type %T is the same as type %T.\n", myStrs, strs)
        fmt.Println()
    }
    // 示例2: 两个不同的类型
    {
        // MyString和string就是两个不同的类型了。这里的MyString2是一个新的类型，不同于其 他任何类型。
        //这种方式也可以被叫做对类型的再定义。我们刚刚把string类型再定义成了另外一个类型 MyString。
        //对于这里的类型再定义来说，string可以被称为MyString的潜在类型。潜在类型的含义是某个类型在本质上是哪个类型或者是哪个类型的集合。
        type MyString string
        str := "BCD"
        myStr1 := MyString(str)
        myStr2 := MyString("A" + str)
        _ = myStr2
        //fmt.Printf("%T(%q) == %T(%q): %v\n",
        //	str, str, myStr1, myStr1, str == myStr1) // 这里的判等不合法，会引发编译错误。
        //fmt.Printf("%T(%q) > %T(%q): %v\n",
        //	str, str, myStr2, myStr2, str > myStr2) // 这里的比较不合法，会引发编译错误。
        fmt.Printf("Type %T is different from type %T.\n", myStr1, str)

        strs := []string{"E", "F", "G"}
        var myStrs []MyString
        //myStrs := []MyString(strs) // 这里的类型转换不合法，会引发编译错误。
        //fmt.Printf("A value of type []MyString: %T(%q)\n",
        //	myStrs, myStrs)
        fmt.Printf("Type %T is different from type %T.\n", myStrs, strs)
        fmt.Println()
    }
    // 示例3。
    {
        type MyString1 = string
        type MyString2 string
        str := "BCD"
        myStr1 := MyString1(str)
        myStr2 := MyString2(str)
        myStr1 = MyString1(myStr2)
        myStr2 = MyString2(myStr1)

        myStr1 = str
        //myStr2 = str // 这里的赋值不合法，会引发编译错误。
        //myStr1 = myStr2 // 这里的赋值不合法，会引发编译错误。
        //myStr2 = myStr1 // 这里的赋值不合法，会引发编译错误。
    }
}
