package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	   1.func NewTimer(d Duration) *Timer
	       创建一个计时器：d时间以后触发，go触发计时器的方法比较特别，就是在计时器的channel中发送值
	*/
	//新建一个计时器：timer
	timer := time.NewTimer(3 * time.Second)
	fmt.Printf("%T\n", timer) //*time.Timer
	fmt.Println(time.Now())   //2019-08-15 10:41:21.800768 +0800 CST m=+0.000461190

	//此处在等待channel中的信号，执行此段代码时会阻塞3秒
	ch2 := timer.C     //<-chan time.Time
	fmt.Println(<-ch2) //2019-08-15 10:41:24.803471 +0800 CST m=+3.003225965
}

