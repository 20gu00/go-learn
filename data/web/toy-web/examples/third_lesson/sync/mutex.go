package main

import (
	"sync"
)

var mutex sync.Mutex
var rwMutex sync.RWMutex

func Mutex() {
	mutex.Lock()
	defer mutex.Unlock() //解锁操作用defer比较稳妥
	// 你的代码
}

func RwMutex() {
	// 加读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 也可以加写锁
	rwMutex.Lock()
	defer rwMutex.Unlock()
}

// 不可重入例子
func Failed1() {
	mutex.Lock()
	defer mutex.Unlock()

	//默认不可重入,跟锁的组成结构有关,它会记住自己的状态,就是有没有上锁,但是不会记住是谁上的锁,我们可以手动改造下,所以不建议在递归函数中使用锁
	// 这一句会死锁
	// 但是如果你只有一个goroutine，那么这一个会导致程序崩溃,其实就算是多个goroutine,panic会沿着函数调用栈向上,如果没有defer也会发生程序的崩溃
	mutex.Lock()
	defer mutex.Unlock()
}

// 不可升级
func Failed2() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 这一句会死锁
	// 但是如果你只有一个goroutine，那么这一个会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}
