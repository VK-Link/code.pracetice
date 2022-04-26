// 匿名函数  defer 返回值等

package main

import (
    "fmt"
)

func main() {
    fmt.Println("=========================")
    fmt.Println("return:", fun1())

    fmt.Println("=========================")
    fmt.Println("return:", fun2())
    fmt.Println("=========================")

    fmt.Println("return:", fun3())
    fmt.Println("=========================")

    fmt.Println("return:", fun4())
}

// 有名返回值;
func fun1() (i int) {
    defer func() {
        i++
        fmt.Println("func1-defer2:", i) // 打印结果为 defer2: 2
    }()

    // 规则二 defer执行顺序为先进后出
    defer func() {
        i++
        fmt.Println("func1-defer1:", i) // 打印结果为 defer1: 1
    }()

    // 规则三 defer可以改变有名返回参数的值
    return 0 // 这里实际结果为2;如果是return 100呢
}

func fun2() int {
    var i int
    defer func() {
        i++
        fmt.Println("defer2:", i) // 打印结果为 defer2: 2
    }()

    defer func() {
        i++
        fmt.Println("defer1:", i) // 打印结果为 defer1: 1
    }()
    return i
}

func fun3() (r int) {
    t := 5
    defer func() {
        t = t + 5
        fmt.Println("func3-defer", t)
    }()
    return t
}

func fun4() int {
    i := 8
    // 规则一 defer后面的函数参数会被实时解析
    defer func(i int) {
        i = 99
        fmt.Println("func4-defer", i)
    }(i)
    i = 19
    return i
}
