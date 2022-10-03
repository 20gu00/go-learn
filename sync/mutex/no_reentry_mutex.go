package main

import (
	"fmt"
	"sync"
)

//当一个线程获取锁时，如果没有其它线程拥有这个锁，那么，这个线程就成功获取到这个锁。之后，如果其它线程再请求这个锁，就会处于阻塞等待的状态。但是，如果拥有这把锁的线程再请求这把锁的话，不会阻塞，而是成功返回，所以叫可重入锁（有时候也叫做递归锁）。只要你拥有这把锁，你可以可着劲儿地调用，比如通过递归实现一些算法，调用者不会阻塞或者死锁。
//了解了可重入锁的概念，那我们来看 Mutex 使用的错误场景。划重点了：Mutex 不是可重入的锁。
//想想也不奇怪，因为 Mutex 的实现中没有记录哪个 goroutine 拥有这把锁。理论上，任何 goroutine 都可以随意地 Unlock 这把锁，所以没办法计算重入条件

func foo1(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}
func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

//sync.Mutext{}结构体实现了Lock和Unlock两个方法,也就是实现了sync.Locaker接口
//注意这里的场景是同一个线程(同一个goroutine),理论上是可以实现可重入锁,但是原生的标准库的sync的mutext不支持罢了
func main() {
	l := &sync.Mutex{} //创建一个mutex的结构体的指针,引用传递
	foo1(l)
}
