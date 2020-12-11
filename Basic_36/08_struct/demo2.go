package main

import "fmt"

type Cat struct {
    name           string // 名字。
    scientificName string // 学名。
    category       string // 动物学基本分类。
}

func New(name, scientificName, category string) Cat {
    return Cat{
        name:           name,
        scientificName: scientificName,
        category:       category,
    }
}

func (cat *Cat) SetName(name string) {
    cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
    cat.name = name
}

func (cat Cat) Name() string {
    return cat.name
}

func (cat Cat) ScientificName() string {
    return cat.scientificName
}

func (cat Cat) Category() string {
    return cat.category
}

func (cat Cat) String() string {
    return fmt.Sprintf("%s (category: %s, name: %q)",
        cat.scientificName, cat.category, cat.name)
}

func main() {
    cat := New("little pig", "American Shorthair", "cat")
    cat.SetName("monster") // (&cat).SetName("monster")
    fmt.Printf("The cat: %s\n", cat)

    cat.SetNameOfCopy("little pig")
    fmt.Printf("The cat: %s\n", cat)

    type Pet interface {
        SetName(name string)
        Name() string
        Category() string
        ScientificName() string
    }

    // 一个类型的方法集合中有哪些方法与它能实现哪些接口类型是息息相关的。
    // 如果一个基本类型和它的指针类型的方法集合是不同的，那么它们具体实现的接口类型的数量就也会有差异，除非这两个数量都是零
    _, ok := interface{}(cat).(Pet)
    fmt.Printf("Cat implements interface Pet: %v\n", ok)
    _, ok = interface{}(&cat).(Pet) // 指针类型实现了接口 Pet
    fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}

/*
The cat: American Shorthair (category: cat, name: "monster")
The cat: American Shorthair (category: cat, name: "monster")
Cat implements interface Pet: false
*Cat implements interface Pet: true
*/



/*
1. 我们可以在结构体类型中嵌入某个类型的指针类型吗?如果可以，有哪些注意事项?
- 可以在结构体中嵌入某个类型的指针类型， 它和普通指针类似，默认初始化为nil,因此在用之前需要人为初始化，否则可能引起错误

2. 字面量struct{}代表了什么?又有什么用处?
- 空结构体不占用内存空间，但是具有结构体的一切属性，如可以拥有方法，可以 写入channel。所以当我们需要使用结构体而又不需要具体属性时可以使用它
 */