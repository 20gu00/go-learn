package main

import (
	"fmt"
	"time"
)

func main() {
	// 这个不能在 main 函数运行，是因为运行起来，
	// 所有的goroutine都被我们搞sleep了，直接就崩了
	//Select()
}

func Select() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	//两个goroutine其实都和快运行了
	go func() {
		time.Sleep(time.Second)
		ch1 <- "msg from channel1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "msg from channel2"
	}()

	for {
		//就是都满足,随便挑一个,也就是select顺序不保证的,不是顺序匹配的
		select { //可以认为两个chanel都被放了数据,阻塞在那了,所以下面两个case都有可能被执行
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		default: //注意default,因为很容易就进来了,搭配了for循环使用,实际效果就是chanel中没有数据就等待一下,再去看chanel中有没有数据
			time.Sleep(time.Second)
		}
	}
}
