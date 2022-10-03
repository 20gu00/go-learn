package main

import (
    "fmt"
    "time"
)

func main()  {
    a := 1
    go func() {
        a = 2
        fmt.Println("子goroutine。。",a)
    }()
    a = 3
    time.Sleep(1)
    fmt.Println("main goroutine。。",a)
}

