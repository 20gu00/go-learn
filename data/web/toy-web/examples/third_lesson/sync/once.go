package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
}

var once sync.Once //注意是全局变量(包级变量),如果是局部变量,那么每次进入局部代码块都会var一次,都能执行Once

// 这个方法，不管调用几次，只会输出一次,一般是用来做初始化,防止多次初始化
func PrintOnce() {
	once.Do(func() {
		fmt.Println("只输出一次")
	})
}
