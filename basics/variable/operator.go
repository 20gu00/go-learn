package main

import "fmt"

func main() {
	/*
	   算术运算符：+，-，*,/,%,++,--
	   +
	   -
	   *,乘法
	   /：取商，两个数相除，取商
	   %：取余，取模，两个数相除，取余数

	   ++：给自己加1
	       i++
	   --：给自己减1
	       i--
	   自增自减运算符的使用要注意它的先后，比如i++是先自增1在整体表达式的值，++i是先整体表达式的值再到i的值自增一
	   整数
	*/
	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	sub := a - b
	fmt.Printf("%d - %d = %d\n", a, b, sub)

	mul := a * b
	fmt.Printf("%d * %d = %d\n", a, b, mul)

	div := a / b //取商，或者说取整
	mod := a % b //取余
	fmt.Printf("%d / %d = %d\n", a, b, div)
	fmt.Printf("%d %% %d = %d\n", a, b, mod)

	c := 3
	c++ //就是给c+1
	fmt.Println(c)

	c-- //给c减1
	fmt.Println(c)

} //+= -= *= %= 取余取商
