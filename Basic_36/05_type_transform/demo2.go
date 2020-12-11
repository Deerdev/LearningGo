package main

import (
    "fmt"
)

// string 和 int 转换
// 首先，对于整数类型值、整数常量之间的类型转换，原则上只要源值在目标类型的可表示范围内 就是合法的。
// 第二，虽然直接把一个整数值转换为一个string类型的值是可行的，但值得关注的是，被转换 的整数值应该可以代表一个有效的 Unicode 代码点，否则转换的结果将会是"�"(仅由高亮的 问号组成的字符串值)。
// 第三个知识点是关于string类型与各种切片类型之间的互转的。

func main() {
    // 重点1的示例。
    /*
    变量srcInt的值是int16类型的-255，而变量dstInt的值是由前者转换而来的，类型是 int8。int16类型的可表示范围可比int8类型大了不少。问题是，dstInt的值是多少?
    首先你要知道，整数在 Go 语言以及计算机中都是以补码的形式存储的。这主要是为了简化计算 机对整数的运算过程。补码其实就是原码个位求反再加 1。
    比如，int16类型的值-255的补码是1111111100000001。如果我们把该值转换为int8类型 的值，那么 Go 语言会把在较高位置(或者说最左边位置)上的 8 位二进制数直接截掉，从而 得到00000001。
    又由于其最左边一位是0，表示它是个正整数，以及正整数的补码就等于其原码，所以dstInt 的值就是1。
    一定要记住，当整数值的类型的有效范围由宽变窄时，只需在补码形式下截掉一定数量的高位二 进制数即可。
    类似的快刀斩乱麻规则还有:当把一个浮点数类型的值转换为整数类型值时，前者的小数部分会 被全部截掉。
     */
    var srcInt = int16(-255)
    // 请注意，之所以要执行uint16(srcInt)，是因为只有这样才能得到全二进制的表示。
    // 例如，fmt.Printf("%b", srcInt)将打印出"-11111111"，后者是负数符号再加上srcInt的绝对值的补码。
    // 而fmt.Printf("%b", uint16(srcInt))才会打印出srcInt原值的补码"1111111100000001"。
    fmt.Printf("The complement of srcInt: %b (%b)\n",
        uint16(srcInt), srcInt)
    dstInt := int8(srcInt)
    fmt.Printf("The complement of dstInt: %b (%b)\n",
        uint8(dstInt), dstInt)
    fmt.Printf("The value of dstInt: %d\n", dstInt)
    fmt.Println()

    // 重点2的示例:字符'�'的 Unicode 代码点是U+FFFD。它是 Unicode 标准中定义的 Replacement Character，专用于替换那些未知的、不被认可的以及无法展示的字符。
    fmt.Printf("The Replacement Character: %s\n", string(-1)) // �
    fmt.Printf("The Unicode codepoint of Replacement Character: %U\n", '�') // unicode U+FFFD
    fmt.Println()

    // 重点3的示例:一个值在从string类型向[]byte类型转换时代表着以 UTF-8 编码的字符串 会被拆分成零散、独立的字节。
    srcStr := "你好"
    fmt.Printf("The string: %q\n", srcStr)
    fmt.Printf("The hex of %q: %x\n", srcStr, srcStr) // "你好": e4bda0e5a5bd
    fmt.Printf("The byte slice of %q: % x\n", srcStr, []byte(srcStr)) // "你好": e4 bd a0 e5 a5 bd
    // 除了与 ASCII 编码兼容的那部分字符集，以 UTF-8 编码的某个单一字节是无法代表一个字符 的。
    // 比如，UTF-8 编码的三个字节\xe4、\xbd和\xa0合在一起才能代表字符'你'，而 \xe5、\xa5和\xbd合在一起才能代表字符'好'。
    fmt.Printf("The string: %q\n", string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}))
    fmt.Printf("The rune slice of %q: %U\n", srcStr, []rune(srcStr))    // "你好": [U+4F60 U+597D]
    // 一个值在从string类型向[]rune类型转换时代表着字符串会被拆分成一个个 Unicode 字符
    fmt.Printf("The string: %q\n", string([]rune{'\u4F60', '\u597D'}))
}

// unicode: https://home.unicode.org/