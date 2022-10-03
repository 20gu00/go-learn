package main

import (
    "time"
    "fmt"
)

func main() {

    /*
        1.func NewTimer(d Duration) *Timer
            创建一个计时器：d时间以后触发，go触发计时器的方法比较特别，就是在计时器的channel中发送值
     */
    //新建一个计时器：timer
    //timer := time.NewTimer(3 * time.Second)
    //fmt.Printf("%T\n", timer) //*time.Timer
    //fmt.Println(time.Now())   //2019-08-15 10:41:21.800768 +0800 CST m=+0.000461190
    //
    此处在等待channel中的信号，执行此段代码时会阻塞3秒
    //ch2 := timer.C     //<-chan time.Time
    //fmt.Println(<-ch2) //2019-08-15 10:41:24.803471 +0800 CST m=+3.003225965

    fmt.Println("-------------------------------")

    //新建计时器，一秒后触发

    timer2 := time.NewTimer(5 * time.Second)

    //新开启一个线程来处理触发后的事件

    go func() {

        //等触发时的信号

        <-timer2.C

        fmt.Println("Timer 2 结束。。")

    }()

    //由于上面的等待信号是在新线程中，所以代码会继续往下执行，停掉计时器

    time.Sleep(3*time.Second)
    stop := timer2.Stop()

    if stop {

        fmt.Println("Timer 2 停止。。")

    }

}


