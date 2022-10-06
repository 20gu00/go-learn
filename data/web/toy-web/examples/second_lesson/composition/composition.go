package main

import "fmt"

func main() {

}

//组合的思维
// Swimming 会游泳的
type Swimming interface {
	Swim()
}

type Duck interface {
	// 鸭子是会游泳的，所以这里组合了它
	Swimming //注意不能有名称
}

type Base struct {
	Name string
}

type Concrete1 struct {
	Base //嵌套,组合,注意这里不是继承  没有成员名,那类型名既是成员名又是类型名
}

//罕见
type Concrete2 struct {
	*Base
}

func (c Concrete1) SayHello() {
	// c.Name 直接访问了Base的Name字段
	fmt.Printf("I am base and my name is: %s \n", c.Name)
	// 这样也是可以的
	fmt.Printf("I am base and my name is: %s \n", c.Base.Name)

	// 调用了被组合的,可以调用组合的内容的方法

	//其实就是接受这类型对应了
	c.Base.SayHello()
	c.SayHello()
}

func (b *Base) SayHello() {
	fmt.Printf("I am base and my name is: %s \n", b.Name)
}
