package main

import (
	"fmt"
	"sync"
)

type Counter2 struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter2
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁  值类型,深拷贝,复制数据传递
}

//Mutex 是一个有状态的对象，它的 state 字段记录这个锁的状态。如果你要复制一个已经加锁的 Mutex 给一个新的变量，那么新的刚初始化的变量居然被加锁了，这显然不符合你的期望，因为你期望的是一个零值的 Mutex。关键是在并发环境下，你根本不知道要复制的 Mutex 状态是什么，因为要复制的 Mutex 是由其它 goroutine 并发访问的，状态可能总是在变化。
// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter2) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
