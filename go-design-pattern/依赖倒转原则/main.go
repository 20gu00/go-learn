package main

import "fmt"

//------------抽象层

//抽象出汽车和司机
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

//------------实现层

//benz
type Benz struct {
}

func (*Benz) Run() {
	fmt.Println("benz run")
}

//bmw
type Bmw struct {
}

func (*Bmw) Run() {
	fmt.Println("bmw run")
}

//张三李四
type Zhangsan struct {
}

func (*Zhangsan) Drive(car Car) {
	fmt.Println("zhangsan drive", car)
	car.Run()
}

type Lisi struct {
}

func (*Lisi) Drive(car Car) {
	fmt.Println("lisi drive", car)
	car.Run()
}

//这样就可以很好实现解耦,不然不同的司机,不同的车,关系很复杂,全都是struct,耦合度会很高
//在添加张三李四的时候只用关注Drive()

//业务逻辑层
func main() {
	//所以实现层和业务逻辑层都是依赖抽象层,都面向抽象层

	var benz Car
	benz = new(Benz)

	var zhangsan Driver
	//不赋值就是未初始化,空指针
	zhangsan = new(Zhangsan)
	zhangsan.Drive(benz)
}