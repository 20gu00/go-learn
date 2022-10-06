package main

import (
	"fmt"
	"time"
)

func main() {

	i := 13
	a := func() {
		fmt.Printf("i is %d \n", i)
	}
	a()

	fmt.Println(ReturnClosure("Tom")())

	Delay()
	time.Sleep(time.Second)
}

func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}

func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		fns = append(fns, func() { //注意这里加进去的是函数,而不是函数调用
			fmt.Printf("hello, this is : %d \n", i)
		})
	}

	for _, fn := range fns {
		fn()
	}
}
