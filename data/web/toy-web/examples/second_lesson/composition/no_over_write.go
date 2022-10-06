package main

import "fmt"

func main() {
	son := Son{
		Parent{}, //这个是被组合的
	}

	//重写
	son.SayHello() //可以直接调用被组合的对象的方法,son.Parent.SayHello
}

type Parent struct {
}

func (p Parent) SayHello() {
	fmt.Println("I am " + p.Name()) //这里调用的依然就是Parent
}

func (p Parent) Name() string {
	return "Parent"
}

type Son struct {
	Parent
}

// 定义了自己的 Name() 方法
func (s Son) Name() string {
	return "Son"
}
