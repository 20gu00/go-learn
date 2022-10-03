package main

import "fmt"

func main() {
	/*
		方法：method
			一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
			所有给定类型的方法属于该类型的方法集
		本质就是讲方法接受者参数作为第一个形参的函数(recieve)

		方法与函数很类似
		方法是包含接受者的函数
		命名类型也就是常见的int之类的，一般用结构体和结构体指针比较多

		语法：
			func （接受者） 方法名(参数列表)(返回值列表){

			}

		总结：method，同函数类似，区别需要有接受者。(也就是调用者)

		对比函数：
			A：意义
				方法：某个类别的行为功能，需要指定的接受者调用
				函数：一段独立功能的代码，可以直接调用

			B：语法
				方法：方法名可以相同，只要接受者不同
				函数：命名不能冲突

	*/
	w1 := Worker{name: "王二狗", age: 30, sex: "男"}
	w1.work()

	w2 := &Worker{name: "Ruby", age: 34, sex: "女"}
	fmt.Printf("%T\n", w2)
	w2.work()

	w2.rest() //不管方法定义时定义的接受者是结构体还是结构体指针，这里w1这个结构体变量和w2这个结构体指针代表的接受着一样
	w1.rest()

	w2.printInfo()
	c1 := Cat{color: "白色的", age: 1}
	c1.printInfo()

}

//1.定义一个工人结构体
type Worker struct {
	//字段
	name string
	age  int
	sex  string
}

type Cat struct {
	color string
	age   int
}

//2.定义行为方法

func (w Worker) work() { //w = w1，值传递
	fmt.Println(w.name, "在工作。。。")
}

func (p *Worker) rest() { //p = w2 ,p = w1的地址
	fmt.Println(p.name, "在休息。。")
}

//同样名称的方法不同接受者调用效果不同，可以在调用处Ctrl+B跳转到实际使用的方法的调用处
func (p *Worker) printInfo() {
	fmt.Printf("工人姓名：%s，工人年龄：%d，工人性别：%s\n", p.name, p.age, p.sex)
}

func (p *Cat) printInfo() {
	fmt.Printf("猫咪的颜色：%s，年龄：%d\n", p.color, p.age)
}
