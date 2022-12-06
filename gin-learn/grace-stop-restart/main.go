package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	server := http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second, //time.Duration  单位默认是纳秒
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //1的20次幂
	}

	go func() {
		//要开启协程,不然就一直卡在ListenAndServer()
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("http server启动失败")
		}
	}()

	//监听信号做优雅关闭
	quit := make(chan os.Signal) //默认0,1效果也差不多
	//kill默认syscall.SIGTERM
	//kill -2 或者 Ctrl+c syscall.SIGINT
	//kill -9不会被捕获
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	//创建一个超时的context
	//父context
	c2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(c2); err != nil {
		log.Fatal("server不正常退出,shutdown") //defer也不执行
	}
	log.Printf("server退出了")

	//go func() {
	//	stop := common.SignalHandler()
	//	select {
	//	case <-stop: //当通道关闭
	//		if *kind == "admin" {
	//			router.HttpServerStop()  //这里边其实就有http的shutdown,也就是某个server可以正常实现优雅退出逻辑,主函数打印提示即可
	//			<-time.NewTimer(10 * time.Second).C  //设置格式件延迟,额外设置了总的优雅退出时间
	//			os.Exit(0)
	//		} else if *kind == "proxy" {
	//			httpProxyServer.HttpProxyServerStop()
	//			httpProxyServer.HttpsProxyServerStop()
	//			<-time.NewTimer(10 * time.Second).C
	//			//<-time.After(10*time.Minute)
	//			os.Exit(0)
	//		} else {
	//			router.HttpServerStop()
	//			httpProxyServer.HttpProxyServerStop()
	//			httpProxyServer.HttpsProxyServerStop()
	//			<-time.NewTimer(10 * time.Second).C
	//			os.Exit(0)
	//		}
	//	}
	//}()

	//优雅重启
	//fvbock/endless替换ListenAndServe()
	//比如说服务升级
	//新开了一个子进程,新的请求就新的服务处理,处理重启前还在处理中的请求父进程处理,等旧的请求处理完了,父进程关闭,平滑更新,子进程的服务全面
	//kill -1 pid发送syscall.SIGINT信号通知程序优雅重启(-2)

	//r := gin.Default()
	//	r.GET("/ping", func(ctx *gin.Context) {
	//		ctx.String(http.StatusOK, "ok")
	//	})

	//endless.ListenAndServe(":8080",r)

	//如果是用了supviser监控进程,自动拉起关闭的进程就不适合使用endless,有可能会冲突
	//云流行的各种发布方式 滚动跟新 蓝绿 金丝雀 ab等,实际上能有针对应用的优雅关闭需求就不错了,很多时候都用不上
}

//监听信号的函数
func SignalHandler() chan struct{} { //<-chan struct{}
	stop := make(chan struct{})
	c := make(chan os.Signal, 2) //可以容纳的两个信号
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	//启用了一个goroutine
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1)
	}()
	return stop //返回一个通道,该通道可以被关闭
}
