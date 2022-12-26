package main

import (
	"context"
	"fmt"
)

func main() {
	/*
		context接口
		type Context interface {
		    Deadline() (deadline time.Time, ok bool)
		    Done() <-chan struct{}
		    Err() error
		    Value(key interface{}) interface{}
		}

		用于信息穿透的上下文
		上下文信息传递 （request-scoped），比如处理 http 请求、在请求处理链路上传递信息；
		控制子 goroutine 的运行；
		超时控制的方法调用；
		可以取消的方法调用。

		Deadline 方法会返回这个 Context 被取消的截止日期。如果没有设置截止日期，ok 的值是 false。后续每次调用这个对象的 Deadline 方法时，都会返回和第一次调用相同的结果。
		Done 方法返回一个 Channel 对象。在 Context 被取消时，此 Channel 会被 close，如果没被取消，可能会返回 nil。后续的 Done 调用总是返回相同的结果。当 Done 被 close 的时候，你可以通过 ctx.Err 获取错误信息。Done 这个方法名其实起得并不好，因为名字太过笼统，不能明确反映 Done 被 close 的原因，因为 cancel、timeout、deadline 都可能导致 Done 被 close，不过，目前还没有一个更合适的方法名称。
		关于 Done 方法，要记住的知识点就是：如果 Done 没有被 close，Err 方法返回 nil；如果 Done 被 close，Err 方法会返回 Done 被 close 的原因。
		Value 返回此 ctx 中和指定的 key 相关联的 value。

		context.Background()：返回一个非 nil 的、空的 Context，没有任何值，不会被 cancel，不会超时，没有截止日期。一般用在主函数、初始化、测试以及创建根 Context 的时候。
		context.TODO()：返回一个非 nil 的、空的 Context，没有任何值，不会被 cancel，不会超时，没有截止日期。当你不清楚是否该用 Context，或者目前还不知道要传递一些什么上下文信息的时候，就可以使用这个方法。
	*/
	ctx := context.Background()
	fmt.Println(ctx) // context.Background
	// Err方法
	fmt.Println(ctx.Err())
	// 返回一个struct{}类型的接口
	ctx.Done()
	fmt.Println(ctx)
	fmt.Println(ctx.Err())      // ctx取消 channel关闭的原因(被cancel timeout 超过截止时间)
	fmt.Println(ctx.Value("a")) // nil
	t, ok := ctx.Deadline()
	fmt.Println(t, ok) // 0001-01-01 00:00:00 +0000 UTC false 没有设置ctx的截止时间
}
