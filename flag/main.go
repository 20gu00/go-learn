package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// 定义命令行参数方式
	var name string
	var age int
	var married bool
	var delay time.Duration
	//flag.String("test","a","testdo")
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	// 解析命令行参数
	flag.Parse() // 一次即可
	fmt.Println(name, age, married, delay)
	// 返回命令行参数后的其他参数(非flag参数)(flag正统的命令行参数,arg也算是命令行参数,称非flag)
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数(非flag参数)
	fmt.Println(flag.NArg())
	// 返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

//go run main.go --help

/*[root@localhost flag]# go run main.go -name pprof --age 28 -married=false -d=1h30m
pprof 28 false 1h30m0s
[]
0
4*/

/*[root@localhost flag]# go run main.go a b c (个数不限)
张三 18 false 0s
[a b c]
3
0
*/
