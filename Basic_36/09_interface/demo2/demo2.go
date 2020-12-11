package main

import (
    "fmt"
    "reflect"
)

type Pet interface {
    Name() string
    Category() string
}

type Dog struct {
    name string // 名字。
}

func (dog *Dog) SetName(name string) {
    dog.name = name
}

func (dog Dog) Name() string {
    return dog.name
}

func (dog Dog) Category() string {
    return "dog"
}

func main() {
    /*
    接口类型 本身是无法被值化的。在我们赋予它实际的值之前，它的值一定会是nil，这也是它的零值。
    一旦它被赋予了某个实现类型的值，它的值就不再是nil了。不过要注意，即使我们 像前面那样把dog的值赋给了pet，pet的值与dog的值也是不同的。
    这不仅仅是副本与原值的 那种不同。
    当我们给一个接口变量赋值的时候，该变量的动态类型会与它的动态值一起被存储在一个专用的数据结构中。
    严格来讲，这样一个变量的值其实是这个专用数据结构的一个实例，而不是我们赋给该变量的那 个实际的值。
    所以我才说，pet的值与dog的值肯定是不同的，无论是从它们存储的内容，还是 存储的结构上来看都是如此。不过，我们可以认为，这时pet的值中包含了dog值的副本。
    我们就把这个专用的数据结构叫做iface吧，在 Go 语言的runtime包中它其实就叫这个名字。
    iface的实例会包含两个指针，一个是指向类型信息的指针，另一个是指向动态值的指针。
    这里 的类型信息是由另一个专用数据结构的实例承载的，其中包含了动态值的类型，以及使它实现了 接口的方法和调用它们的途径，等等。
    */
    // 示例1:
    var dog1 *Dog
    fmt.Println("The first dog is nil.")
    dog2 := dog1
    fmt.Println("The second dog is nil.")
    // nil 的 struct 赋值给 interface，interface 并不是 nil
    var pet Pet = dog2
    if pet == nil {
        fmt.Println("The pet is nil.")
    } else {
        fmt.Println("The pet is not nil.")
    }
    fmt.Printf("The type of pet is %T.\n", pet)
    fmt.Printf("The type of pet is %s.\n", reflect.TypeOf(pet).String())
    fmt.Printf("The type of second dog is %T.\n", dog2)
    fmt.Println()

    // 示例2。
    wrap := func(dog *Dog) Pet {
        if dog == nil {
            return nil
        }
        return dog
    }
    pet = wrap(dog2)
    if pet == nil {
        fmt.Println("The pet is nil.")
    } else {
        fmt.Println("The pet is not nil.")
    }
}
/*
The first dog is nil.
The second dog is nil.
The pet is not nil.
The type of pet is *main.Dog.
The type of pet is *main.Dog.
The type of second dog is *main.Dog.
 */