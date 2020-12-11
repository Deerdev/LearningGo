package main

import "fmt"

func main() {
    // 数组类型的值(以下简称数组)的长度是固定的，而切片类型的值 (以下简称切片)是可变长的
    // 示例1
    s1 := make([]int, 5) // 长度和容量都是 5 的切片
    fmt.Printf("The length of s1: %d\n", len(s1))
    fmt.Printf("The capacity of s1: %d\n", cap(s1))
    fmt.Printf("The value of s1: %d\n", s1)
    s2 := make([]int, 5, 8) // 长度是 5，容量是 8（指向的底层数组总长 8，但是本切片只能看到 5 个），对底层数组的索引位置：[0, 4]
    fmt.Printf("The length of s2: %d\n", len(s2))
    fmt.Printf("The capacity of s2: %d\n", cap(s2))
    fmt.Printf("The value of s2: %d\n", s2)
    fmt.Println()
    /*
    The length of s1: 5
    The capacity of s1: 5
    The value of s1: [0 0 0 0 0]
    The length of s2: 5
    The capacity of s2: 8
    The value of s2: [0 0 0 0 0]
    */

    // 示例2。
    s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
    s4 := s3[3:6] // s4 切片长度 6-3=3，容量计算：切片只能从索引 3 开始 向右滑动，一直移动到数据结尾，所以容量是：8-3=5
    fmt.Printf("The length of s4: %d\n", len(s4))
    fmt.Printf("The capacity of s4: %d\n", cap(s4))
    fmt.Printf("The value of s4: %d\n", s4)
    fmt.Println()
    /*
    The length of s4: 3
    The capacity of s4: 5
    The value of s4: [4 5 6]
    */

    // 示例3。
    s5 := s4[:cap(s4)] // 取 s4
    fmt.Printf("The length of s5: %d\n", len(s5))
    fmt.Printf("The capacity of s5: %d\n", cap(s5))
    fmt.Printf("The value of s5: %d\n", s5)
    /*
    The length of s5: 5
    The capacity of s5: 5
    The value of s5: [4 5 6 7 8]
    */

    // --- 切片容量增长 ---
    /*
    在一般的情况下，你可以简单地认为新切片的容量(以下简称新容量)将会是原切片容量(以下 简称原容量)的 2 倍。
    但是，当原切片的长度(以下简称原长度)大于或等于1024时，Go 语言将会以原容量的1.25 倍作为新容量的基准(以下新容量基准)。
    新容量基准会被调整(不断地与1.25相乘)，直到结果不小于原长度与要追加的元素数量之和 (以下简称新长度)。最终，新容量往往会比新长度大一些，当然，相等也是可能的。
    另外，如果我们一次追加的元素过多，以至于使新长度比原容量的 2 倍还要大，那么新容量就 会以新长度为基准
     */
    // 示例1。
    s6 := make([]int, 0)
    fmt.Printf("The capacity of s6: %d\n", cap(s6))
    for i := 1; i <= 5; i++ {
        s6 = append(s6, i)
        fmt.Printf("s6(%d): len: %d, cap: %d\n", i, len(s6), cap(s6))
    }
    fmt.Println()
    /*
    The capacity of s6: 0
    s6(1): len: 1, cap: 1
    s6(2): len: 2, cap: 2
    s6(3): len: 3, cap: 4
    s6(4): len: 4, cap: 4
    s6(5): len: 5, cap: 8
    */

    // 示例2。
    s7 := make([]int, 1024)
    fmt.Printf("The capacity of s7: %d\n", cap(s7))
    s7e1 := append(s7, make([]int, 200)...)
    fmt.Printf("s7e1: len: %d, cap: %d\n", len(s7e1), cap(s7e1))
    s7e2 := append(s7, make([]int, 400)...)
    fmt.Printf("s7e2: len: %d, cap: %d\n", len(s7e2), cap(s7e2))
    s7e3 := append(s7, make([]int, 600)...)
    fmt.Printf("s7e3: len: %d, cap: %d\n", len(s7e3), cap(s7e3))
    fmt.Println()
    /*
    The capacity of s7: 1024
    s7e1: len: 1224, cap: 1280
    s7e2: len: 1424, cap: 1696
    s7e3: len: 1624, cap: 2048

    */

    // 示例3。
    s8 := make([]int, 10)
    fmt.Printf("The capacity of s8: %d\n", cap(s8))
    s8a := append(s8, make([]int, 11)...)
    fmt.Printf("s8a: len: %d, cap: %d\n", len(s8a), cap(s8a))
    s8b := append(s8a, make([]int, 23)...)
    fmt.Printf("s8b: len: %d, cap: %d\n", len(s8b), cap(s8b))
    s8c := append(s8b, make([]int, 45)...)
    fmt.Printf("s8c: len: %d, cap: %d\n", len(s8c), cap(s8c))
    /*
    The capacity of s8: 10
    s8a: len: 21, cap: 22
    s8b: len: 44, cap: 44
    s8c: len: 89, cap: 96

    */

    // ---
    /*
    在扩容的过程中，底层数组时不会改动的；当容量超出时，会返回新的切片，底层数组新建；

    确切地说，一个切片的底层数组永远不会被替换。为什么?虽然在扩容的时候 Go 语言一定会生 成新的底层数组，
    但是它也同时生成了新的切片。它是把新的切片作为了新底层数组的窗口，而 没有对原切片及其底层数组做任何改动。
     */
    a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
    fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))
    s9 := a1[1:4]
    //s9[0] = 1
    fmt.Printf("s9: %v (len: %d, cap: %d)\n", s9, len(s9), cap(s9))
    for i := 1; i <= 5; i++ {
        s9 = append(s9, i)
        fmt.Printf("s9(%d): %v (len: %d, cap: %d)\n", i, s9, len(s9), cap(s9))
    }
    fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))
    fmt.Println()
    /*
       a1: [1 2 3 4 5 6 7] (len: 7, cap: 7)
       s9: [2 3 4] (len: 3, cap: 6)
       s9(1): [2 3 4 1] (len: 4, cap: 6)
       s9(2): [2 3 4 1 2] (len: 5, cap: 6)
       s9(3): [2 3 4 1 2 3] (len: 6, cap: 6)
       s9(4): [2 3 4 1 2 3 4] (len: 7, cap: 12)
       s9(5): [2 3 4 1 2 3 4 5] (len: 8, cap: 12)
       a1: [1 2 3 4 1 2 3] (len: 7, cap: 7)
     */
}

/*
1.当两个长度不一的切片使用同一个底层数组，并且两切片的长度均小于数组的容量时，对其 中长度较小的一个切片进行append操作，但不超过底层数组容量，这时会影响长度较长切片
中原来比较小切片多看到的值，因为底层数组被修改了。
2. 数组缩容：可以截取切片的部分数据，然后创建新数组来缩容
 */
