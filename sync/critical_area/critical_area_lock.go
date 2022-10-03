package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//其实换个角度来看,锁也是竞争的,有个方式就是使用条件变量去通知,就不同go程去竞争了
//全局变量
var ticket = 10 // 100张票

var wg sync.WaitGroup
var matex sync.Mutex // 创建锁头

func main() {
	/*
	   4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。
	*/
	wg.Add(4)
	go saleTickets("售票口1") // g1,100
	go saleTickets("售票口2") // g2,100
	go saleTickets("售票口3") //g3,100
	go saleTickets("售票口4") //g4,100
	wg.Wait()              // main要等待。。。

	//time.Sleep(5*time.Second)
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	//for i:=1;i<=100;i++{
	//  fmt.Println(name,"售出：",i)
	//}
	for { //ticket=1
		matex.Lock()
		if ticket > 0 { //g1,g3,g2,g4
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket) // 1 , 0, -1 , -2
			ticket--                         //0 , -1 ,-2 , -3
		} else {
			matex.Unlock() //解锁
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
		matex.Unlock() //解锁
	}
}
