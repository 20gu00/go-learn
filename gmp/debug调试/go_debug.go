package main

import (
	"fmt"
	"time"
)

//go build -o godebug go_debug.go
//GODEBUG=schedtrace=1000 ./godebug
//1000毫秒，终端上debug调试就是1秒打印一次数据
//SCHED 8047ms: gomaxprocs=4 idleprocs=4 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0]
//调试信息 程序启动到输出的时间 P的数量 空闲的P 线程数量包括M0和终端debug调试线程 自旋的线程数量 空闲的线程数量 全局的G队列中的等待运行的G数量 每个P的本地队列
func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("gmp")
	}
}
