package main

import (
	"fmt"
	"time"
)

//FilterBuilder
type FilterBuilder func(next Filter) Filter //形参是下一个filter,返回一个filter

type handlerFunc func(c *Context)

//洋葱结构
type Filter func(c *Context)

//这是一种很方便的判断类型的方法
//var _ a=b  判断b是不是a类型
var _ FilterBuilder = MetricsFilterBuilder //函数名

//用一个函数作为一个返回值,形参也是个函数,形参函数调用时使用返回值函数的参数
//实现一个FilterBuilder
//一般就是handler中持有多个filter,由handler调用filter
func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		//请求前记录时间戳
		start := time.Now().Nanosecond()
		next(c) //下一个filter
		//请求后记录时间戳
		end := time.Now().Nanosecond()
		fmt.Printf("用了%d纳秒", end-start) //执行时间
	}
}
