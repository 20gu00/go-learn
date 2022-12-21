package main

import "fmt"

//结构体实现了接口中的所有方法，那么我们就说这个结构体实现了这个接口
//类型转换，结构体转接口，接口转结构体
//实现接口后，就能实现特定的功能，比如 golang自定义排序 swap() less() len() sort()

//同一个结构对象,实际值可以不同,也就对应了不同的方法,多种形态
type Animal interface {
	Say() string
}

type Cat struct {
	Name string
}

func (c *Cat) Say() string {
	return c.Name + "喵喵喵"
}

type Dog struct {
	Name string
}

func (d *Dog) Say() string {
	return d.Name + "汪汪汪"
}

func main() {
	c := &Cat{Name: "小白猫"}
	d := &Dog{Name: "大黄狗"}
	var p1 Animal
	p1 = c
	fmt.Println(p1.Say())
	p1 = d
	fmt.Println(p1.Say())
	fmt.Println(transData(c))
	fmt.Println(transData(d))
}

//类型转换举例,隐式转换，将结构体类型转换为接口类型
func transData(a Animal) string {
	return fmt.Sprintf("%s%s", a.Say(), "处理后")
}
