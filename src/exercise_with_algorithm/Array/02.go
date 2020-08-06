/*
题目14: 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
示例1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释:

输入不存在公共前缀。
说明：

所有输入只包含小写字母 a-z

*/
package main

import (
    "fmt"
    "strings"
)

/*
["flow","flower","flight"]
以第一个元素为基础，和后面的元素对比：
1. strings.Index("flower", "flow") == 0, "flow" 为最长前缀
2. strings.Index("flight", "flow") != 0, "flow" 不是最长前缀，"flow"向左截断，使用"flo"继续尝试
3. strings.Index("flight", "flo") != 0, "flo" 不是最长前缀，"flo"向左截断，使用"fl"继续尝试
3. strings.Index("flight", "fl") == 0, "fl" 是最长前缀
 */

func commonPrefix(a1 []string) string  {
    if len(a1) < 1 {
        return ""
    }
    prefix := a1[0]
    for _, v := range a1 {
        for strings.Index(v, prefix) != 0 {
            if len(prefix) < 1 {
                return ""
            }
            prefix = prefix[:len(prefix) - 1]
        }
    }
    return prefix
}

func main() {
    fmt.Print(commonPrefix([]string{"flower", "flow", "flight"}))
}