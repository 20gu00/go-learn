package main

import (
	"context"
	"fmt"
	"time"
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

	/*
		WithValue 基于 parent Context 生成一个新的 Context，保存了一个 key-value 键值对。它常常用来传递上下文。(传递给函数)
		WithValue 方法其实是创建了一个类型为 valueCtx 的 Context，它的类型定义如下：
		type valueCtx struct {
		    Context
		    key, val interface{}
		}

		它持有一个 key-value 键值对，还持有 parent 的 Context。它覆盖了 Value 方法，优先从自己的存储中检查这个 key，不存在的话会从 parent 中继续检查。
		Go 标准库实现的 Context 还实现了链式查找。如果不存在，还会向 parent Context 去查找，如果 parent 还是 valueCtx 的话，还是遵循相同的原则：valueCtx 会嵌入 parent，所以还是会查找 parent 的 Value 方法的。
	*/
	ctx = context.TODO()
	ctx = context.WithValue(ctx, "key1", "0001")
	ctx = context.WithValue(ctx, "key2", "0002")
	ctx = context.WithValue(ctx, "key3", "0003")
	ctx = context.WithValue(ctx, "key4", "0004")
	fmt.Println(ctx.Value("key1"))
	//0004->0003->0002->0001

	/*
		WithCancel
		WithCancel 方法返回 parent 的副本，只是副本中的 Done Channel 是新建的对象，它的类型是 cancelCtx。
		我们常常在一些需要主动取消长时间的任务时，创建这种类型的 Context，然后把这个 Context 传给长时间执行任务的 goroutine。当需要中止任务时，我们就可以 cancel 这个 Context，这样长时间执行任务的 goroutine，就可以通过检查这个 Context，知道 Context 已经被取消了。
		WithCancel 返回值中的第二个值是一个 cancel 函数。其实，这个返回值的名称（cancel）和类型（Cancel）也非常迷惑人。
		记住，不是只有你想中途放弃，才去调用 cancel，只要你的任务正常完成了，就需要调用 cancel，这样，这个 Context 才能释放它的资源（通知它的 children 处理 cancel，从它的 parent 中把自己移除，甚至释放相关的 goroutine）。很多同学在使用这个方法的时候，都会忘记调用 cancel，切记切记，而且一定尽早释放。

		当这个 cancelCtx 的 cancel 函数被调用的时候，或者 parent 的 Done 被 close 的时候，这个 cancelCtx 的 Done 才会被 close。
		cancel 是向下传递的，如果一个 WithCancel 生成的 Context 被 cancel 时，如果它的子 Context（也有可能是孙，或者更低，依赖子的类型）也是 cancelCtx 类型的，就会被 cancel，但是不会向上传递。parent Context 不会因为子 Context 被 cancel 而 cancel。
		cancelCtx 被取消时，它的 Err 字段就是下面这个 Canceled 错误：
		var Canceled = errors.New("context canceled")

		withCancel的parent context调用cancel或者Done的channel关闭时,这个cancel类型的context的Done才会关闭
	*/

	ctx2 := context.TODO()
	ctxC, cancel := context.WithCancel(ctx2)
	ctxC2, _ := context.WithCancel(ctxC)
	cancel()
	fmt.Println(ctxC.Err()) // context canceled  !=nil
	fmt.Println(ctxC2.Err())

	// WithTimeout 其实是和 WithDeadline 一样，只不过一个参数是超时时间，一个参数是截止时间。超时时间加上当前时间，其实就是截止时间
	c1 := context.TODO()
	ctx3, _ := context.WithTimeout(c1, 1*time.Second) // cancelFunc
	ctxC4, _ := context.WithCancel(ctx3)
	ctxC5, _ := context.WithCancel(ctxC4)
	//cancel()
	time.Sleep(2 * time.Second)
	fmt.Println(ctxC4.Err()) // context deadline exceeded
	fmt.Println(ctxC5.Err())

	//WithDeadline 会返回一个 parent 的副本，并且设置了一个不晚于参数 d 的截止时间，类型为 timerCtx（或者是 cancelCtx）。
	//如果它的截止时间晚于 parent 的截止时间，那么就以 parent 的截止时间为准，并返回一个类型为 cancelCtx 的 Context，因为 parent 的截止时间到了，就会取消这个 cancelCtx。
	//如果当前时间已经超过了截止时间，就直接返回一个已经被 cancel 的 timerCtx。否则就会启动一个定时器，到截止时间取消这个 timerCtx。
	//综合起来，timerCtx 的 Done 被 Close 掉，主要是由下面的某个事件触发的：
	//截止时间到了；
	//cancel 函数被调用；
	//parent 的 Done 被 close。

	// 两个都要考虑父context,这个决定管控的子context

	// 和 cancelCtx 一样，WithDeadline（WithTimeout）返回的 cancel 一定要调用，并且要尽可能早地被调用，这样才能尽早释放资源，不要单纯地依赖截止时间被动取消
}
