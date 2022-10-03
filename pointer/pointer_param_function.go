package main

import "fmt"

func main() {
	/*
		指针作为参数：

		参数的传递：值传递，引用传递
	*/
	//对于切片就没必要设置指针了，它本身就是个指针，存放的就是地址，是引用了类型，一般是对值类型设置指针，会节省资源，值类型在传递的时候是深拷贝，会赋值一份，浪费资源
	a := 10
	fmt.Println("fun1()函数调用前，a:", a)
	fun1(a)
	fmt.Println("fun1()函数调用后，a：", a)

	fun2(&a) //指针作为参数，那么就传递地址
	fmt.Println("fun2()函数调用后，a:", a)

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("fun3()函数调用前：", arr1)
	fun3(arr1)
	fmt.Println("fun3()函数调用后：", arr1)

	fun4(&arr1)
	fmt.Println("fun4()函数调用后：", arr1)

	s1 := []int{1, 2, 3, 4, 5}
}
func fun4(p2 *[4]int) { // 引用传递
	fmt.Println("fun4()函数中的数组指针：", p2)
	p2[0] = 200
	fmt.Println("fun4()函数中的数组指针：", p2)
}

func fun3(arr2 [4]int) { // 值传递
	fmt.Println("fun3()函数中数组的：", arr2)
	arr2[0] = 100
	fmt.Println("fun3()函数中修改数组：", arr2)
}

func fun1(num int) { // 值传递：num = a = 10
	fmt.Println("fun1()函数中，num的值：", num)
	num = 100
	fmt.Println("fun1()函数中修改num：", num)
}

func fun2(p1 *int) { //传递的是a的地址，就是引用传递，p1=&a
	fmt.Println("fun2()函数中，p1：", *p1)
	*p1 = 200 //函数调用完成后，参数也会销毁了，但是是通过地址更改的，所以main函数中的局部环境变量会修改了
	fmt.Println("fun2()函数中，修改p1：", *p1)
}
