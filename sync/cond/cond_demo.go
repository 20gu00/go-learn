package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

//修改条件变量和调用wait()时要加锁
func main() {
	// 返回一个带锁的条件变量
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	for i := 0; i < 10; i++ {
		// 启用了10个goroutine
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()
			log.Printf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者(c.wait)
			c.Broadcast()
		}(i)
	}
	c.L.Lock()
	for ready != 10 {
		// 等待自动解锁 唤醒
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()
	//所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}
