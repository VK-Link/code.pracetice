// 测试 竟态
// 对于变量的更改;
package main

import (
    "fmt"
    "sync"
    "time"
)

var Wait sync.WaitGroup
var Counter int = 0

func main() {
    // Counter = 0
    for routine := 1; routine <= 2; routine++ {
        Wait.Add(1)
        go Routine(routine)
    }

    Wait.Wait()
    fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {
    for count := 0; count < 2; count++ {
        value := Counter
        time.Sleep(1 * time.Nanosecond)
        value++
        Counter = value
    }

    Wait.Done()
}
