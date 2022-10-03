package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		    随机数的生成依靠的是种子，所以种子如果相同那么生成的随机数也就相同，用seed来设置种子
			   生成随机数random:其实是伪随机数，一位使用一定的算法公式算出来的
			   math.rand
	*/

	num1 := rand.Int() //随机整数
	fmt.Println(num1)

	//范围比较大
	//指定范围
	for i := 0; i < 10; i++ {
		num := rand.Intn(10) //加上个指定范围的功能的，rand.Intn(n)表示范围是[0,n)
		fmt.Println(num)
	}
	rand.Seed(1) //手动更改可以改变随机数值，但这样不够自动化
	num2 := rand.Intn(10)
	fmt.Println(num2)

	//时间每分每秒不一样，根据时间来设置seed
	t1 := time.Now() //获取当前的时间
	fmt.Println(t1)
	fmt.Printf("%T", t1) //打印t1的类型，可见是time类型
	//seed要用的是整形，所以将时间转换成整形
	//时间戳，距离1970.1.1 0点0分0秒，单位：s
	timestamp1 := t1.Unix()
	fmt.Println(timestamp1)

	//第一步生成种子
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

}
