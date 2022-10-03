package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

//Goroutine的任何返回值被忽略,一般go语句后是函数,一般称为go函数,可以用channel来在不同的goroutine之间通信
func main() {
	go hello()
	fmt.Println("main function")
}
