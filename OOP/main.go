package main

import "fmt"

// 用匿名结构体来模拟继承关系,如果不是匿名,那么就是聚合关系,简单的嵌套关系,用的多,不会自动处理继承关系

//父结构体 实际go中没有子结构体父结构体的说法
type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p *Person) SayHello() {
	fmt.Println(p)
	fmt.Println("this is from Person")
}

//子结构体
type Student struct {
	School string
	Person
}

//一般情况下不使用golang继承，用嵌套结构体替代，因为嵌套结构体比较简单，易读
//type Student struct {
//	School string
//	People *Person
//}

func main() {
	stu := &Student{
		School: "middle",
	}
	//模拟继承的关系，帮我们自动处理
	//stu.Name = stu.Person.Name
	//stu.SayHello() = stu.Person.SayHello()
	//属性和方法都能继承
	stu.Name = "Leo"
	stu.Age = 30
	stu.Sex = "Male"
	fmt.Println(stu) // {  {}}嵌套结构体模拟继承关系
	stu.SayHello()
}
