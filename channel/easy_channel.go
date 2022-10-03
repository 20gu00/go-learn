package main

import "fmt"

func main() {
	var a chan int //不建议这种方式
	if a == nil {
		//fmt.Println("channel 是 nil 的, 不能使用，需要先创建通道。。")
		//nil chan，不能使用，类似于nil map，不能直接存储键值对
		//就是零值不可用,不像是slice零值可用
		a = make(chan int) //这里重新make,没关系
		fmt.Printf("数据类型是： %T", a)
	}
}
