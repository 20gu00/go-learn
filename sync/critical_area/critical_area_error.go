package main

import (
	"fmt"
	"math/rand"
	"time"
)

//全局变量
var ticket = 10 // 100张票

//Duration 将两个瞬间之间的经过时间表示为 int64 纳秒计数。 该表示将最大可表示持续时间限制为大约 290 年。
//
//time.Duration(seconds) 是进行的类型转换，把我们的整型转换成了time.Duration类型
//
//然后把我们传递的10 * 1000 * 1000 ，这样就是我们想要的结果了
//
//默认是纳秒单位
//
//如果想传递一个10秒的时间进去，需要这样转换，其实就是把我们传递的整型进行了乘法
//
//second := 10
//
//time.Duration(seconds)*time.Second  前面是
//
//time.Second是一个常量

func main() {
	/*
	   4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。
	*/
	go saleTickets("售票口1") // g1,100
	go saleTickets("售票口2") // g2,100
	go saleTickets("售票口3") //g3,100
	go saleTickets("售票口4") //g4,100

	time.Sleep(5 * time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	//for i:=1;i<=100;i++{
	//  fmt.Println(name,"售出：",i)
	//}
	for { //ticket=1
		if ticket > 0 { //g1,g3,g2,g4
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket) // 1 , 0, -1 , -2
			ticket--                         //0 , -1 ,-2 , -3
		} else {
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
	}
}
