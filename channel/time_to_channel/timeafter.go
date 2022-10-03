package main

import (
    "time"
    "fmt"
)

func main() {

    /*
        func After(d Duration) <-chan Time
            返回一个通道：chan，存储的是d时间间隔后的当前时间。
     */
    ch1 := time.After(3 * time.Second) //3s后
    fmt.Printf("%T\n", ch1) // <-chan time.Time
    fmt.Println(time.Now()) //2019-08-15 09:56:41.529883 +0800 CST m=+0.000465158
    time2 := <-ch1
    fmt.Println(time2) //2019-08-15 09:56:44.532047 +0800 CST m=+3.002662179

}

