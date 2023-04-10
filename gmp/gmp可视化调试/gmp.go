package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

/*do
1.创建trace文件
2.启动
3.停止
*/

//go tool trace trace.out 打开trace文件  -> view
//goroutine 每一小块都是goroutine，蓝色块的长度意味着数量的，点一下蓝块就会显示当时微秒有多少种不同状态的goroutines，values是数值，Gcwaiting等待再次调度运行，runningable可以运行，runnging运行中
//heap
//thread 跟goroutine显示差不多，可以知道开启了多少个goroutine多少个m
//procs 可以看到设置了多少个processor，还有每个processor做什么，当前这个程序proc主要是比如G1 running main

func main() {
	//trace文件
	//*file结构体
	//程序运行的目录
	//1.创建文件
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	//开始trace
	//io.Writer接口
	//2.启动
	if err = trace.Start(file); err != nil {
		panic(err)
	}

	//3.关闭
	defer trace.Stop()

	//要调试业务
	fmt.Println("trace to do")
}
