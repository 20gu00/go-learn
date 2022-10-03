package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包：字符串和基本类型之前的转换
			string convert
	*/
	//fmt.Println("aa"+100)
	//1.bool类型
	s1 := "true" //得是true或者false
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1)

	ss1 := strconv.FormatBool(b1) //bool->string
	fmt.Printf("%T,%s\n", ss1, ss1)

	//2.整数
	//s：数字的字符串形式
	//base：数字字符串的进制，比如：2进制、10进制。
	//bitSize：预期数值的bit大小，用于数值上限限制，最终返回的还是int64类型
	s2 := "100"
	i2, err := strconv.ParseInt(s2, 2, 64) //(s2,10,0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i2, i2)

	ss2 := strconv.FormatInt(i2, 10)
	fmt.Printf("%T,%s\n", ss2, ss2)

	//itoa(),atoi()
	i3, err := strconv.Atoi("-42") //转为int类型
	fmt.Printf("%T,%d\n", i3, i3)
	ss3 := strconv.Itoa(-42)
	fmt.Printf("%T,%s\n", ss3, ss3)
}
