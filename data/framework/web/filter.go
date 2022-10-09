package web

import (
	"fmt"
	"time"
)

type Filter func(c *Context)

type FilterBuilder func(next Filter) Filter

//请求进来之前经过的一些列的过滤器,filter演示,主要是用来实现AOP,类似其他框架爱的gin的middlerware,go-zero的interceptor
func MetricFilterBuilder(nextFilter Filter) Filter {
	return func(c *Context) {
		startTime := time.Now().UnixNano()
		nextFilter(c)
		endTime := time.Now().UnixNano()
		fmt.Printf("run time: %d \n", endTime-startTime)
	}
}
