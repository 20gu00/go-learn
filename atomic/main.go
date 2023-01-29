package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// add store load compare swap
	var a int64
	go func() {
		atomic.AddInt64(&a, 10)
		fmt.Println(a)
	}()

	// 比较是不是等于10,是就替换成20
	go func() {
		atomic.CompareAndSwapInt64(&a, 10, 20)
		fmt.Println(a)
	}()

	go func() {
		atomic.SwapInt64(&a, 100)
		fmt.Println(a)
	}()

	go func() {
		fmt.Println("load", atomic.LoadInt64(&a))
	}()

	go func() {
		atomic.StoreInt64(&a, 111)
		fmt.Println("load2", atomic.LoadInt64(&a))
	}()

	time.Sleep(10 * time.Second)
}
