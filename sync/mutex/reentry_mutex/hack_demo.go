package recursiveMutex

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

//首先，我们获取goroutine的 g 指针，反解出对应的 g 的结构。每个运行的 goroutine 结构的 g 指针保存在当前 goroutine 的一个叫做 TLS 对象中。
//第一步：我们先获取到 TLS 对象；
//第二步：再从 TLS 中获取 goroutine 结构的 g 指针；
//第三步：再从 g 指针中取出 goroutine id。
//需要注意的是，不同 Go 版本的 goroutine 的结构可能不同，所以需要根据 Go 的不同版本进行调整。当然了，如果想要搞清楚各个版本的 goroutine 结构差异，所涉及的内容又过于底层而且复杂，直接使用第三方的库来获取 goroutine id 就可以了。
//现在已经有很多成熟的方法了，可以支持多个 Go 版本的 goroutine id，给你推荐一个常用的库：petermattis/goid。
//尽管拥有者可以多次调用 Lock，但是也必须调用相同次数的 Unlock，这样才能把锁释放掉。这是一个合理的设计，可以保证 Lock 和 Unlock 一一对应。

// RecursiveMutex 包装一个Mutex,实现可重入
type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 这个goroutine 重入的次数
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
	if atomic.LoadInt64(&m.owner) == gid { //原子加载*int,返回值是int
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.owner, gid) //原子操作保存到*int中
	m.recursion = 1                  //重入次数唯一,即第一次上锁
}
func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
