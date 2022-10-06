// Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package middleware

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gohade/hade/framework/gin"
)

var defaultTimeOut = TimeOut{
	d: 10 * time.Second,
}

// TimeOut struct
type TimeOut struct {
	d time.Duration
}

// TimeOutOption ...
type TimeOutOption func(out *TimeOut)

// WithTimeOut ...
func WithTimeOut(d time.Duration) TimeOutOption {
	return func(out *TimeOut) {
		out.d = d
	}
}

// NewTimeOut ...
func NewTimeOut(opts ...TimeOutOption) *TimeOut {
	out := defaultTimeOut
	for _, opt := range opts {
		opt(&out)
	}

	return &out
}

func (t *TimeOut) Func() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		// 执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), t.d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 使用next执行具体的业务逻辑
			c.Next()

			finish <- struct{}{}
		}()
		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			c.ISetStatus(500).IJson("time out")
			log.Println(p)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.ISetStatus(500).IJson("time out")
		}
	}
}
