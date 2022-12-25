package web

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Hook 是一个钩子函数。注意，
// ctx 是一个有超时机制的 context.Context
// 所以你必须处理超时的问题,防止hook运行太久,钩子的意思,初始化或者优雅关闭,或者一些用户的操作比如文件关闭,数据库连接断开.这里利用了context包的Context做同步
type Hook func(ctx context.Context) error

// BuildCloseServerHook 这里其实可以考虑使用 errgroup，
// 但是我们这里不用是希望每个 service 单独关闭
// 互相之间不影响
//这里构造了一个hook的builder函数
func BuildCloseServerHook(servers ...Server) Hook {
	return func(ctx context.Context) error {
		wg := sync.WaitGroup{}
		doneCh := make(chan struct{}) //当goroutine全部关了就玩这里发个信号,空结构体拿来做信号
		wg.Add(len(servers))          //其实就是server的数量

		for _, s := range servers {
			go func(svr Server) {
				err := svr.Shutdown(ctx)
				if err != nil {
					fmt.Printf("service shutdown error: %v \n", err)
				}
				time.Sleep(time.Second)
				wg.Done() //最好在goroutine中关
			}(s)
		}
		//原始的写法是
		//wg.Wait()  就是停下来等待,阻塞,等goroutine全部关了,有个缺点就是可能有个server一直关不了
		go func() {
			wg.Wait() //这里只是阻塞了这个goroutine,知道全部的goroutine关闭
			doneCh <- struct{}{}
		}()
		//一个是超时了没关掉,一个是全部关了
		select { //如果case都不满足,会阻塞在这里,直到有个case满足
		case <-ctx.Done(): //该方法返回个 <-chan struct{},超时的时候会调用这个方法
			fmt.Printf("closing servers timeout \n")
			return ErrorHookTimeout
		case <-doneCh:
			fmt.Printf("close all servers \n")
			return nil
		}
	}
}
