package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("cat", "Tom") //存放进去
	m.Store("mouse", "Jerry")

	// 这里重新读取出来的，就是
	val, ok := m.Load("cat") //拿出来,返回值是any,其实就是interface{}的别名
	if ok {
		fmt.Println(len(val.(string))) //接口的断言,string类型,string的值或者零值,注意这里面如果判断的对象是结构体,那么它也可以是结构体指针,还有nil不能被断言,也就是val不能是nil,结果是false
	}
}
