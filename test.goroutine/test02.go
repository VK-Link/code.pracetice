package main

import (
    "fmt"
    "math/rand"
    "time"
)

func GetAwardUserName(users map[string]int64) (name string) {
    sizeOfUsers := len(users)
    awardIndex := rand.Intn(sizeOfUsers)

    var index int
    for uName, _ := range users {
        if index == awardIndex {
            name = uName
            return
        }
        index += 1
    }
    return
}

func main() {
    var users map[string]int64 = map[string]int64{
        "a": 10,
        "b": 6,
        "c": 3,
        "d": 12,
        "e": 20,
        "f": 1,     // struct 初始化 var a int =
    }

    rand.Seed(time.Now().Unix())
    awardStat := make(map[string]int64)
    for i := 0; i < 1000; i += 1 {
        name := GetAwardUserName(users)
        if count, ok := awardStat[name]; ok {
            awardStat[name] = count + 1
        } else {
            awardStat[name] = 1
        }
    }

    for name, count := range awardStat {
        fmt.Printf("user: %s, award count: %d\n", name, count)
    }

    return
}