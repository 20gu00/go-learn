package main

import "fmt"

func main() {
	/*
	   常量：
	   1.概念：同变量类似，程序执行过程中数值不能改变
	   2.语法：
	       显式类型定义
	       隐式类型定义，就是不指定类型,无类型常量

	   3.常数：
	       固定的数值：100，"abc"


	     跟变量不一样，常量定义了不使用不会报错
	     常量值类型只能是布尔值，数字型和字符串型
	*/
	fmt.Println(100)
	fmt.Println("hello")

	//1.定义常量
	const PATH string = "http:www.baidu.com"
	const PI = 3.14
	fmt.Println(PATH)
	//fmt.Println(PI)

	//2.尝试修改常量的数值
	//PATH = "http://www.sina.com" //cannot assign to PATH

	//3.定义一组常量
	const C1, C2, C3 = 100, 3.14, "haha"
	const (
		MALE   = 0
		FEMALE = 1
		UNKNOW = 3
	)
	//4.一组常量中，如果某个常量没有初始值，默认和上一行一致，值和类型
	const (
		a int = 100
		b
		c string = "ruby"
		d
		e
	)
	fmt.Printf("%T,%d\n", a, a) //格式化标准输出用的格式符合变量一样
	fmt.Printf("%T,%d\n", b, b)
	fmt.Printf("%T,%s\n", c, c)
	fmt.Printf("%T,%s\n", d, d)
	fmt.Printf("%T,%s\n", e, e) //%T输出类型

	//5. 枚举类型：使用常量组作为枚举类型。一组相关数值的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)

}
