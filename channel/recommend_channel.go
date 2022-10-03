package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	fmt.Printf("%T,%p\n", ch1, ch1)

	test1(ch1)

}

//channel是引用类型的数据，在作为参数传递的时候，传递的是内存地址。

func test1(ch chan int) {
	fmt.Printf("%T,%p\n", ch, ch)
}
